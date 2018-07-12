package tasks

import (
	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/core"
	"github.com/urfave/cli"
)

// UpdateFlags - Flags that will prevent prompts
var UpdateFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "y",
		Usage: "dont prompt if there is an update",
	},
}

// Update kombustion
func Update(c *cli.Context) {
	printer.Progress("Kombusting")

	core.Update(c.App.Version, c.Bool("y"))
}
