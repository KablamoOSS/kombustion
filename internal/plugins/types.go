package plugins

import (
	apiTypes "github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
)

// PluginLoaded is a fully loaded plugin
type PluginLoaded struct {
	Resources *map[string]func(
		name string,
		data string,
	) []byte

	Outputs *map[string]func(
		name string,
		data string,
	) []byte

	Mappings *map[string]func(
		name string,
		data string,
	) []byte

	Config         apiTypes.Config
	InternalConfig struct {
		Prefix     string
		PathOnDisk string
	}
}
