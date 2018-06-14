package plugins

import (
	"fmt"

	printer "github.com/KablamoOSS/go-cli-printer"
	kombustionTypes "github.com/KablamoOSS/kombustion/types"
)

// ExtractResourcesFromPlugins and ensure there are no clashes for plugin resource names
func ExtractResourcesFromPlugins(
	loadedPlugins []*PluginLoaded,
	resources *map[string]kombustionTypes.ParserFunc,
) {
	for _, plugin := range loadedPlugins {
		if *plugin.Resources != nil {
			for key, parserFunc := range *plugin.Resources {
				pluginKey := fmt.Sprintf("%s::%s", plugin.InternalConfig.Prefix, key)
				if _, ok := (*resources)[pluginKey]; ok { // Check for duplicates
					printer.Fatal(
						fmt.Errorf("Plugin `%s` tried to load resource `%s` but it already exists", plugin.Config.Name, pluginKey),
						fmt.Sprintf(
							"You can add a `prefix` to this plugin in kombustion.yaml to resolve this.",
						),
						"",
					)
				} else {
					wrappedParserFunc := func(
						name string, data string,
					) (
						kombustionTypes.TemplateObject,
						error,
					) {
						return loadResource(parserFunc(name, data))
					}
					(*resources)[pluginKey] = wrappedParserFunc
				}
			}
		}
	}
	return
}

// ExtractMappingsFromPlugins and ensure there are no clashes for plugin resource names
func ExtractMappingsFromPlugins(
	loadedPlugins []*PluginLoaded,
	mappings *map[string]kombustionTypes.ParserFunc,
) {
	for _, plugin := range loadedPlugins {
		if *plugin.Mappings != nil {
			for key, parserFunc := range *plugin.Mappings {
				pluginKey := fmt.Sprintf("%s::%s", plugin.InternalConfig.Prefix, key)
				if _, ok := (*mappings)[pluginKey]; ok { // Check for duplicates
					printer.Fatal(
						fmt.Errorf("Plugin `%s` tried to load mapping `%s` but it already exists", plugin.Config.Name, pluginKey),
						fmt.Sprintf(
							"You can add a `prefix` to this plugin in kombustion.yaml to resolve this.",
						),
						"",
					)
				} else {
					wrappedParserFunc := func(
						name string, data string,
					) (
						kombustionTypes.TemplateObject,
						error,
					) {
						return loadResource(parserFunc(name, data))
					}
					(*mappings)[pluginKey] = wrappedParserFunc
				}
			}
		}
	}

	return
}

// ExtractOutputsFromPlugins and ensure there are no clashes for plugin resource names
func ExtractOutputsFromPlugins(
	loadedPlugins []*PluginLoaded,
	outputs *map[string]kombustionTypes.ParserFunc,
) {
	for _, plugin := range loadedPlugins {
		if *plugin.Outputs != nil {
			for key, parserFunc := range *plugin.Outputs {
				pluginKey := fmt.Sprintf("%s::%s", plugin.InternalConfig.Prefix, key)
				if _, ok := (*outputs)[pluginKey]; ok { // Check for duplicates

					printer.Fatal(
						fmt.Errorf("Plugin `%s` tried to load output `%s` but it already exists", plugin.Config.Name, pluginKey),
						fmt.Sprintf(
							"You can add a `prefix` to this plugin in kombustion.yaml to resolve this.",
						),
						"",
					)
				} else {
					wrappedParserFunc := func(
						name string, data string,
					) (
						kombustionTypes.TemplateObject,
						error,
					) {
						return loadResource(parserFunc(name, data))
					}
					(*outputs)[pluginKey] = wrappedParserFunc
				}
			}
		}
	}

	return
}
