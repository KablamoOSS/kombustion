package tasks

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/KablamoOSS/kombustion/cloudformation"
	"github.com/KablamoOSS/kombustion/types"
	"github.com/aws/aws-sdk-go/aws"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/urfave/cli"
)

var Upsert_Flags = []cli.Flag{
	cli.StringFlag{
		Name:  "region, r",
		Usage: "region to deploy to",
		Value: "ap-southeast-2",
	},
	cli.StringFlag{
		Name:  "stackName",
		Usage: "stack name to deploy (defaults to filename)",
	},
	cli.StringFlag{
		Name:  "url",
		Usage: "Url to use for the stack definition (s3)",
	},

	// cf generate flags
	cli.StringFlag{
		Name:  "format, f",
		Usage: "cf output format (only yaml is supported)",
		Value: "yaml",
	},
	cli.StringFlag{
		Name:  "env",
		Usage: "environment config to use from ./config/environment.yaml",
	},
	cli.StringFlag{
		Name:  "envFile",
		Usage: "path to the environment.yaml file",
	},
	cli.StringSliceFlag{
		Name:  "param, p",
		Usage: "cloudformation parameters. eg. ( --param Env=dev --param BucketName=test )",
	},
	cli.BoolFlag{
		Name:  "noBaseOutputs, b",
		Usage: "disable generation of outputs for Base AWS types",
	},
	cli.BoolFlag{
		Name:  "allowIAMUpsert, i",
		Usage: "gives the capability to perform upserts of IAM resources",
	},
}

func Upsert(c *cli.Context) {
	upsertStack(c, getCF(c.GlobalString("profile"), c.String("region")))
}

func upsertStack(c *cli.Context, cf *awsCF.CloudFormation) {
	var err error
	var status *awsCF.DescribeStacksOutput

	stackName := c.Args().Get(0)
	if len(c.String("stackName")) > 0 {
		stackName = c.String("stackName")
	}

	capabilities := aws.StringSlice([]string{})
	if c.Bool("allowIAMUpsert") {
		capabilities = aws.StringSlice([]string{"CAPABILITY_NAMED_IAM"})
	}

	if len(c.String("url")) > 0 {
		// use cf template url
		_, err = cf.DescribeStacks(&awsCF.DescribeStacksInput{StackName: aws.String(stackName)})
		if err == nil { //update
			_, err = cf.UpdateStack(&awsCF.UpdateStackInput{
				StackName:    aws.String(stackName),
				TemplateURL:  aws.String(c.String("url")),
				Parameters:   resolveParametersS3(c),
				Capabilities: capabilities,
			})
		} else { //create
			_, err = cf.CreateStack(&awsCF.CreateStackInput{
				StackName:    aws.String(stackName),
				TemplateURL:  aws.String(c.String("url")),
				Parameters:   resolveParametersS3(c),
				Capabilities: capabilities,
			})
		}
		checkError(err)
	} else {
		// use template from file
		data, cfYaml := generateTemplate(c)
		_, err = cf.DescribeStacks(&awsCF.DescribeStacksInput{StackName: aws.String(stackName)})
		if err == nil { //update
			_, err = cf.UpdateStack(&awsCF.UpdateStackInput{
				StackName:    aws.String(stackName),
				TemplateBody: aws.String(string(data)),
				Parameters:   resolveParameters(c, cfYaml),
				Capabilities: capabilities,
			})
		} else {
			_, err = cf.CreateStack(&awsCF.CreateStackInput{
				StackName:    aws.String(stackName),
				TemplateBody: aws.String(string(data)),
				Parameters:   resolveParameters(c, cfYaml),
				Capabilities: capabilities,
			})
		}
		checkError(err)
	}

	// Make sure upsert works
	for {
		time.Sleep(2 * time.Second)
		status, err = cf.DescribeStacks(&awsCF.DescribeStacksInput{StackName: aws.String(stackName)})
		checkError(err)
		if len(status.Stacks) > 0 {
			stack := status.Stacks[0]
			stackStatus := *stack.StackStatus
			fmt.Println(stackStatus)
			if stackStatus != awsCF.StackStatusCreateInProgress &&
				stackStatus != awsCF.StackStatusUpdateInProgress &&
				stackStatus != awsCF.StackStatusUpdateCompleteCleanupInProgress {
				if stackStatus == awsCF.StackStatusCreateComplete ||
					stackStatus == awsCF.StackStatusUpdateComplete {
					os.Exit(0)
				} else {
					log.Error("Upsert Failed: ")
					time.Sleep(2 * time.Second)
					printStackEvents(cf, stackName)
					os.Exit(1)
				}
			}
		}
	}
}

func resolveParameters(c *cli.Context, cfYaml cloudformation.YamlCloudformation) []*awsCF.Parameter {
	results := []*awsCF.Parameter{}

	// Get params from the envFile
	env := cloudformation.ResolveEnvironment(c.String("envFile"), c.String("env"))

	// override envFile values with optional --param values
	params := getParamMap(c)
	for k, v := range params {
		env[k] = v
	}

	// convert to aws Parameter list
	for paramK := range cfYaml.Parameters {
		for k, v := range env {
			if paramK == k {
				if s, ok := v.(string); ok {
					// Filter to params in the stack
					results = append(results, &awsCF.Parameter{
						ParameterKey:   aws.String(k),
						ParameterValue: aws.String(s),
					})
				}
			}
		}
	}

	return results
}

func resolveParametersS3(c *cli.Context) []*awsCF.Parameter {
	results := []*awsCF.Parameter{}

	var params types.ValueMap

	// override envFile values with optional --param values
	paramMap := getParamMap(c)
	for k, v := range paramMap {
		params[k] = v
	}

	// convert to aws Parameter list
	for k, v := range params {
		if s, ok := v.(string); ok {
			// Filter to params in the stack
			results = append(results, &awsCF.Parameter{
				ParameterKey:   aws.String(k),
				ParameterValue: aws.String(s),
			})
		}
	}

	return results
}
