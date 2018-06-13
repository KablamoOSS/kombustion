package tasks

import (
	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/cloudformation/tasks"
	"github.com/urfave/cli"
)

// PrintEvents outputs the events of a stack
func PrintEvents(c *cli.Context) {
	printer.Step("Print events")
	printer.Progress("Kombusting")

	cf := tasks.GetCloudformationClient(c.GlobalString("profile"), c.String("region"))
	stackName := c.Args().Get(0)
	tasks.PrintStackEvents(cf, stackName)
}
