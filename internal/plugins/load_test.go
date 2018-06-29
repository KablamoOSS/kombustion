package plugins

import (
	"fmt"
	"testing"

	printer "github.com/KablamoOSS/go-cli-printer"
	pluginTypes "github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
	"github.com/stretchr/testify/assert"
)

func TestConfigValid(t *testing.T) {
	printer.Test()

	type input struct {
		config        pluginTypes.Config
		pluginName    string
		pluginVersion string
	}

	tests := []struct {
		input  input
		output bool
	}{
		{
			input: input{
				config: pluginTypes.Config{
					Name:   "Example",
					Prefix: "Example",
					Help:   pluginTypes.Help{},
				},
				pluginName:    "Example",
				pluginVersion: "0.1.0",
			},
			output: true,
		},
		{
			input: input{
				config: pluginTypes.Config{
					Name:   "Example",
					Prefix: "AWS",
					Help:   pluginTypes.Help{},
				},
				pluginName:    "Example",
				pluginVersion: "0.1.0",
			},
			output: false,
		},
		{
			input: input{
				config: pluginTypes.Config{
					Name:   "Example",
					Prefix: "Custom",
					Help:   pluginTypes.Help{},
				},
				pluginName:    "Example",
				pluginVersion: "0.1.0",
			},
			output: false,
		},
		{
			input: input{
				config: pluginTypes.Config{
					Name:   "Example",
					Prefix: "Kombustion",
					Help:   pluginTypes.Help{},
				},
				pluginName:    "Example",
				pluginVersion: "0.1.0",
			},
			output: false,
		},
		{
			input: input{
				config: pluginTypes.Config{
					Name:   "",
					Prefix: "Kombustion",
					Help:   pluginTypes.Help{},
				},
				pluginName:    "Example",
				pluginVersion: "0.1.0",
			},
			output: false,
		},
		{
			input: input{
				config: pluginTypes.Config{
					Name:   "Example",
					Prefix: "",
					Help:   pluginTypes.Help{},
				},
				pluginName:    "Example",
				pluginVersion: "0.1.0",
			},
			output: false,
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput := configIsValid(
			test.input.config,
			test.input.pluginName,
			test.input.pluginVersion,
		)
			assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d:", i))
	}
}
