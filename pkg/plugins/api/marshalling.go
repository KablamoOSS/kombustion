package api

import (
	"fmt"
	"log"

	"github.com/KablamoOSS/yaml"

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
		Conditions: cleanResult(conditions),
		Metadata:   cleanResult(metadata),
		Mappings:   cleanResult(mappings),
		Outputs:    cleanResult(outputs),
		Parameters: cleanResult(parameters),
		Resources:  cleanResult(resources),
		Transform:  cleanResult(transform),
		Errors:     errors,
	}
	blob, err := msgpack.Marshal(&result)
	if err != nil {
		log.Fatal("Resource marshalling err:", err)
	}
	return
}

// Clean the templateObject by parsing to YAML and back
// This needs to happen before being marshalled to binary (msgpack)
// because the tags on the struct's are not stored in msgpack,
// therefore the omitempty directive is lost
//
// This prevents null values being output
func cleanResult(objects kombustionTypes.TemplateObject) (result kombustionTypes.TemplateObject) {
	result = make(kombustionTypes.TemplateObject)

	for k, v := range objects {
		// We need to check the value is not empty, to prevent a nil pointer in the yaml.Marshal
		// it will be empty when the user leaves a key blank in their template
		if v != nil {
			obj, err := yaml.Marshal(v)
			if err == nil {
				var tempObject kombustionTypes.TemplateObject
				err = yaml.Unmarshal(obj, &tempObject)
				if err == nil {
					result[k] = tempObject
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		}
	}

	return
}
