package cloudformation

import (
	"fmt"
	"testing"

	"github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/KablamoOSS/kombustion/types"
	"github.com/aws/aws-sdk-go/aws"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

type mockCloudFormationClient struct {
	cloudformationiface.CloudFormationAPI
}

func TestResolveEnvironmentParameters(t *testing.T) {
	tests := []struct {
		name        string
		environment string
		manifest    manifest.Manifest
		output      map[string]string
	}{
		{
			name:        "Returns map of env vars",
			environment: "development",
			manifest: manifest.Manifest{
				Name:               "TestManifestWithEnvironment",
				Plugins:            nil,
				Architectures:      []string(nil),
				GenerateDefaultOutputs: false,
				Environments: map[string]manifest.Environment{
					"development": {
						AccountIDs: nil,
						Parameters: map[string]string{
							"parameterOneName":   "parameterOneValue",
							"parameterTwoName":   "8654238642489624862",
							"parameterThreeName": "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s",
							"parameterFourName":  "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-",
						},
					},
				},
			},
			output: map[string]string{
				"parameterOneName":   "parameterOneValue",
				"parameterTwoName":   "8654238642489624862",
				"parameterThreeName": "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s",
				"parameterFourName":  "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-",
			},
		},
		{
			name:        "Returns emtpy map",
			environment: "production",
			manifest: manifest.Manifest{
				Name:               "TestManifestWithEnvironment",
				Plugins:            nil,
				Architectures:      []string(nil),
				GenerateDefaultOutputs: false,
				Environments: map[string]manifest.Environment{
					"development": {
						AccountIDs: nil,
						Parameters: map[string]string{
							"parameterOneName":   "parameterOneValue",
							"parameterTwoName":   "8654238642489624862",
							"parameterThreeName": "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s",
							"parameterFourName":  "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-",
						},
					},
				},
			},
			output: nil,
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput := resolveEnvironmentParameters(&test.manifest, test.environment)

		matches := 0
		for key, value := range testOutput {
			for paramKey, paramVal := range test.output {
				if key == paramKey && value == paramVal {
					matches = matches + 1
				}
			}
		}
		assert.Equal(len(test.output), matches, fmt.Sprintf("Test %d: %s", i, test.name))
	}
}

func TestResolveParameters(t *testing.T) {
	type input struct {
		ctx          *cli.Context
		cfYaml       YamlCloudformation
		cfClient     *awsCF.CloudFormation
		manifestFile *manifest.Manifest
	}

	tests := []struct {
		name   string
		input  input
		output []*awsCF.Parameter
	}{
		{
			name: "Dev",
			input: input{
				ctx: func() *cli.Context {
					var context *cli.Context

					app := cli.NewApp()

					app.Flags = []cli.Flag{
						cli.StringSliceFlag{
							Name:  "param, p",
							Usage: "cloudformation parameters. eg. ( -p Env=dev -p BucketName=test )",
						},
						cli.StringFlag{
							Name:  "environment, e",
							Usage: "environment config to use from ./kombustion.yaml",
						},
					}

					app.Action = func(c *cli.Context) error {
						context = c
						return nil
					}

					app.Run([]string{
						"",
						"--environment", "development",
						"--param", "parameterOneName=parameterOneValue",
						"--param", "parameterTwoName=8654238642489624862",
					})
					return context
				}(),
				cfYaml: YamlCloudformation{
					AWSTemplateFormatVersion: "version",
					Description:              "Test Template",
					Parameters: types.TemplateObject{
						"parameterOneName":   "",
						"parameterTwoName":   "",
						"parameterThreeName": "",
						"parameterFourName":  "",
					},
					Mappings:   types.TemplateObject{},
					Conditions: types.TemplateObject{},
					Transform:  types.TemplateObject{},
					Resources:  types.TemplateObject{},
					Outputs:    types.TemplateObject{},
				},
				manifestFile: &manifest.Manifest{
					Name:               "TestManifestWithEnvironment",
					Plugins:            nil,
					Architectures:      []string(nil),
					GenerateDefaultOutputs: false,
					Environments: map[string]manifest.Environment{
						"development": {
							AccountIDs: nil,
							Parameters: map[string]string{
								"parameterThreeName": "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s",
								"parameterFourName":  "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-",
							},
						},
					},
				},
			},
			output: []*awsCF.Parameter{
				{
					ParameterKey:   aws.String("parameterOneName"),
					ParameterValue: aws.String("parameterOneValue"),
				},
				{
					ParameterKey:   aws.String("parameterTwoName"),
					ParameterValue: aws.String("8654238642489624862"),
				},
				{
					ParameterKey:   aws.String("parameterThreeName"),
					ParameterValue: aws.String("3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s"),
				},
				{
					ParameterKey:   aws.String("parameterFourName"),
					ParameterValue: aws.String("hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-"),
				},
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput := ResolveParameters(
			test.input.ctx,
			test.input.cfYaml,
			test.input.manifestFile,
		)
		matches := 0
		for _, outputParam := range testOutput {
			for _, param := range test.output {
				if *outputParam.ParameterKey == *param.ParameterKey &&
					*outputParam.ParameterValue == *param.ParameterValue {
					matches = matches + 1
				}
			}
		}

		assert.Equal(len(test.output), matches, fmt.Sprintf("Test %d: %s", i, test.name))
	}
}

func TestResolveParametersS3(t *testing.T) {
	type input struct {
		ctx          *cli.Context
		cfClient     *awsCF.CloudFormation
		manifestFile *manifest.Manifest
	}

	tests := []struct {
		name   string
		input  input
		output []*awsCF.Parameter
	}{
		{
			name: "Dev",
			input: input{
				ctx: func() *cli.Context {
					var context *cli.Context

					app := cli.NewApp()

					app.Flags = []cli.Flag{
						cli.StringFlag{
							Name:  "environment, e",
							Usage: "environment config to use from ./kombustion.yaml",
						},
					}

					app.Action = func(c *cli.Context) error {
						context = c
						return nil
					}

					app.Run([]string{
						"",
						"--environment", "development",
					})
					return context
				}(),
				manifestFile: &manifest.Manifest{
					Name:               "TestManifestWithEnvironment",
					Plugins:            nil,
					Architectures:      []string(nil),
					GenerateDefaultOutputs: false,
					Environments: map[string]manifest.Environment{
						"development": {
							AccountIDs: nil,
							Parameters: map[string]string{
								"parameterOneName":   "parameterOneValue",
								"parameterTwoName":   "8654238642489624862",
								"parameterThreeName": "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s",
								"parameterFourName":  "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-",
							},
						},
					},
				},
			},
			output: []*awsCF.Parameter{
				{
					ParameterKey:   aws.String("parameterOneName"),
					ParameterValue: aws.String("parameterOneValue"),
				},
				{
					ParameterKey:   aws.String("parameterTwoName"),
					ParameterValue: aws.String("8654238642489624862"),
				},
				{
					ParameterKey:   aws.String("parameterThreeName"),
					ParameterValue: aws.String("3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s"),
				},
				{
					ParameterKey:   aws.String("parameterFourName"),
					ParameterValue: aws.String("hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-"),
				},
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput := ResolveParametersS3(
			test.input.ctx,
			test.input.manifestFile,
		)

		matches := 0
		for _, outputParam := range testOutput {
			for _, param := range test.output {
				if *outputParam.ParameterKey == *param.ParameterKey &&
					*outputParam.ParameterValue == *param.ParameterValue {
					matches = matches + 1
				}
			}
		}

		assert.Equal(len(test.output), matches, fmt.Sprintf("Test %d: %s", i, test.name))
	}
}

// func ResolveParameters(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		input struct {
// 			ctx          *cli.Context
// 			cfYaml       YamlCloudformation
// 			cfClient     *awsCF.CloudFormation
// 			manifestFile *manifest.Manifest
// 		}
// 		output []*awsCF.Parameter
// 		throws bool
// 	}{
// 		{
// 			name:        "Dev",
// 			environment: "development",
// 			output:      map[string]string{},
// 			throws:      false,
// 		},
// 	}

// 	for i, test := range tests {
// 		assert := assert.New(t)
// 		testOutput, err := ResolveParameters(
// 			test.input.ctx,
// 			test.input.cfYaml,
// 			test.input.cfClient,
// 			test.input.manifestFile,
// 		)
// 		if test.throws {
// 			assert.NotNil(err)
// 		} else {
// 			assert.Nil(err)
// 			assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d: %s", i, test.name))
// 		}
// 	}
// }
