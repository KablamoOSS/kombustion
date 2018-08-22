package plugins

import (
	"fmt"
	"os"
	"plugin"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/KablamoOSS/kombustion/internal/plugins/lock"
	pluginTypes "github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
)

// LoadPlugins from a manifest and lockfile
func LoadPlugins(manifestFile *manifest.Manifest, lockFile *lock.Lock) (loadedPlugins []*PluginLoaded) {
	// Load all plugins the manifest
	for _, manifestPlugin := range manifestFile.Plugins {
		for _, plugin := range lockFile.Plugins {
			// Find the matching plugin in the lock file
			if plugin.Match(&manifestPlugin) {
				resolution := plugin.ResolveForRuntime()
				if resolution == nil {
					printer.Fatal(
						fmt.Errorf("Could not resolve plugin %s", plugin.Name),
						"Ensure plugin supports your operating system and architecture",
						"",
					)
				}
				loadedPlugins = append(
					loadedPlugins,
					loadPlugin(plugin.Name, plugin.Version, resolution.PathOnDisk, resolution.Hash, false, manifestPlugin.Alias),
				)
			} else {
				printer.Fatal(
					fmt.Errorf("Plugin `%s@%s` is not installed, but is included in kombustion.yaml", manifestPlugin.Name, manifestPlugin.Version),
					"Run `kombustion install`",
					"",
				)
			}
		}
	}

	return
}

// LoadDevPlugin loads an arbitrary plugin for plugin developers, to ease plugin development.
// Only works with a kombustion binary that was built from source
func LoadDevPlugin(pluginPath string) *PluginLoaded {
	return loadPlugin(
		"dev-loaded-plugin",
		"DEV",
		pluginPath,
		"",
		true,
		"",
	)
}

func loadPlugin(
	name string,
	version string,
	path string,
	expectedHash string,
	isDevPlugin bool,
	alias string,
) *PluginLoaded {

	loadedPlugin := PluginLoaded{}

	// TODO: Make the help messages for users much friendlier
	if !pluginExists(path) {
		if isDevPlugin {
			printer.Fatal(
				fmt.Errorf("Plugin `%s` could not be found", path),
				"Check the path you provided with --load-plugin is correct.",
				"https://www.kombustion.io/api/cli/#load-plugin",
			)
		}
		printer.Fatal(
			fmt.Errorf("Plugin `%s` is not installed, but is included in kombustion.lock", name),
			"Run `kombustion install` to fix.",
			"",
		)
	}

	loadedPlugin.InternalConfig.PathOnDisk = path

	currentHash, err := getHashOfFile(path)
	if err != nil {
		printer.Fatal(
			err,
			fmt.Sprintf("Check user has permissions to read %s", path),
			"",
		)
	}
	if !isDevPlugin && currentHash != expectedHash {
		printer.Fatal(
			fmt.Errorf("Hash of plugin %s@%s does not match lock file", name, version),
			"Reinstall the plugin by removing %s and running `kombustion install`",
			"",
		)
	}

	// Open the plugin
	p, err := plugin.Open(path)
	if err != nil {
		printer.Fatal(
			fmt.Errorf("Plugin `%s` could not be loaded, this is likely an issue with the plugin", name),
			"Try your command again, but if it fails file an issue with the plugin author.",
			"",
		)
	}

	// Config
	RegisterConstructor, _ := p.Lookup("Register")
	configFunc := RegisterConstructor.(func() []byte)
	config, err := loadConfig(configFunc())
	if err != nil {
		printer.Fatal(
			fmt.Errorf("Plugin `%s` does not have a valid config", name),
			"Try your command again, but if it fails file an issue with the plugin author.",
			"",
		)
	}

	if !configIsValid(config, name, version) {
		printer.Fatal(
			fmt.Errorf("Plugin `%s` does not have a valid config", name),
			"Contact the plugin author.",
			"",
		)
	}

	loadedPlugin.Config = config

	loadedPlugin.InternalConfig.Prefix = config.Prefix

	if alias != "" {
		loadedPlugin.InternalConfig.Prefix = alias
	}

	// Load Parsers
	parserConstructor, _ := p.Lookup("Parsers")
	if parserConstructor != nil {
		loadedPlugin.Parsers = parserConstructor.(*map[string]func(string, string) []byte)
	}

	return &loadedPlugin
}

// Helper function to ensure a plugin file exists before we attempt to load it
func pluginExists(filePath string) bool {
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		return true
	}
	return false
}

func configIsValid(config pluginTypes.Config, pluginName string, pluginVersion string) (ok bool) {
	// TODO: improve these error messages, and provide links to the docs for plugin devs

	// NOTE: Even though `printer.Fatal` normally terminates the program,
	// that's bypassed in testing so the `return false` is important.
	if config.Name == "" {
		printer.Fatal(
			fmt.Errorf("Plugin `%s` did not supply a name, this plugin cannot be loaded", pluginName),
			"Try your command again, but if it fails file an issue with the plugin author.",
			"",
		)
		return false
	}

	switch config.Prefix {
	case "":
		printer.Fatal(
			fmt.Errorf("Plugin `%s` did not supply a prefix, this plugin cannot be loaded", pluginName),
			"Try your command again, but if it fails file an issue with the plugin author.",
			"",
		)
		return false

	case "AWS":
		printer.Fatal(
			fmt.Errorf("Plugin `%s` tried to use 'AWS' as prefix, this plugin cannot be loaded", pluginName),
			"'AWS' is a restricted prefix, and cannt be used by a plugin. This is an issue with the plugin.",
			"",
		)
		return false

	case "Custom":
		printer.Fatal(
			fmt.Errorf("Plugin `%s` tried to use 'Custom' as prefix, this plugin cannot be loaded", pluginName),
			"'Custom' is a restricted prefix, and cannt be used by a plugin. This is an issue with the plugin.",
			"",
		)
		return false

	case "Kombustion":
		printer.Fatal(
			fmt.Errorf("Plugin `%s` tried to use 'Kombustion' as prefix, this plugin cannot be loaded", pluginName),
			"'Kombustion' is a restricted prefix, and cannt be used by a plugin. This is an issue with the plugin.",
			"",
		)
		return false
	}

	return true
}
