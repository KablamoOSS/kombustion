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
							"You can add a `Alias` to this plugin in kombustion.yaml to resolve this.",
						),
						"https://www.kombustion.io/api/manifest/#alias-optional",
					)
				} else {
					wrappedParserFunc := func(
						name string, data string,
					) (kombustionTypes.TemplateObject, error) {
						resources, errs := loadResource(parserFunc(name, data))
						hasErrors := false
						for _, err := range errs {
							if err != nil {
								hasErrors = true
								printer.Error(
									err,
									fmt.Sprintf(
										"\n   ├─ Name:    %s\n   ├─ Plugin:  %s\n   └─ Type:    %s",
										name,
										plugin.Config.Name,
										pluginKey,
									),
									"",
								)
							}
						}

						if hasErrors {
							return resources, fmt.Errorf("There were errors parsing %s", name)
						}
						return resources, nil
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
							"You can add a `Alias` to this plugin in kombustion.yaml to resolve this.",
						),
						"https://www.kombustion.io/api/manifest/#alias-optional",
					)
				} else {
					wrappedParserFunc := func(
						name string, data string,
					) (kombustionTypes.TemplateObject, error) {
						mapping, errs := loadMapping(parserFunc(name, data))

						hasErrors := false
						for _, err := range errs {
							if err != nil {
								hasErrors = true

								printer.Error(
									err,
									fmt.Sprintf(
										"\n   ├─ Name:    %s\n   ├─ Plugin:  %s\n   └─ Type:    %s",
										name,
										plugin.Config.Name,
										pluginKey,
									),
									"",
								)
							}
						}

						if hasErrors {
							return mapping, fmt.Errorf("There were errors parsing %s", name)
						}
						return mapping, nil
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
							"You can add a `Alias` to this plugin in kombustion.yaml to resolve this.",
						),
						"https://www.kombustion.io/api/manifest/#alias-optional",
					)
				} else {
					wrappedParserFunc := func(
						name string, data string,
					) (kombustionTypes.TemplateObject, error) {
						outputs, errs := loadOutput(parserFunc(name, data))

						hasErrors := false
						for _, err := range errs {
							if err != nil {
								hasErrors = true

								printer.Error(
									err,
									fmt.Sprintf(
										"\n   ├─ Name:    %s\n   ├─ Plugin:  %s\n   └─ Type:    %s",
										name,
										plugin.Config.Name,
										pluginKey,
									),
									"",
								)
							}
						}

						if hasErrors {
							return outputs, fmt.Errorf("There were errors parsing %s", name)
						}
						return outputs, nil
					}
					(*outputs)[pluginKey] = wrappedParserFunc
				}
			}
		}
	}

	return
}
