package tasks

import (
	"github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/config"
	"github.com/KablamoOSS/kombustion/internal/core"
	"github.com/KablamoOSS/kombustion/internal/plugins"
	"github.com/KablamoOSS/kombustion/internal/plugins/lock"
	"github.com/urfave/cli"
)

// InstallPlugins to reconcile the manifest to the lock file, and then the lock file to the disk
// Which is to say, ensure the manifest and lock file agree with the state of the plugins
// and then ensure the lock file agrees with the disk on the state of the plugins
func InstallPlugins(c *cli.Context) {
	objectStore := core.NewFilesystemStore(".")
	installPlugins(objectStore)
}

func installPlugins(objectStore core.ObjectStore) {
	printer.Step("Install plugins")
	printer.Progress("Kombusting")

	lockFile, err := lock.GetLockObject(objectStore, "kombustion.lock")
	if err != nil {
		printer.Fatal(err, config.ErrorHelpInfo, "")
	}

	lockFile = plugins.InstallPlugins(lockFile)

	lockFile.Save(objectStore)
	if err != nil {
		printer.Fatal(err, config.ErrorHelpInfo, "")
	}
}
