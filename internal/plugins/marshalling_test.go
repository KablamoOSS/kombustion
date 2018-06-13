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
					Name:               "example-plugin",
					Version:            "0.1.0",
					Prefix:             "Test",
					RequiresAWSSession: false,
				})

				if err != nil {
					t.Fatalf("Config marshalling err")
				}

				return blob
			}(),
			output: pluginTypes.Config{
				Name:               "example-plugin",
				Version:            "0.1.0",
				Prefix:             "Test",
				RequiresAWSSession: false,
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
		throws bool
	}{
		{
			input: func() []byte {
				blob, err := msgpack.Marshal(&kombustionTypes.TemplateObject{
					"Name": "example-plugin",
				})

				if err != nil {
					t.Fatalf("Config marshalling err")
				}

				return blob
			}(),
			output: kombustionTypes.TemplateObject{
				"Name": "example-plugin",
			},
			throws: false,
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput, err := loadResource(
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

func TestLoadMapping(t *testing.T) {
	tests := []struct {
		input  []byte
		output kombustionTypes.TemplateObject
		throws bool
	}{
		{
			input: func() []byte {
				blob, err := msgpack.Marshal(&kombustionTypes.TemplateObject{
					"Name": "example-plugin",
				})

				if err != nil {
					t.Fatalf("Config marshalling err")
				}

				return blob
			}(),
			output: kombustionTypes.TemplateObject{
				"Name": "example-plugin",
			},
			throws: false,
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput, err := loadMapping(
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

func TestLoadOutput(t *testing.T) {
	tests := []struct {
		input  []byte
		output kombustionTypes.TemplateObject
		throws bool
	}{
		{
			input: func() []byte {
				blob, err := msgpack.Marshal(&kombustionTypes.TemplateObject{
					"Name": "example-plugin",
				})

				if err != nil {
					t.Fatalf("Config marshalling err")
				}

				return blob
			}(),
			output: kombustionTypes.TemplateObject{
				"Name": "example-plugin",
			},
			throws: false,
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput, err := loadOutput(
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
