package tasks

import (
	"log"

	"github.com/KablamoOSS/go-cli-printer"

	"github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/KablamoOSS/kombustion/internal/plugins"
	"github.com/urfave/cli"
)

// AddPluginToManifest file and update it
func AddPluginToManifest(c *cli.Context) error {
	printer.Step("Adding plugins")
	// Try and load the manifest
	manifestFile := manifest.FindAndLoadManifest()

	// Get the plugin to add
	pluginNames := c.Args()

	// Add them
	_, err := plugins.AddPluginsToManifest(manifestFile, pluginNames)
	if err != nil {
		log.Fatal(err)
	}

	// Now install them
	err = plugins.InstallPlugins()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
