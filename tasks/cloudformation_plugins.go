package tasks

import (
	"fmt"

	"github.com/KablamoOSS/kombustion/cloudformation"
	"github.com/urfave/cli"
)

var GetPlugin_Flags = []cli.Flag{}
var PrintPlugins_Flags = []cli.Flag{}
var DeletePlugin_Flags = []cli.Flag{}

func PrintPlugins(c *cli.Context) {
	docs := cloudformation.PluginDocs()
	for name, _ := range docs {
		fmt.Println(name)
	}
}

func GetPlugin(c *cli.Context) {
	cloudformation.DownloadPlugin(c.Args().Get(0))
}

func DeletePlugin(c *cli.Context) {
	cloudformation.DeletePlugin(c.Args().Get(0))
}
