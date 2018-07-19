package plugins

import (
	"fmt"
	"os"
	"plugin"
	"runtime"

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
			if manifestPlugin.Name == plugin.Name && manifestPlugin.Version == plugin.Version {
				for _, resolved := range plugin.Resolved {
					// Find the right plugin for the current OS/Arch
					if runtime.GOOS == resolved.OperatingSystem &&
						runtime.GOARCH == resolved.Architecture {
						loadedPlugins = append(
							loadedPlugins,
							loadPlugin(manifestPlugin, plugin.Name, plugin.Version, resolved.PathOnDisk, false),
						)
					}
				}
			} else {
				printer.Fatal(
					fmt.Errorf("Plugin `%s` is not installed, but is included in kombustion.yaml", manifestPlugin.Name),
					fmt.Sprintf(
						"Run `kombustion install` to fix.",
					),
					"",
				)
			}
		}
	}

	return
}

// LoadDevPlugin loads an arbitrary plugin for plugin developers, to ease plugin development.
// Only works with a kombustion binary that was built from source
func LoadDevPlugin(
	pluginPath string,
) *PluginLoaded {
	return loadPlugin(
		manifest.Plugin{},
		"dev-loaded-plugin",
		"DEV",
		pluginPath,
		true,
	)
}

func loadPlugin(
	manifestPlugin manifest.Plugin,
	pluginName string,
	pluginVersion string,
	pluginPath string,
	isDevPlugin bool,
) *PluginLoaded {

	loadedPlugin := PluginLoaded{}

	// TODO: Make the help messages for users much friendlier
	if !pluginExists(pluginPath) {
		if isDevPlugin {
			printer.Fatal(
				fmt.Errorf("Plugin `%s` could not be found", pluginPath),
				fmt.Sprintf(
					"Check the path you provided with --load-plugin is correct.",
				),
				"https://www.kombustion.io/api/cli/#load-plugin",
			)
		}
		printer.Fatal(
			fmt.Errorf("Plugin `%s` is not installed, but is included in kombustion.lock", manifestPlugin.Name),
			fmt.Sprintf(
				"Run `kombustion install` to fix.",
			),
			"",
		)
	}

	loadedPlugin.InternalConfig.PathOnDisk = pluginPath

	// TODO: Check the hash of the plugin to load matches the lockfile

	// Open the plugin
	p, err := plugin.Open(pluginPath)
	if err != nil {
		printer.Fatal(
			fmt.Errorf("Plugin `%s` could not be loaded, this is likely an issue with the plugin", manifestPlugin.Name),
			fmt.Sprintf(
				"Try your command again, but if it fails file an issue with the plugin author.",
			),
			"",
		)
	}

	// Config
	RegisterConstructor, _ := p.Lookup("Register")
	configFunc := RegisterConstructor.(func() []byte)
	config, err := loadConfig(configFunc())
	if err != nil {
		printer.Fatal(
			fmt.Errorf("Plugin `%s` does not have a valid config", pluginName),
			"Try your command again, but if it fails file an issue with the plugin author.",
			"",
		)
	}

	if configIsValid(config, pluginName, pluginVersion) == false {
		printer.Fatal(
			fmt.Errorf("Plugin `%s` does not have a valid config", pluginName),
			"Contact the plugin author.",
			"",
		)
	}

	loadedPlugin.Config = config

	loadedPlugin.InternalConfig.Prefix = config.Prefix

	if manifestPlugin.Alias != "" {
		loadedPlugin.InternalConfig.Prefix = manifestPlugin.Alias
	}

	// Load Parsers
	parserConstructor, _ := p.Lookup("Parsers")
	if parserConstructor != nil {
		loadedPlugin.Parsers = parserConstructor.(*map[string]func(
			name string,
			data string,
		) []byte)
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
	foundIssue := false
	// TODO: improve these error messages, and provide links to the docs for plugin devs
	if config.Name == "" {
		printer.Fatal(
			fmt.Errorf("Plugin `%s` did not supply a name, this plugin cannot be loaded", pluginName),
			"Try your command again, but if it fails file an issue with the plugin author.",
			"",
		)

		foundIssue = true
	}

	if config.Prefix == "" {
		printer.Fatal(
			fmt.Errorf("Plugin `%s` did not supply a prefix, this plugin cannot be loaded", pluginName),
			"Try your command again, but if it fails file an issue with the plugin author.",
			"",
		)
		foundIssue = true
	}

	if config.Prefix == "AWS" {
		printer.Fatal(
			fmt.Errorf("Plugin `%s` tried to use 'AWS' as prefix, this plugin cannot be loaded", pluginName),
			"'AWS' is a restricted prefix, and cannt be used by a plugin. This is an issue with the plugin.",
			"",
		)
		foundIssue = true
	}

	if config.Prefix == "Custom" {
		printer.Fatal(
			fmt.Errorf("Plugin `%s` tried to use 'Custom' as prefix, this plugin cannot be loaded", pluginName),
			"'Custom' is a restricted prefix, and cannt be used by a plugin. This is an issue with the plugin.",
			"",
		)
		foundIssue = true
	}

	if config.Prefix == "Kombustion" {
		printer.Fatal(
			fmt.Errorf("Plugin `%s` tried to use 'Kombustion' as prefix, this plugin cannot be loaded", pluginName),
			"'Kombustion' is a restricted prefix, and cannt be used by a plugin. This is an issue with the plugin.",
			"",
		)
		foundIssue = true
	}

	if foundIssue == false {
		ok = true
	}
	return
}
