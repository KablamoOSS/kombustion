package tasks

import (
	"fmt"

	"github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/config"
	"github.com/KablamoOSS/kombustion/internal/plugins"
	"github.com/urfave/cli"
)

// InstallPlugins to reconcile the manifest to the lock file, and then the lock file to the disk
// Which is to say, ensure the manifest and lock file agree with the state of the plugins
// and then ensure the lock file agrees with the disk on the state of the plugins
func InstallPlugins(c *cli.Context) {
	printer.Step("Install plugins")
	printer.Progress("Kombusting")

	err := plugins.InstallPlugins()
	if err != nil {
		printer.Fatal(err, fmt.Sprintf("%s\n%s", "Instaling plugins failed. Try again.", config.ErrorHelpInfo), "")
	}
}
