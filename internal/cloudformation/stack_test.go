package cloudformation

import (
	"fmt"
	"testing"

	"github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/stretchr/testify/assert"
)

func TestGetStackName(t *testing.T) {
	type input struct {
		manifestFile  *manifest.Manifest
		fileName      string
		environment   string
		stackNameFlag string
	}
	tests := []struct {
		input  input
		output string
	}{
		{
			input: input{
				manifestFile: &manifest.Manifest{
					Name: "KombustionDemo",
				},
				fileName:      "stacks/MyDemoStack.yaml",
				environment:   "development",
				stackNameFlag: "",
			},
			output: "KombustionDemo-MyDemoStack-development",
		},
		{
			input: input{
				manifestFile:  &manifest.Manifest{},
				fileName:      "stacks/MyDemoStack.yaml",
				environment:   "development",
				stackNameFlag: "MyCustomStackName",
			},
			output: "MyCustomStackName",
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput := GetStackName(
			test.input.manifestFile,
			test.input.fileName,
			test.input.environment,
			test.input.stackNameFlag,
		)

		assert.Equal(test.output, testOutput, fmt.Sprintf("Test %d", i))
	}
}

func TestCleanStackName(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		// TODO: Add more test cases
		{
			input:  "KombustionDemo-/MyDemoStack-development",
			output: "KombustionDemo-MyDemoStack-development",
		},
		{
			input:  "';[},.;'][~!@#$%^&*()_+/KombustionDemo-MyDemoStack-development",
			output: "KombustionDemo-MyDemoStack-development",
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput := cleanStackName(test.input)

		assert.Equal(test.output, testOutput, fmt.Sprintf("Test %d: %s", i, test.input))
	}
}
