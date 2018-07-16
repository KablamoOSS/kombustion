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

func TestLoadResource(t *testing.T) {
	tests := []struct {
		input  []byte
		output kombustionTypes.TemplateObject
		errs   []error
	}{
		{
			input: func() []byte {
				blob, err := msgpack.Marshal(&pluginTypes.PluginResult{
					Data: kombustionTypes.TemplateObject{
						"Name": "example-plugin",
					},
					Errors: []error{nil},
				})

				if err != nil {
					t.Fatalf("Config marshalling err")
				}

				return blob
			}(),
			output: kombustionTypes.TemplateObject{
				"Name": "example-plugin",
			},
			errs: []error{nil},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput, testErrs := loadResource(
			test.input,
		)
		assert.Equal(testErrs, test.errs, fmt.Sprintf("Test %d:", i))
		assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d:", i))
	}

}

func TestLoadMapping(t *testing.T) {
	tests := []struct {
		input  []byte
		output kombustionTypes.TemplateObject
		errs   []error
	}{
		{
			input: func() []byte {
				blob, err := msgpack.Marshal(&pluginTypes.PluginResult{
					Data: kombustionTypes.TemplateObject{
						"Name": "example-plugin",
					},
					Errors: []error{nil},
				})

				if err != nil {
					t.Fatalf("Config marshalling err")
				}

				return blob
			}(),
			output: kombustionTypes.TemplateObject{
				"Name": "example-plugin",
			},
			errs: []error{nil},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput, testErrs := loadMapping(
			test.input,
		)
		assert.Equal(testErrs, test.errs, fmt.Sprintf("Test %d:", i))
		assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d:", i))
	}
}

func TestLoadOutput(t *testing.T) {
	tests := []struct {
		input  []byte
		output kombustionTypes.TemplateObject
		errs   []error
	}{
		{
			input: func() []byte {
				blob, err := msgpack.Marshal(&pluginTypes.PluginResult{
					Data: kombustionTypes.TemplateObject{
						"Name": "example-plugin",
					},
					Errors: []error{nil},
				})

				if err != nil {
					t.Fatalf("Config marshalling err")
				}

				return blob
			}(),
			output: kombustionTypes.TemplateObject{
				"Name": "example-plugin",
			},
			errs: []error{nil},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput, testErrs := loadOutput(
			test.input,
		)
		assert.Equal(testErrs, test.errs, fmt.Sprintf("Test %d:", i))
		assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d:", i))
	}
}
