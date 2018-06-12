package plugins

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOSArchFromFilename(t *testing.T) {
	type input struct {
		pluginName string
		fileName   string
	}
	type output struct {
		operatingSystem string
		architecture    string
		valid           bool
	}
	tests := []struct {
		input  input
		output output
	}{
		{
			input: input{
				pluginName: "kombustion-plugin-serverless",
				fileName:   "kombustion-plugin-serverless-darwin-10.6-amd64.tgz",
			},
			output: output{
				operatingSystem: "darwin",
				architecture:    "amd64",
				valid:           true,
			},
		},
		{
			input: input{
				pluginName: "kombustion-plugin-serverless",
				fileName:   "kombustion-plugin-serverless-linux-386.tgz",
			},
			output: output{
				operatingSystem: "linux",
				architecture:    "386",
				valid:           true,
			},
		},
		{
			input: input{
				pluginName: "kombustion-plugin-serverless",
				fileName:   "kombustion-plugin-serverless-linux-amd64.tgz",
			},
			output: output{
				operatingSystem: "linux",
				architecture:    "amd64",
				valid:           true,
			},
		},
		{
			input: input{
				pluginName: "kombustion-plugin-serverless",
				fileName:   "kombustion-plugin-serverless-linux-arm64.tgz",
			},
			output: output{
				operatingSystem: "linux",
				architecture:    "arm64",
				valid:           true,
			},
		},
		{
			input: input{
				pluginName: "invalid-plugin-name",
				fileName:   "kombustion-plugin-serverless-linux-arm64.tgz",
			},
			output: output{
				operatingSystem: "",
				architecture:    "",
				valid:           false,
			},
		},
		{
			input: input{
				pluginName: "",
				fileName:   "",
			},
			output: output{
				operatingSystem: "",
				architecture:    "",
				valid:           false,
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		operatingSystem, architecture, valid := getOSArchFromFilename(
			test.input.pluginName,
			test.input.fileName,
		)
		testOutput := output{
			operatingSystem: operatingSystem,
			architecture:    architecture,
			valid:           valid,
		}
		assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d: %s", i, test.input))
	}
}

func TestCheckValidOS(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{
			input:  "darwin",
			output: true,
		},
		{
			input:  "freebsd",
			output: true,
		},

		{
			input:  "linux",
			output: true,
		},

		{
			input:  "fail",
			output: false,
		},

		{
			input:  "fail-123",
			output: false,
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput := checkValidOS(
			test.input,
		)
		assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d: %s", i, test.input))
	}
}

func TestCheckValidArch(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{
			input:  "386",
			output: true,
		},

		{
			input:  "380",
			output: false,
		},

		{
			input:  "amd64",
			output: true,
		},

		{
			input:  "arm64",
			output: true,
		},

		{
			input:  "fail",
			output: false,
		},

		{
			input:  "fail-123",
			output: false,
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput := checkValidArch(
			test.input,
		)
		assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d: %s", i, test.input))
	}
}
