package api

import (
	"fmt"
	"testing"

	apiTypes "github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
	"github.com/KablamoOSS/kombustion/types"
	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack"
)

func TestRegisterPlugin(t *testing.T) {
	tests := []struct {
		input apiTypes.Config
	}{
		{
			input: apiTypes.Config{
				Name:    "Name",
				Version: "0.1.0",
				Prefix:  "Test",
				Help:    apiTypes.Help{},
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)

		testOutput := RegisterPlugin(test.input)

		var testConfig apiTypes.Config
		err := msgpack.Unmarshal(testOutput, &testConfig)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(testConfig, test.input, fmt.Sprintf("Test %d", i))
	}
}
func TestRegisterResource(t *testing.T) {
	tests := []struct {
		input func(
			name string,
			data string,
		) (types.TemplateObject, []error)
		output types.TemplateObject
	}{
		{
			input: func(
				name string,
				data string,
			) (types.TemplateObject, []error) {
				return types.TemplateObject{
					"Name": name,
					"Data": data,
				}, []error{nil}
			},
			output: types.TemplateObject{
				"Name": "TestName",
				"Data": "TestData",
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)

		name := "TestName"
		data := "TestData"
		testFunc := RegisterResource(test.input)

		testOutput := testFunc(name, data)

		var pluginResult apiTypes.PluginResult
		err := msgpack.Unmarshal(testOutput, &pluginResult)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(pluginResult.Data, test.output, fmt.Sprintf("Test %d", i))
	}
}

func TestRegisterOutput(t *testing.T) {
	tests := []struct {
		input func(
			name string,
			data string,
		) (types.TemplateObject, []error)
		output types.TemplateObject
	}{
		{
			input: func(
				name string,
				data string,
			) (types.TemplateObject, []error) {
				return types.TemplateObject{
					"Name": name,
					"Data": data,
				}, []error{nil}
			},
			output: types.TemplateObject{
				"Name": "TestName",
				"Data": "TestData",
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)

		name := "TestName"
		data := "TestData"
		testFunc := RegisterOutput(test.input)

		testOutput := testFunc(name, data)

		var pluginResult apiTypes.PluginResult
		err := msgpack.Unmarshal(testOutput, &pluginResult)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(pluginResult.Data, test.output, fmt.Sprintf("Test %d", i))
	}
}

func TestRegisterMapping(t *testing.T) {
	tests := []struct {
		input func(
			name string,
			data string,
		) (types.TemplateObject, []error)
		output types.TemplateObject
	}{
		{
			input: func(
				name string,
				data string,
			) (types.TemplateObject, []error) {
				return types.TemplateObject{
					"Name": name,
					"Data": data,
				}, []error{nil}
			},
			output: types.TemplateObject{
				"Name": "TestName",
				"Data": "TestData",
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)

		name := "TestName"
		data := "TestData"
		testFunc := RegisterMapping(test.input)

		testOutput := testFunc(name, data)

		var pluginResult apiTypes.PluginResult
		err := msgpack.Unmarshal(testOutput, &pluginResult)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(pluginResult.Data, test.output, fmt.Sprintf("Test %d", i))
	}
}
