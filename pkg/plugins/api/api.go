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
		ctx map[string]interface{},
		name string,
		data string,
	) (cf types.TemplateObject),
) func(
	ctx map[string]interface{},
	name string,
	data string,
) []byte {
	return func(
		ctx map[string]interface{},
		name string,
		data string,
	) []byte {
		return marshallResource(resource(ctx, name, data))
	}
}

// [ Mapping ]--------------------------------------------------------------------------------------

// RegisterMapping for your plugin
func RegisterMapping(
	mapping func(
		ctx map[string]interface{},
		name string,
		data string,
	) (cf types.TemplateObject),
) func(
	ctx map[string]interface{},
	name string,
	data string,
) []byte {
	return func(
		ctx map[string]interface{},
		name string,
		data string,
	) []byte {
		return marshallMapping(mapping(ctx, name, data))
	}
}

// [ Outputs ]--------------------------------------------------------------------------------------

// RegisterOutput for your plugin
func RegisterOutput(
	output func(
		ctx map[string]interface{},
		name string,
		data string,
	) (cf types.TemplateObject),
) func(
	ctx map[string]interface{},
	name string,
	data string,
) []byte {
	return func(
		ctx map[string]interface{},
		name string,
		data string,
	) []byte {
		return marshallOutput(output(ctx, name, data))
	}
}
