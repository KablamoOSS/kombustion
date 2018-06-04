package manifest

// Manifest - kombustion.yaml manifest file
type Manifest struct {
	// Name of this Kombustion project, this can be used to prefix stack names
	Name string `yaml:"name"`

	// A list of plugins used in this project
	Plugins map[string]Plugin `yaml:"plugins,omitempty"`

	// An array of architectures this project will be used in. This can be used to limit the number
	// of plugins downloaded.
	// If omitted a plugin will download for all achitectures.
	Architectures []string `yaml:"architectures,omitempty"`

	// A map of enviroment specific configuration
	Environments map[string]Environment `yaml:"environments,omitempty"`

	// A flag to indicate if default exports should be added to stacks in this project
	// this defaults to false
	HideDefaultExports bool `yaml:"hideDefaultExports"`
}

// Plugin specification inside the manifest
type Plugin struct {
	// The name of the plugin to include
	// In the initial version of plugin management, this must be a github url
	Name string `yaml:"name"`

	// The version to download of the plugin
	// In the initial version this must be a tag on the github url, or
	// `latest` which will use the latest tag.
	Version string `yaml:"version"`

	// TODO: implement Plugin Aliases
	// This will be useful if two plugins want to use the same namespace
	// Ideally, that wouldn't happen. Ideally.
	Alias string `yaml:"alias"`
}

// Environment specific parameters
type Environment struct {
	// A whitelist of allowed accountID's for this enviroment.
	// This restricts where stacks in this project can be deployed to.
	// This is a harm minimisation feature, intended to limit damage from human error.
	AccountIDs []string `yaml:"accountIDs,omitempty"`

	// A list of Parameters to substiture into the Parameters field of all stacks in this
	// project.
	Parameters map[string]string `yaml:"parameters,omitempty"`
}
