package tasks

import (
	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/cloudformation"
	"github.com/KablamoOSS/kombustion/internal/cloudformation/tasks"
	"github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/aws/aws-sdk-go/aws"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/urfave/cli"
)

// UpsertFlags for kombustion upsert
var UpsertFlags = []cli.Flag{
	cli.StringSliceFlag{
		Name:  "param, p",
		Usage: "cloudformation parameters. eg. [ --param Env=dev --param BucketName=test ]",
	},
	cli.BoolFlag{
		Name:  "no-base-outputs, b",
		Usage: "disable generation of outputs for Base AWS types",
	},
	cli.BoolFlag{
		Name:  "iam, i",
		Usage: "gives the capability to perform upserts of IAM resources",
	},
	cli.StringSliceFlag{
		Name:  "capability",
		Usage: "set capabilities for the upsert, eg [ --capability CAPABILITY_IAM ]",
	},
}

func init() {
	UpsertFlags = append(CloudFormationStackFlags)
}

// Upsert a stack
func Upsert(c *cli.Context) {
	printer.Step("Upserting stack")
	manifest := manifest.FindAndLoadManifest()

	cfClient := tasks.GetCloudformationClient(
		c.GlobalString("profile"),
		c.String("region"),
	)

	paramMap := cloudformation.GetParamMap(c)

	// Template generation parameters
	generateParams := cloudformation.GenerateParams{
		Filename:           c.Args().Get(0),
		EnvFile:            c.String("env-file"),
		Env:                c.String("env"),
		DisableBaseOutputs: c.Bool("no-base-outputs"),
		ParamMap:           paramMap,
	}

	capabilities := getCapabilities(c)

	// Cloudformation Stack parameters
	var parameters []*awsCF.Parameter

	stackName := c.Args().Get(0)
	if len(c.String("stack-name")) > 0 {
		stackName = c.String("stack-name")
	}
	if len(c.String("url")) > 0 {
		// TODO: We probably need to download the template to determine what params
		// it needs, and filter the available params only to those
		parameters = cloudformation.ResolveParametersS3(c, manifest)

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
		parameters = cloudformation.ResolveParameters(c, cfYaml, manifest)

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
