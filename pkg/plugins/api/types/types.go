package types

import (
	kombustionTypes "github.com/KablamoOSS/kombustion/types"
)

// PluginParserResult result of a parserFunc
type PluginParserResult struct {
	Conditions kombustionTypes.TemplateObject
	Metadata   kombustionTypes.TemplateObject
	Mappings   kombustionTypes.TemplateObject
	Outputs    kombustionTypes.TemplateObject
	Parameters kombustionTypes.TemplateObject
	Resources  kombustionTypes.TemplateObject
	Transform  kombustionTypes.TemplateObject

	Errors []error
}

// Config provides Kombustion with information about your plugin
type Config struct {
	Name    string
	Version string
	Prefix  string
	Help    Help
}

// Help - a set of available documentation fields
type Help struct {
	// A short description of what the plugin does
	Description string

	// Examples/Snippets of how this plugin can be used
	Snippets []string
	Types    []TypeMapping
}

// TypeMapping - recursive list of types with its associated config object
type TypeMapping struct {
	Name        string
	Description string
	Config      interface{}
}
