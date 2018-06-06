package tasks

import (
	"github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/urfave/cli"
)

// InitManifestFlags - Flags that will prevent prompts
var InitManifestFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "name, n",
		Usage: "Set the name of the project",
	},
	cli.StringFlag{
		Name:  "environments",
		Usage: "comma seperated environments eg: production,development",
	},
}

// InitaliseNewManifestTask - Create a new manifest file, and prompt to fill out
// the default required fields
func InitaliseNewManifestTask(c *cli.Context) error {

	// This funciton is a thin layer between the task, and the cli wrapper
	err := manifest.InitaliseNewManifest()

	return err
}
