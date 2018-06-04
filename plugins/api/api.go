package api

import (
	"github.com/KablamoOSS/kombustion/plugins"
	apiTypes "github.com/KablamoOSS/kombustion/plugins/api/types"
	"github.com/KablamoOSS/kombustion/types"
)

// This file contains the public api for plugins to interface with Kombustion
// Note that all the types in out and out are only from teh stdlib.
// To cross the plugin boundary we marshall to binary and back.

// RegisterPlugin to provide the name, prefix and version, and requiresAWSSession
func RegisterPlugin(config apiTypes.Config) []byte {
	return plugins.MarshallConfig(config)
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
		return plugins.MarshallResource(resource(ctx, name, data))
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
		return plugins.MarshallMapping(mapping(ctx, name, data))
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
		return plugins.MarshallOutput(output(ctx, name, data))
	}
}
