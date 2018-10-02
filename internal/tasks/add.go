package tasks

import (
	"github.com/KablamoOSS/go-cli-printer"

	"github.com/KablamoOSS/kombustion/config"
	"github.com/KablamoOSS/kombustion/internal/core"
	"github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/KablamoOSS/kombustion/internal/plugins"
	"github.com/KablamoOSS/kombustion/internal/plugins/lock"
	"github.com/urfave/cli"
)

// AddPluginToManifest file and update it
func AddPluginToManifest(c *cli.Context) error {
	// Get the plugin to add
	pluginNames := c.Args()
	manifestLocation := c.GlobalString("manifest-file")

	objectStore := core.NewFilesystemStore(".")
	addPluginToManifest(objectStore, pluginNames, manifestLocation)
	return nil
}

func addPluginToManifest(objectStore core.ObjectStore, pluginNames []string, manifestLocation string) {
	printer.Step("Add plugins")
	// Try and load the manifest
	manifestFile, err := manifest.GetManifestObject(objectStore, manifestLocation)
	if err != nil {
		printer.Fatal(err, config.ErrorHelpInfo, "")
	}

	// Add them
	_, err = plugins.AddPluginsToManifest(objectStore, manifestFile, pluginNames, manifestLocation)
	if err != nil {
		printer.Fatal(err, config.ErrorHelpInfo, "")
	}

	lockFile, err := lock.GetLockObject(objectStore, "kombustion.lock")
	if err != nil {
		printer.Fatal(err, config.ErrorHelpInfo, "")
	}

	// Now install them
	lockFile = plugins.InstallPlugins(lockFile)

	lockFile.Save(objectStore)
	if err != nil {
		printer.Fatal(err, config.ErrorHelpInfo, "")
	}
}
