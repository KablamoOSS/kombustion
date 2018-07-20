package api

import (
	apiTypes "github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
	"github.com/KablamoOSS/kombustion/types"
)

// This file contains the public api for plugins to interface with Kombustion
// Note that all the types in out and out are only from teh stdlib.
// To cross the plugin boundary we marshall to binary and back.

// RegisterPlugin to provide the name, prefix and version, and requiresAWSSession
func RegisterPlugin(config apiTypes.Config) []byte {
	return marshallConfig(config)
}

// RegisterParser for your plugin
func RegisterParser(
	// parser has the same type signature as types.ParserFunc
	parser func(
		name string,
		data string,
	) (
		conditions types.TemplateObject,
		metadata types.TemplateObject,
		mappings types.TemplateObject,
		outputs types.TemplateObject,
		parameters types.TemplateObject,
		resources types.TemplateObject,
		transform types.TemplateObject,
		errors []error,
	),
) func(
	name string,
	data string,
) []byte {
	return func(
		name string,
		data string,
	) []byte {
		conditions,
			metadata,
			mappings,
			outputs,
			parameters,
			resources,
			transform,
			errors := parser(name, data)
		return marshallParserResult(
			conditions,
			metadata,
			mappings,
			outputs,
			parameters,
			resources,
			transform,
			errors,
		)
	}
}
