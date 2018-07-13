package tasks

import (
	"fmt"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/cloudformation"
	"github.com/KablamoOSS/kombustion/internal/cloudformation/tasks"
	"github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/KablamoOSS/kombustion/internal/plugins"
	"github.com/KablamoOSS/kombustion/internal/plugins/lock"
	"github.com/aws/aws-sdk-go/aws"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/urfave/cli"
)

// UpsertFlags for kombustion upsert
var UpsertFlags = []cli.Flag{
	cli.StringSliceFlag{
		Name:  "param, p",
		Usage: "cloudformation parameters. eg `BucketName=test`",
	},
	cli.BoolFlag{
		Name:  "generate-default-outputs, b",
		Usage: "disable generation of outputs for Base AWS types",
	},
	cli.BoolFlag{
		Name:  "iam, i",
		Usage: "gives the capability to perform upserts of IAM resources",
	},
	cli.StringSliceFlag{
		Name:  "capability",
		Usage: "set capabilities for the upsert eg `CAPABILITY_IAM`",
	},
}

func init() {
	UpsertFlags = append(CloudFormationStackFlags, UpsertFlags...)
}

// Upsert a stack
func Upsert(c *cli.Context) {
	printer.Progress("Kombusting")

	fileName := c.Args().Get(0)
	if fileName == "" {
		printer.Fatal(
			fmt.Errorf("Can't upsert file, no source template provided"),
			fmt.Sprintf(
				"Add the path to the source template file you want to generate like: `kombustion upsert template.yaml`.",
			),
			"https://www.kombustion.io/api/cli/#upsert",
		)
	}

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

	region := c.String("region")
	if region == "" {
		// If no region was provided by the cli flag, check for the default in the manifest
		if manifestFile.Region != "" {
			region = manifestFile.Region
		}
	}

	cfClient := tasks.GetCloudformationClient(
		c.GlobalString("profile"),
		region,
	)

	paramMap := cloudformation.GetParamMap(c)

	environment := c.String("environment")

	printer.Progress("Generating template")
	// Template generation parameters
	generateParams := cloudformation.GenerateParams{
		Filename: fileName,
		Env:      environment,
		GenerateDefaultOutputs: c.Bool("generate-default-outputs"),
		ParamMap:               paramMap,
		Plugins:                loadedPlugins,
	}

	capabilities := getCapabilities(c)

	// CloudFormation Stack parameters
	var parameters []*awsCF.Parameter

	stackName := cloudformation.GetStackName(manifestFile, fileName, environment, c.String("stack-name"))

	printer.Progress("Upserting template")
	if len(c.String("url")) > 0 {
		// TODO: We probably need to download the template to determine what params
		// it needs, and filter the available params only to those
		parameters = cloudformation.ResolveParametersS3(c, manifestFile)

		templateURL := c.String("url")

		tasks.UpsertStackViaS3(
			templateURL,
			parameters,
			capabilities,
			stackName,
			cfClient,
		)
	} else {

		templateBody, cfYaml := tasks.GenerateYamlTemplate(generateParams)
		parameters = cloudformation.ResolveParameters(c, cfYaml, manifestFile)

		tasks.UpsertStack(
			templateBody,
			parameters,
			capabilities,
			stackName,
			cfClient,
		)
	}
}

// Extract capabilities from the cli call
func getCapabilities(c *cli.Context) []*string {
	capabilities := aws.StringSlice([]string{})

	capabilitySlices := c.StringSlice("capability")

	iamCapability := "CAPABILITY_NAMED_IAM"
	haveAddedIam := false

	for _, capability := range capabilitySlices {
		if capability == iamCapability {
			haveAddedIam = true
		}
		capabilities = append(capabilities, &capability)
	}

	if c.Bool("iam") {
		if haveAddedIam == false {
			capabilities = append(capabilities, &iamCapability)
		}
	}

	return capabilities
}
