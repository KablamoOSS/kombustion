package cloudformation

import (
	"strings"

	"github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/aws/aws-sdk-go/aws"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/urfave/cli"
)

func GetParamMap(c *cli.Context) map[string]string {
	paramMap := make(map[string]string)
	params := c.StringSlice("param")
	for _, param := range params {
		parts := strings.Split(param, "=")
		if len(parts) > 1 {
			paramMap[parts[0]] = strings.Join(parts[1:], "=")
		}
	}
	return paramMap
}

func ResolveParameters(
	c *cli.Context,
	cfYaml YamlCloudformation,
	cfClient *awsCF.CloudFormation,
) []*awsCF.Parameter {
	results := []*awsCF.Parameter{}

	env := resolveEnvironmentParameters(c.String("environment"))

	// override envFile values with optional --param values
	params := GetParamMap(c)
	for key, value := range params {
		env[key] = value
	}

	// convert to aws Parameter list
	for paramKey := range cfYaml.Parameters {
		for key, value := range env {
			if paramKey == key {
				// Filter to params in the stack
				results = append(results, &awsCF.Parameter{
					ParameterKey:   aws.String(key),
					ParameterValue: aws.String(value),
				})
			}
		}
	}

	return results
}

func ResolveParametersS3(
	c *cli.Context,
	cfClient *awsCF.CloudFormation,
) []*awsCF.Parameter {

	results := []*awsCF.Parameter{}

	params := make(map[string]string)

	env := resolveEnvironmentParameters(c.String("environment"))
	for key, value := range params {
		env[key] = value
	}

	// convert to aws Parameter list
	for key, value := range params {
		// Filter to params in the stack
		results = append(results, &awsCF.Parameter{
			ParameterKey:   aws.String(key),
			ParameterValue: aws.String(value),
		})
	}

	return results
}

func resolveEnvironmentParameters(environment string) (parameters map[string]string) {
	manifestFile := manifest.FindAndLoadManifest()
	if manifestFile != nil {
		envParams := manifestFile.Environments[environment].Parameters
		if envParams != nil {
			parameters = envParams
		}
	}
	return
}
