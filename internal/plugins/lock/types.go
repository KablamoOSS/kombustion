package lock

// Lock - The Lock file
type Lock struct {
	// An array of the plugins we resolved and installed
	Plugins map[string]Plugin `yaml:"plugins"`
}

// Plugin - Store the information of how the plugin was resolved and saved
type Plugin struct {

	// The name of the plugin
	Name string `yaml:"name"`

	// The version downloaded
	Version string `yaml:"version"`

	// A map key by architecture of the resolved plugin
	Resolved map[string]PluginResolution
}

// PluginResolution -
type PluginResolution struct {

	// Exactly where it was downloaded from
	// A list of direct download urls for each architecture
	URL string `yaml:"urls"`

	OperatingSystem string `yaml:"operatingSystem"`

	Architecture string `yaml:"architecture"`

	// Exactly where it is store on disk, relative to the manifest file
	PathOnDisk string `yaml:"pathOnDisk"`

	// A sha256 hash of the plugin
	Hash string `yaml:"hash"`

	// A sha256 hash of the plugin archive (used to compare to the cache)
	ArchiveHash string `yaml:"archiveHash"`

	// The filename of the release archive
	ArchiveName string `yaml:"archiveName"`
}
