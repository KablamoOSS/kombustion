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

// [ Resources ]------------------------------------------------------------------------------------

// RegisterResource for your plugin
func RegisterResource(
	resource func(
		name string,
		data string,
	) (cf types.TemplateObject, err []error),
) func(
	name string,
	data string,
) []byte {
	return func(
		name string,
		data string,
	) []byte {
		result, resultErr := resource(name, data)
		return marshallResource(result, resultErr)
	}
}

// [ Mapping ]--------------------------------------------------------------------------------------

// RegisterMapping for your plugin
func RegisterMapping(
	mapping func(
		name string,
		data string,
	) (cf types.TemplateObject, err []error),
) func(
	name string,
	data string,
) []byte {
	return func(
		name string,
		data string,
	) []byte {
		result, resultErr := mapping(name, data)
		return marshallMapping(result, resultErr)
	}
}

// [ Outputs ]--------------------------------------------------------------------------------------

// RegisterOutput for your plugin
func RegisterOutput(
	output func(
		name string,
		data string,
	) (cf types.TemplateObject, err []error),
) func(
	name string,
	data string,
) []byte {
	return func(
		name string,
		data string,
	) []byte {
		result, resultErr := output(name, data)
		return marshallOutput(result, resultErr)
	}
}
