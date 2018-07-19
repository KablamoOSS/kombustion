package types

import (
	kombustionTypes "github.com/KablamoOSS/kombustion/types"
)

// PluginResult result of a parserFunc
type PluginParserResult struct {
	Conditions kombustionTypes.TemplateObject
	Metadata   kombustionTypes.TemplateObject
	Mappings   kombustionTypes.TemplateObject
	Outputs    kombustionTypes.TemplateObject
	Parameters kombustionTypes.TemplateObject
	Resources  kombustionTypes.TemplateObject

	Errors []error
}

// Config provides Kombustion with information about your plugin
type Config struct {
	Name    string
	Version string
	Prefix  string
	Types   map[string]TypeMapping
	Help    Help
}

// Help - a set of available documentation fields
type Help struct {
	// A short description of what the plugin does
	Description string

	// Examples/Snippets of how this plugin can be used
	Snippets []string
}

// TypeMapping - recursive list of types with its associated config object
type TypeMapping struct {
	Name        string
	Description string
	Config      interface{}
	Parser      func(
		name string,
		data string,
	) (
		conditions kombustionTypes.TemplateObject,
		metadata kombustionTypes.TemplateObject,
		mappings kombustionTypes.TemplateObject,
		outputs kombustionTypes.TemplateObject,
		parameters kombustionTypes.TemplateObject,
		resources kombustionTypes.TemplateObject,
		errors []error,
	)
}
