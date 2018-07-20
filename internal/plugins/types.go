package plugins

import (
	apiTypes "github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
)

// PluginLoaded is a fully loaded plugin
type PluginLoaded struct {
	Config         apiTypes.Config
	InternalConfig struct {
		Prefix     string
		PathOnDisk string
	}

	// The Parser functions from the plugin
	Parsers *map[string]func(
		name string,
		data string,
	) []byte
}
