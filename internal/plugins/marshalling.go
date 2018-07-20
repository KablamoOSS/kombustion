package plugins

import (
	pluginTypes "github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
	kombustionTypes "github.com/KablamoOSS/kombustion/types"

	"github.com/vmihailenco/msgpack"
)

func loadConfig(blob []byte) (config pluginTypes.Config, err error) {
	err = msgpack.Unmarshal(blob, &config)
	if err != nil {
		panic(err)
	}
	return
}

// unmarshallParser unmarshalls the result of the parser function
func unmarshallParser(blob []byte) (
	conditions kombustionTypes.TemplateObject,
	metadata kombustionTypes.TemplateObject,
	mappings kombustionTypes.TemplateObject,
	outputs kombustionTypes.TemplateObject,
	parameters kombustionTypes.TemplateObject,
	resources kombustionTypes.TemplateObject,
	transform kombustionTypes.TemplateObject,
	errors []error,
) {
	var pluginResult pluginTypes.PluginParserResult
	err := msgpack.Unmarshal(blob, &pluginResult)
	if err != nil {
		panic(err)
	}

	if pluginResult.Conditions != nil {
		conditions = pluginResult.Conditions
	}

	if pluginResult.Metadata != nil {
		metadata = pluginResult.Metadata
	}

	if pluginResult.Mappings != nil {
		mappings = pluginResult.Mappings
	}

	if pluginResult.Outputs != nil {
		outputs = pluginResult.Outputs
	}

	if pluginResult.Parameters != nil {
		parameters = pluginResult.Parameters
	}

	if pluginResult.Resources != nil {
		resources = pluginResult.Resources
	}

	if pluginResult.Transform != nil {
		transform = pluginResult.Transform
	}

	if pluginResult.Errors != nil {
		errors = pluginResult.Errors
	}

	return
}
