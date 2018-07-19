package plugins

import (
	"fmt"

	printer "github.com/KablamoOSS/go-cli-printer"
	kombustionTypes "github.com/KablamoOSS/kombustion/types"
)

// ExtractParsersFromPlugins and ensure there are no clashes for plugin resource names
func ExtractParsersFromPlugins(
	loadedPlugins []*PluginLoaded,
) {
	var parsers map[string]kombustionTypes.ParserFunc
	for _, plugin := range loadedPlugins {
		if *plugin.Parsers != nil {
			for key, parserFunc := range *plugin.Parsers {
				pluginKey := fmt.Sprintf("%s::%s", plugin.InternalConfig.Prefix, key)
				if _, ok := parsers[pluginKey]; ok { // Check for duplicates
					printer.Fatal(
						fmt.Errorf("Plugin `%s` tried to load resource `%s` but it already exists", plugin.Config.Name, pluginKey),
						fmt.Sprintf(
							"You can add a `Alias` to this plugin in kombustion.yaml to resolve this.",
						),
						"https://www.kombustion.io/api/manifest/#alias-optional",
					)
				} else {
					// This function has the same type as types.ParserFunc, but we need to define the type
					// as follows, becuase it's crossing the host/plugin boundary
					wrappedParserFunc := func(
						name string,
						data string,
					) (
						source string,
						sourceResource string,
						conditions kombustionTypes.TemplateObject,
						metadata kombustionTypes.TemplateObject,
						mappings kombustionTypes.TemplateObject,
						outputs kombustionTypes.TemplateObject,
						parameters kombustionTypes.TemplateObject,
						resources kombustionTypes.TemplateObject,
						errors []error,
					) {

						source = plugin.Config.Name

						conditions,
							metadata,
							mappings,
							outputs,
							parameters,
							resources,
							errors = unmarshallParser(parserFunc(name, data))

						// hasErrors := false

						// for _, err := range errs {
						// 	if err != nil {
						// 		hasErrors = true
						// 		printer.Error(
						// 			err,
						// 			fmt.Sprintf(
						// 				"\n   ├─ Name:    %s\n   ├─ Source: %s\n   └─ Type:    %s",
						// 				name,
						// 				plugin.Config.Name,
						// 				pluginKey,
						// 			),
						// 			"",
						// 		)
						// 	}
						// }

						// if hasErrors {
						// 	return parsers, fmt.Errorf("There were errors parsing %s", name)
						// }

						return
					} // </wrappedParserFunc>

					// Add the parser func to our parsers
					parsers[pluginKey] = wrappedParserFunc
				}
			}
		}
	}
	return
}
