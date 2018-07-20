package cloudformation

import (
	"github.com/KablamoOSS/kombustion/internal/plugins"
	"github.com/KablamoOSS/kombustion/types"
)

type (
	// YamlConfig -
	YamlConfig struct {
		AWSTemplateFormatVersion string                      `yaml:"AWSTemplateFormatVersion,omitempty"`
		Description              string                      `yaml:"Description,omitempty"`
		Metadata                 types.TemplateObject        `yaml:"Metadata,omitempty"`
		Parameters               types.TemplateObject        `yaml:"Parameters,omitempty"`
		Mappings                 types.TemplateObject        `yaml:"Mappings,omitempty"`
		Conditions               types.TemplateObject        `yaml:"Conditions,omitempty"`
		Transform                types.TemplateObject        `yaml:"Transform,omitempty"`
		Resources                map[string]types.CfResource `yaml:"Resources"`
		Outputs                  types.TemplateObject        `yaml:"Outputs,omitempty"`
	}

	// YamlCloudformation -
	YamlCloudformation struct {
		AWSTemplateFormatVersion string               `yaml:"AWSTemplateFormatVersion,omitempty"`
		Description              string               `yaml:"Description,omitempty"`
		Metadata                 types.TemplateObject `yaml:"Metadata,omitempty"`
		Parameters               types.TemplateObject `yaml:"Parameters,omitempty"`
		Mappings                 types.TemplateObject `yaml:"Mappings,omitempty"`
		Conditions               types.TemplateObject `yaml:"Conditions,omitempty"`
		Transform                types.TemplateObject `yaml:"Transform,omitempty"`
		Resources                types.TemplateObject `yaml:"Resources"`
		Outputs                  types.TemplateObject `yaml:"Outputs,omitempty"`
	}

	// GenerateParams are required to generate a cloudformation yaml template
	GenerateParams struct {
		Filename               string
		EnvFile                string
		Env                    string
		GenerateDefaultOutputs bool
		ParamMap               map[string]string
		Plugins                []*plugins.PluginLoaded
	}
)
