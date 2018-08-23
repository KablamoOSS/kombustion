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
		input input
		valid bool
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
			valid: true,
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
			valid: false,
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
			valid: false,
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
			valid: false,
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
			valid: false,
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
			valid: false,
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		if test.valid {
			assert.NotPanics(
				func() {
					validateConfig(
						test.input.config,
						test.input.pluginName,
						test.input.pluginVersion,
					)
				},
				fmt.Sprintf("Test %d:", i),
			)
		} else {
			assert.Panics(
				func() {
					validateConfig(
						test.input.config,
						test.input.pluginName,
						test.input.pluginVersion,
					)
				},
				fmt.Sprintf("Test %d:", i),
			)
		}
	}
}
