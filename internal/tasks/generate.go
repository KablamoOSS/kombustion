package tasks

import (
	"fmt"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/cloudformation"
	"github.com/KablamoOSS/kombustion/internal/cloudformation/tasks"
	"github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/KablamoOSS/kombustion/internal/plugins"
	"github.com/KablamoOSS/kombustion/internal/plugins/lock"
	"github.com/urfave/cli"
)

// GenerateFlags for kombustion upsert
var GenerateFlags = []cli.Flag{
	cli.StringSliceFlag{
		Name:  "param, p",
		Usage: "cloudformation parameters. eg. `--param Env=dev --param BucketName=test`",
	},
	cli.BoolFlag{
		Name:  "generate-default-outputs, b",
		Usage: "disable generation of outputs for Base AWS types",
	},
	cli.BoolFlag{
		Name:  "write-parameters, w",
		Usage: "Write parameters to a file",
	},
	cli.StringFlag{
		Name:  "output-directory, d",
		Usage: "Directory to write generated yaml to",
		Value: "compiled",
	},
}

func init() {
	GenerateFlags = append(GenerateFlags, CloudFormationStackFlags...)
}

// Generate a template and save it to disk, without upserting it
func Generate(c *cli.Context) {
	printer.Step("Generate template")
	printer.Progress("Kombusting")

	fileName := c.Args().Get(0)
	if fileName == "" {
		printer.Fatal(
			fmt.Errorf("Can't generate file, no source template provided"),
			fmt.Sprintf(
				"Add the path to the source template file you want to generate like: `kombustion generate template.yaml`.",
			),
			"https://www.kombustion.io/api/cli/#generate",
		)
	}

	paramMap := cloudformation.GetParamMap(c)

	lockFile := lock.FindAndLoadLock()

	manifestFile := manifest.FindAndLoadManifest()

	// load all plugins
	loadedPlugins := plugins.LoadPlugins(manifestFile, lockFile)

	// if in devMode optionally load a devMode plugin
	devPluginPath := c.GlobalString("load-plugin")

	if devPluginPath != "" {
		devPluginLoaded := plugins.LoadDevPlugin(devPluginPath)
		loadedPlugins = append(loadedPlugins, devPluginLoaded)
	}

	printer.Progress("Generating template")
	tasks.GenerateTemplate(cloudformation.GenerateParams{
		Filename:    fileName,
		Directory:   c.String("output-directory"),
		WriteParams: c.Bool("write-parameters"),
		Env:         c.String("env"),
		GenerateDefaultOutputs: c.Bool("generate-default-outputs") || manifestFile.GenerateDefaultOutputs,
		ParamMap:               paramMap,
		Plugins:                loadedPlugins,
	})
}
