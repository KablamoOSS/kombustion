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
				Name:               "Name",
				Version:            "0.1.0",
				Prefix:             "Test",
				RequiresAWSSession: false,
				Help:               apiTypes.Help{},
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
			ctx map[string]interface{},
			name string,
			data string,
		) types.TemplateObject
		output types.TemplateObject
	}{
		{
			input: func(
				ctx map[string]interface{},
				name string,
				data string,
			) types.TemplateObject {
				return types.TemplateObject{
					"Key":  "Value",
					"Name": name,
					"Data": data,
				}
			},
			output: types.TemplateObject{
				"Key":  "Value",
				"Name": "TestName",
				"Data": "TestData",
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)

		ctx := types.TemplateObject{
			"Key": "Value",
		}
		name := "TestName"
		data := "TestData"
		testFunc := RegisterResource(test.input)

		testOutput := testFunc(ctx, name, data)

		var testResource types.TemplateObject

		err := msgpack.Unmarshal(testOutput, &testResource)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(testResource, test.output, fmt.Sprintf("Test %d", i))
	}
}

func TestRegisterOutput(t *testing.T) {
	tests := []struct {
		input func(
			ctx map[string]interface{},
			name string,
			data string,
		) types.TemplateObject
		output types.TemplateObject
	}{
		{
			input: func(
				ctx map[string]interface{},
				name string,
				data string,
			) types.TemplateObject {
				return types.TemplateObject{
					"Key":  "Value",
					"Name": name,
					"Data": data,
				}
			},
			output: types.TemplateObject{
				"Key":  "Value",
				"Name": "TestName",
				"Data": "TestData",
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)

		ctx := types.TemplateObject{
			"Key": "Value",
		}
		name := "TestName"
		data := "TestData"
		testFunc := RegisterOutput(test.input)

		testOutput := testFunc(ctx, name, data)

		var testMapping types.TemplateObject

		err := msgpack.Unmarshal(testOutput, &testMapping)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(testMapping, test.output, fmt.Sprintf("Test %d", i))
	}
}

func TestRegisterMapping(t *testing.T) {
	tests := []struct {
		input func(
			ctx map[string]interface{},
			name string,
			data string,
		) types.TemplateObject
		output types.TemplateObject
	}{
		{
			input: func(
				ctx map[string]interface{},
				name string,
				data string,
			) types.TemplateObject {
				return types.TemplateObject{
					"Key":  "Value",
					"Name": name,
					"Data": data,
				}
			},
			output: types.TemplateObject{
				"Key":  "Value",
				"Name": "TestName",
				"Data": "TestData",
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)

		ctx := types.TemplateObject{
			"Key": "Value",
		}
		name := "TestName"
		data := "TestData"
		testFunc := RegisterMapping(test.input)

		testOutput := testFunc(ctx, name, data)

		var testMapping types.TemplateObject

		err := msgpack.Unmarshal(testOutput, &testMapping)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(testMapping, test.output, fmt.Sprintf("Test %d", i))
	}
}
