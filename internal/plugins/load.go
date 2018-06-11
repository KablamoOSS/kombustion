package plugins

import (
	"fmt"
	"os"
	"plugin"
	"runtime"

	log "github.com/sirupsen/logrus"

	"github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/KablamoOSS/kombustion/internal/plugins/lock"
	pluginTypes "github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
	kombustionTypes "github.com/KablamoOSS/kombustion/types"
)

// LoadPlugins for the project
func LoadPlugins() (resources, outputs, mappings map[string]kombustionTypes.ParserFunc) {
	resources, outputs, mappings =
		make(map[string]kombustionTypes.ParserFunc),
		make(map[string]kombustionTypes.ParserFunc),
		make(map[string]kombustionTypes.ParserFunc)

	lockFile, err := lock.FindAndLoadLock()
	if err != nil {
		log.Fatal(err)
	}

	manifestFile := manifest.FindAndLoadManifest()
	if err != nil {
		log.Fatal(err)
	}

	// Load all plugins the manifest
	for _, manifestPlugin := range manifestFile.Plugins {
		for _, plugin := range lockFile.Plugins {
			// Find the matching plugin in the lock file
			if manifestPlugin.Name == plugin.Name && manifestPlugin.Version == plugin.Version {
				for _, resolved := range plugin.Resolved {
					// Find the right plugin for the current OS/Arch
					if runtime.GOOS == resolved.OperatingSystem &&
						runtime.GOARCH == resolved.Architecture {
						loadPlugin(manifestPlugin, plugin.Name, plugin.Version, resolved.PathOnDisk, resources, outputs, mappings)
					}
				}
			} else {
				log.Fatal(fmt.Sprintf("Plugin %s is not installed, but is included in kombustion.yaml. Run `kombustion install` to fix.", manifestPlugin.Name))
			}
		}
	}
	fmt.Println(resources)
	return resources, outputs, mappings
}

func loadPlugin(manifestPlugin manifest.Plugin, pluginName string, pluginVersion string, pluginPath string, resources, outputs, mappings map[string]kombustionTypes.ParserFunc) {
	resources, outputs, mappings =
		make(map[string]kombustionTypes.ParserFunc),
		make(map[string]kombustionTypes.ParserFunc),
		make(map[string]kombustionTypes.ParserFunc)

		// TODO: Make the help messages for users much friendlier
	if !pluginExists(pluginPath) {
		fmt.Fprintf(os.Stderr, "error: invalid plugin file: %s\n", pluginPath)
		os.Exit(1)
	}

	// TODO: Check the hash of the plugin to load matches the lockfile

	// open the plug-in file, causing its package init function to run,
	// thereby registering the module
	p, err := plugin.Open(pluginPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to load plugin: %v\n", err)
	}

	// Config
	RegisterConstructor, _ := p.Lookup("Register")
	configFunc := RegisterConstructor.(func() []byte)
	config, err := loadConfig(configFunc())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to load config for plugin: %v\n", err)
	}
	fmt.Println(config)

	if configIsValid(config, pluginName, pluginVersion, false) == false {
		log.Fatal(fmt.Sprintf("Unable to load plugin: %s", pluginName))
	}

	prefix := config.Prefix

	if manifestPlugin.Alias != "" {
		prefix = manifestPlugin.Alias
	}

	// [ Resources ]----------------------------------------------------------------------------------

	resourcesConstructor, _ := p.Lookup("Resources")
	if resourcesConstructor != nil {

		resourcesFuncs := *resourcesConstructor.(*map[string]func(
			ctx map[string]interface{},
			name string,
			data string,
		) []byte)

		for key, parserFunc := range resourcesFuncs {
			pluginKey := fmt.Sprintf("%s::%s", prefix, key)
			if _, ok := resources[pluginKey]; ok { // Check for duplicates
				log.WithFields(log.Fields{
					"resource": pluginKey,
				}).Warn("duplicate resource definition for resource")
			} else {
				wrappedParserFunc := func(ctx map[string]interface{}, name string, data string) (kombustionTypes.TemplateObject, error) {
					return loadResource(parserFunc(ctx, name, data))
				}
				resources[pluginKey] = wrappedParserFunc
			}
		}
	}

	// [ Mapping ]------------------------------------------------------------------------------------

	mappingsConstructor, _ := p.Lookup("Mappings")
	if mappingsConstructor != nil {

		mappingsFuncs := *mappingsConstructor.(*map[string]func(
			ctx map[string]interface{},
			name string,
			data string,
		) []byte)

		for key, parserFunc := range mappingsFuncs {
			pluginKey := fmt.Sprintf("%s::%s", prefix, key)
			if _, ok := mappings[pluginKey]; ok { // Check for duplicates
				log.WithFields(log.Fields{
					"resource": pluginKey,
				}).Warn("duplicate resource definition for mapping")
			} else {
				wrappedParserFunc := func(ctx map[string]interface{}, name string, data string) (kombustionTypes.TemplateObject, error) {
					return loadResource(parserFunc(ctx, name, data))
				}
				mappings[pluginKey] = wrappedParserFunc
			}
		}
	}

	// [ Outputs ]------------------------------------------------------------------------------------

	outputsConstructor, _ := p.Lookup("Mapping")
	if outputsConstructor != nil {

		outputsFuncs := *outputsConstructor.(*map[string]func(
			ctx map[string]interface{},
			name string,
			data string,
		) []byte)

		for key, parserFunc := range outputsFuncs {
			pluginKey := fmt.Sprintf("%s::%s", prefix, key)
			if _, ok := outputs[pluginKey]; ok { // Check for duplicates
				log.WithFields(log.Fields{
					"resource": key,
				}).Warn("duplicate resource definition for output")
			} else {
				wrappedParserFunc := func(ctx map[string]interface{}, name string, data string) (kombustionTypes.TemplateObject, error) {
					return loadResource(parserFunc(ctx, name, data))
				}
				outputs[pluginKey] = wrappedParserFunc
			}
		}
	}

	return
}

// Helper function to ensure a plugin file exists before we attempt to load it
func pluginExists(filePath string) bool {
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		return true
	}
	return false
}

func configIsValid(config pluginTypes.Config, pluginName string, pluginVersion string, requiresAWSSession bool) (ok bool) {
	foundIssue := false
	// TODO: improve these error messages, and provide links to the docs for plugin devs
	if config.Name == "" {
		// err
		log.Fatal(fmt.Sprintf("%s did not supply a name, this plugin cannot be loaded", config.Name))
		foundIssue = true
	}
	// } else if config.Name != pluginName {
	// 	// warn
	// 	foundIssue = true
	// 	log.Fatal(fmt.Sprintf("%s name did not match the name in kombustion.yaml, this plugin cannot be loaded"))
	// }

	if config.Prefix == "" {
		log.Fatal(fmt.Sprintf("%s did not supply a prefix, this plugin cannot be loaded", config.Name))
		foundIssue = true
	}

	if config.RequiresAWSSession != requiresAWSSession {
		// Warn about the need to add the config val to the manifest file
		// foundIssue = true
	}
	if foundIssue == false {
		ok = true
	}
	return
}
