package api

import (
	"log"

	pluginTypes "github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
	kombustionTypes "github.com/KablamoOSS/kombustion/types"

	"github.com/vmihailenco/msgpack"
)

func marshallConfig(config pluginTypes.Config) (blob []byte) {
	blob, err := msgpack.Marshal(&config)
	if err != nil {
		log.Fatal("Config marshalling err:", err)
	}
	return
}

func marshallParserResult(
	conditions kombustionTypes.TemplateObject,
	metadata kombustionTypes.TemplateObject,
	mappings kombustionTypes.TemplateObject,
	outputs kombustionTypes.TemplateObject,
	parameters kombustionTypes.TemplateObject,
	resources kombustionTypes.TemplateObject,
	transform kombustionTypes.TemplateObject,
	errors []error,

) (blob []byte) {
	result := pluginTypes.PluginParserResult{
		Conditions: conditions,
		Metadata:   metadata,
		Mappings:   mappings,
		Outputs:    outputs,
		Parameters: parameters,
		Resources:  resources,
		Transform:  transform,
		Errors:     errors,
	}
	blob, err := msgpack.Marshal(&result)
	if err != nil {
		log.Fatal("Resource marshalling err:", err)
	}
	return
}
