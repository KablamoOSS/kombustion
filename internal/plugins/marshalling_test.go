package plugins

import (
	"fmt"
	"testing"

	pluginTypes "github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
	kombustionTypes "github.com/KablamoOSS/kombustion/types"
	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		input  []byte
		output pluginTypes.Config
		throws bool
	}{
		{
			input: func() []byte {
				blob, err := msgpack.Marshal(&pluginTypes.Config{
					Name:    "example-plugin",
					Version: "0.1.0",
					Prefix:  "Test",
				})

				if err != nil {
					t.Fatalf("Config marshalling err")
				}

				return blob
			}(),
			output: pluginTypes.Config{
				Name:    "example-plugin",
				Version: "0.1.0",
				Prefix:  "Test",
			},
			throws: false,
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput, err := loadConfig(
			test.input,
		)
		if test.throws {
			assert.NotNil(err)
		} else {
			assert.Nil(err)
			assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d:", i))
		}
	}
}

func TestLoadParser(t *testing.T) {
	type Output struct {
		conditions kombustionTypes.TemplateObject
		metadata   kombustionTypes.TemplateObject
		mappings   kombustionTypes.TemplateObject
		outputs    kombustionTypes.TemplateObject
		parameters kombustionTypes.TemplateObject
		resources  kombustionTypes.TemplateObject
		transform  kombustionTypes.TemplateObject
		errors     []error
	}
	tests := []struct {
		input  []byte
		output Output
	}{
		{
			input: func() []byte {
				blob, err := msgpack.Marshal(
					&pluginTypes.PluginParserResult{
						Conditions: kombustionTypes.TemplateObject{},
						Metadata:   kombustionTypes.TemplateObject{},
						Mappings:   kombustionTypes.TemplateObject{},
						Outputs:    kombustionTypes.TemplateObject{},
						Parameters: kombustionTypes.TemplateObject{},
						Resources:  kombustionTypes.TemplateObject{},
						Transform:  kombustionTypes.TemplateObject{},
						Errors:     []error{nil},
					})

				if err != nil {
					t.Fatalf("Config marshalling err")
				}

				return blob
			}(),
			output: Output{
				conditions: kombustionTypes.TemplateObject{},
				metadata:   kombustionTypes.TemplateObject{},
				mappings:   kombustionTypes.TemplateObject{},
				outputs:    kombustionTypes.TemplateObject{},
				parameters: kombustionTypes.TemplateObject{},
				resources:  kombustionTypes.TemplateObject{},
				transform:  kombustionTypes.TemplateObject{},
				errors:     []error{nil},
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		conditions,
			metadata,
			mappings,
			outputs,
			parameters,
			resources,
			transform,
			testErrs := unmarshallParser(
			test.input,
		)

		assert.Equal(testErrs, test.output.errors, fmt.Sprintf("Test %d:", i))
		assert.Equal(conditions, test.output.conditions, fmt.Sprintf("Test %d:", i))
		assert.Equal(metadata, test.output.metadata, fmt.Sprintf("Test %d:", i))
		assert.Equal(mappings, test.output.mappings, fmt.Sprintf("Test %d:", i))
		assert.Equal(outputs, test.output.outputs, fmt.Sprintf("Test %d:", i))
		assert.Equal(parameters, test.output.parameters, fmt.Sprintf("Test %d:", i))
		assert.Equal(resources, test.output.resources, fmt.Sprintf("Test %d:", i))
		assert.Equal(transform, test.output.transform, fmt.Sprintf("Test %d:", i))

	}

}
