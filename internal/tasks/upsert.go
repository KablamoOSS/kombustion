package tasks

import (
	"fmt"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/cloudformation"
	"github.com/KablamoOSS/kombustion/internal/cloudformation/tasks"
	"github.com/KablamoOSS/kombustion/internal/core"
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
	cli.StringSliceFlag{
		Name:  "tag, t",
		Usage: "tags to add to cloudformation stack. eg `CostCenter=example`",
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
		Name:  "iam, i",
		Usage: "gives the capability to perform upserts of IAM resources",
	},
	cli.StringSliceFlag{
		Name:  "capability",
		Usage: "set capabilities for the upsert eg `CAPABILITY_IAM`",
	},
	cli.BoolFlag{
		Name:  "confirm",
		Usage: "Manually confirm required changes before applying",
	},
}

func init() {
	UpsertFlags = append(CloudFormationStackFlags, UpsertFlags...)
}

// Upsert a stack
func Upsert(c *cli.Context) {
	objectStore := core.NewFilesystemStore(".")

	fileName := c.Args().Get(0)
	if fileName == "" {
		printer.Fatal(
			fmt.Errorf("Can't upsert file, no source template provided"),
			"Add the path to the source template file you want to generate like: `kombustion upsert template.yaml`.",
			"https://www.kombustion.io/api/cli/#upsert",
		)
	}

	paramsSlice := c.StringSlice("param")
	paramsMap := cliSliceMap(paramsSlice)

	tagsSlice := c.StringSlice("tag")
	tagsMap := cliSliceMap(tagsSlice)

	stackName := c.String("stack-name")
	profile := c.GlobalString("profile")
	region := c.String("region")
	devPluginPath := c.GlobalString("load-plugin")
	inputParameters := c.String("read-parameters")
	env := c.String("environment")
	generateDefaultOutputs := c.Bool("generate-default-outputs")
	capabilities := getCapabilities(c)
	confirm := c.Bool("confirm")
	manifestFile := c.GlobalString("manifest-file")

	client := &cloudformation.Wrapper{}

	upsert(
		client,
		objectStore,
		fileName,
		stackName,
		profile,
		region,
		paramsMap,
		inputParameters,
		tagsMap,
		devPluginPath,
		env,
		generateDefaultOutputs,
		capabilities,
		confirm,
		manifestFile,
	)
}
func upsert(
	client cloudformation.StackUpserter,
	objectStore core.ObjectStore,
	templatePath string,
	stackName string,
	profile string,
	region string,
	cliParams map[string]string,
	paramsPath string,
	cliTags map[string]string,
	devPluginPath string,
	envName string,
	generateDefaultOutputs bool,
	capabilities []*string,
	confirm bool,
	manifestLocation string,
) {
	printer.Progress("Kombusting")

	lockFile, err := lock.GetLockObject(objectStore, "kombustion.lock")
	if err != nil {
		printer.Fatal(
			fmt.Errorf("Couldn't load lock file: %v", err),
			"Check the file exists, and your user has read permissions",
			"",
		)
	}

	// manifestFile := manifest.FindAndLoadManifest()
	manifestFile, err := manifest.GetManifestObject(objectStore, manifestLocation)
	if err != nil {
		printer.Fatal(
			fmt.Errorf("Couldn't load manifest file: %v", err),
			"Check the file exists, and your user has read permissions",
			"",
		)
	}

	// load all plugins
	loadedPlugins := plugins.LoadPlugins(manifestFile, lockFile)

	// if in devMode optionally load a devMode plugin
	if devPluginPath != "" {
		devPluginLoaded := plugins.LoadDevPlugin(devPluginPath)
		loadedPlugins = append(loadedPlugins, devPluginLoaded)
	}

	if region == "" {
		// If no region was provided by the cli flag, check for the default in the manifest
		if manifestFile.Region != "" {
			region = manifestFile.Region
		}
	}

	acctID := client.Open(profile, region)

	paramMap := make(map[string]string)
	if paramsPath != "" {
		paramMap = readParamsObject(objectStore, paramsPath)
	}

	for key, value := range cliParams {
		paramMap[key] = value
	}

	if env, ok := manifestFile.Environments[envName]; ok {
		if !env.IsAllowlistedAccount(acctID) {
			printer.Fatal(
				fmt.Errorf("Account %s is not allowed for environment %s", acctID, envName),
				fmt.Sprintf("Use allowlisted account, or add account %s to environment accounts in kombustion.yaml", acctID),
				"",
			)
		}
	}

	tags := manifestFile.Tags
	if tags == nil {
		tags = make(map[string]string)
	}
	if env, ok := manifestFile.Environments[envName]; ok {
		for key, value := range env.Tags {
			tags[key] = value
		}
	}
	for key, value := range cliTags {
		tags[key] = value
	}

	printer.Progress("Generating template")
	// Template generation parameters
	generateParams := cloudformation.GenerateParams{
		ObjectStore: objectStore,
		Filename:    templatePath,
		Env:         envName,
		GenerateDefaultOutputs: generateDefaultOutputs || manifestFile.GenerateDefaultOutputs,
		ParamMap:               paramMap,
		Plugins:                loadedPlugins,
	}

	// CloudFormation Stack parameters
	var parameters []*awsCF.Parameter

	fullStackName := cloudformation.GetStackName(manifestFile, templatePath, envName, stackName)

	printer.Progress("Upserting template")

	// FIXME - this previously looked for a --url param to use a template in
	// S3. This should probably be reimplemented around an S3ObjectStore
	templateBody, cfYaml := tasks.GenerateYamlTemplate(generateParams)
	parameters = cloudformation.ResolveParameters(envName, paramMap, cfYaml, manifestFile)

	tasks.UpsertStackBody(
		templateBody,
		parameters,
		capabilities,
		fullStackName,
		client,
		tags,
		confirm,
	)
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
