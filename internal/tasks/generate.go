package tasks

import (
	"encoding/json"
	"fmt"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/cloudformation"
	"github.com/KablamoOSS/kombustion/internal/cloudformation/tasks"
	"github.com/KablamoOSS/kombustion/internal/core"
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
	cli.StringFlag{
		Name:  "read-parameters",
		Usage: "Read parameters from a file",
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
	objectStore := core.NewFilesystemStore(".")

	fileName := c.Args().Get(0)
	if fileName == "" {
		printer.Fatal(
			fmt.Errorf("Can't generate file, no source template provided"),
			"Add the path to the source template file you want to generate like: `kombustion generate template.yaml`.",
			"https://www.kombustion.io/api/cli/#generate",
		)
	}

	paramsSlice := c.StringSlice("param")
	paramsMap := cliSliceMap(paramsSlice)

	// if in devMode optionally load a devMode plugin
	devPluginPath := c.GlobalString("load-plugin")
	outputDirectory := c.String("output-directory")
	inputParameters := c.String("read-parameters")
	outputParameters := c.Bool("write-parameters")
	env := c.String("env")
	generateDefaultOutputs := c.Bool("generate-default-outputs")

	generate(
		objectStore,
		fileName,
		paramsMap,
		inputParameters,
		devPluginPath,
		outputDirectory,
		outputParameters,
		env,
		generateDefaultOutputs,
	)
}

func generate(
	objectStore core.ObjectStore,
	templatePath string,
	cliParams map[string]string,
	paramsPath string,
	devPluginPath string,
	outputDirectory string,
	outputParameters bool,
	env string,
	generateDefaultOutputs bool,
) {
	printer.Step("Generate template")
	printer.Progress("Kombusting")

	paramMap := make(map[string]string)
	if paramsPath != "" {
		paramMap = readParamsObject(objectStore, paramsPath)
	}

	for key, value := range cliParams {
		paramMap[key] = value
	}

	lockFile, err := lock.GetLockObject(objectStore, "kombustion.lock")
	if err != nil {
		printer.Fatal(
			fmt.Errorf("Couldn't load lock file: %v", err),
			"Check the file exists, and your user has read permissions",
			"",
		)
	}

	// manifestFile := manifest.FindAndLoadManifest()
	manifestFile, err := manifest.GetManifestObject(objectStore)
	if err != nil {
		printer.Fatal(
			fmt.Errorf("Couldn't load manifest file: %v", err),
			"Check the file exists, and your user has read permissions",
			"",
		)
	}

	// load all plugins
	loadedPlugins := plugins.LoadPlugins(manifestFile, lockFile)

	if devPluginPath != "" {
		devPluginLoaded := plugins.LoadDevPlugin(devPluginPath)
		loadedPlugins = append(loadedPlugins, devPluginLoaded)
	}

	printer.Progress("Generating template")
	tasks.GenerateTemplate(cloudformation.GenerateParams{
		ObjectStore:            objectStore,
		Filename:               templatePath,
		Directory:              outputDirectory,
		WriteParams:            outputParameters,
		Env:                    env,
		ParamMap:               paramMap,
		Plugins:                loadedPlugins,
		GenerateDefaultOutputs: generateDefaultOutputs || manifestFile.GenerateDefaultOutputs,
	})
}

func readParamsObject(objStore core.ObjectStore, path string) map[string]string {
	body, err := objStore.Get(path)
	if err != nil {
		printer.Fatal(
			fmt.Errorf("Couldn't read params file: %v", err),
			"Check the file exists, and your user has read permissions",
			"",
		)
	}

	cfParams := []cloudformation.Parameter{}
	if err = json.Unmarshal(body, &cfParams); err != nil {
		printer.Fatal(
			fmt.Errorf("Couldn't unmarshal params file: %v", err),
			"Check the file is valid JSON, in the standard AWS cli format",
			"https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Parameter.html",
		)
	}

	params := make(map[string]string)
	for _, param := range cfParams {
		params[param.ParameterKey] = param.ParameterValue
	}

	return params
}
