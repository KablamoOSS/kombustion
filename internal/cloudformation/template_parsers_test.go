package cloudformation

import (
	"fmt"
	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/pkg/parsers/resources"
	"github.com/KablamoOSS/kombustion/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateYamlTemplate(t *testing.T) {
	// Prevent the printer from exiting
	printer.Test()

	tests := []struct {
		input  GenerateParams
		output YamlCloudformation
		err    error
	}{
		{
			input: GenerateParams{
				Filename: "testdata/test.yaml",
			},
			output: YamlCloudformation{
				AWSTemplateFormatVersion: "2010-09-09",
				Description:              "A Demo Template for testing Kombustion",
				Metadata:                 types.TemplateObject{},
				Parameters: types.TemplateObject{
					"Environment": map[interface{}]interface{}{
						"Type": "String", "Default": "UnknownEnvironment",
					},
				},
				Mappings:   types.TemplateObject{},
				Conditions: types.TemplateObject{},
				Transform:  types.TemplateObject{},
				Resources:  types.TemplateObject{},
				Outputs:    types.TemplateObject{},
			},
			err: nil,
		},
		{
			input: GenerateParams{
				Filename:               "testdata/test.yaml",
				GenerateDefaultOutputs: true,
			},
			output: YamlCloudformation{
				AWSTemplateFormatVersion: "2010-09-09",
				Description:              "A Demo Template for testing Kombustion",
				Metadata:                 types.TemplateObject{},
				Parameters: types.TemplateObject{
					"Environment": map[interface{}]interface{}{
						"Type": "String", "Default": "UnknownEnvironment",
					},
				},
				Mappings:   types.TemplateObject{},
				Conditions: types.TemplateObject{},
				Transform:  types.TemplateObject{},
				Resources:  types.TemplateObject{},
				Outputs: types.TemplateObject{
					"MyDemoLogGroup3": types.TemplateObject{
						"Description": "MyDemoLogGroup3 Object",
						"Value": map[string]interface{}{
							"Ref": "MyDemoLogGroup3",
						},
						"Export": map[string]interface{}{
							"Name": map[string]interface{}{
								"Fn::Sub": "${AWS::StackName}-LogsLogGroup-MyDemoLogGroup3",
							},
						},
					},
					"MyDemoLogGroup4": types.TemplateObject{
						"Description": "MyDemoLogGroup4 Object",
						"Value": map[string]interface{}{
							"Ref": "MyDemoLogGroup4",
						},
						"Export": map[string]interface{}{
							"Name": map[string]interface{}{
								"Fn::Sub": "${AWS::StackName}-LogsLogGroup-MyDemoLogGroup4",
							},
						},
					},
					"MyDemoLogGroup4Arn": types.TemplateObject{
						"Value": map[string]interface{}{"Fn::GetAtt": []string{"MyDemoLogGroup4", "Arn"},
						},
						"Export": map[string]interface{}{
							"Name": map[string]interface{}{
								"Fn::Sub": "${AWS::StackName}-LogsLogGroup-MyDemoLogGroup4-Arn",
							},
						},
						"Description": "MyDemoLogGroup4 Object",
					},
					"MyDemoLogGroup": types.TemplateObject{
						"Export": map[string]interface{}{
							"Name": map[string]interface{}{
								"Fn::Sub": "${AWS::StackName}-LogsLogGroup-MyDemoLogGroup",
							},
						},
						"Description": "MyDemoLogGroup Object",
						"Value": map[string]interface{}{
							"Ref": "MyDemoLogGroup",
						},
					},
					"MyDemoLogGroupArn": types.TemplateObject{
						"Description": "MyDemoLogGroup Object",
						"Value": map[string]interface{}{
							"Fn::GetAtt": []string{"MyDemoLogGroup", "Arn"},
						},
						"Export": map[string]interface{}{
							"Name": map[string]interface{}{
								"Fn::Sub": "${AWS::StackName}-LogsLogGroup-MyDemoLogGroup-Arn",
							},
						},
					},
					"MyDemoLogGroup2": types.TemplateObject{
						"Description": "MyDemoLogGroup2 Object",
						"Value": map[string]interface{}{
							"Ref": "MyDemoLogGroup2",
						},
						"Export": map[string]interface{}{
							"Name": map[string]interface{}{
								"Fn::Sub": "${AWS::StackName}-LogsLogGroup-MyDemoLogGroup2",
							},
						},
					},
					"MyDemoLogGroup2Arn": types.TemplateObject{
						"Description": "MyDemoLogGroup2 Object",
						"Value": map[string]interface{}{
							"Fn::GetAtt": []string{"MyDemoLogGroup2", "Arn"},
						},
						"Export": map[string]interface{}{
							"Name": map[string]interface{}{
								"Fn::Sub": "${AWS::StackName}-LogsLogGroup-MyDemoLogGroup2-Arn",
							},
						},
					},
					"MyDemoLogGroup3Arn": types.TemplateObject{"Description": "MyDemoLogGroup3 Object",
						"Value": map[string]interface{}{
							"Fn::GetAtt": []string{"MyDemoLogGroup3", "Arn"},
						},
						"Export": map[string]interface{}{
							"Name": map[string]interface{}{
								"Fn::Sub": "${AWS::StackName}-LogsLogGroup-MyDemoLogGroup3-Arn",
							},
						},
					},
				},
			},
			err: nil,
		},
	}

	for i, test := range tests {
		assert := assert.New(t)

		output, err := GenerateYamlTemplate(
			test.input,
		)

		assert.Equal(
			test.err,
			err,
			fmt.Sprintf("Test: %d", i),
		)

		assert.Equal(
			test.output,
			output,
			fmt.Sprintf("Test: %d", i),
		)
	}
}

func TestProcessParsers(t *testing.T) {
	// Prevent the printer from exiting
	printer.Test()

	type Input struct {
		templateResources types.ResourceMap
		parserFuncs       map[string]types.ParserFunc
	}
	type Output struct {
		conditions types.TemplateObject
		metadata   types.TemplateObject
		mappings   types.TemplateObject
		outputs    types.TemplateObject
		parameters types.TemplateObject
		resources  types.TemplateObject
		transform  types.TemplateObject
	}

	tests := []struct {
		input  Input
		output Output
	}{
		{
			input: Input{
				templateResources: types.ResourceMap{
					"Test": types.CfResource{
						Type: "AWS::Logs::LogsGroup",
						Properties: map[string]interface{}{
							"LogGroupName": "TestLogGroup",
						},
						Condition: map[string]interface{}{
							"ConditionName": "ConditionValue",
						},
						Metadata: map[string]interface{}{
							"MetadataName": "MetadataValue",
						},
						DependsOn: []interface{}{
							"OtherResource",
						},
					},
				},
				parserFuncs: map[string]types.ParserFunc{
					"AWS::Logs::LogsGroup": resources.ParseLogsLogGroup,
				},
			},
			output: Output{
				conditions: types.TemplateObject{},
				metadata:   types.TemplateObject{},
				mappings:   types.TemplateObject{},
				outputs:    types.TemplateObject{},
				parameters: types.TemplateObject{},
				resources: types.TemplateObject{
					"Test": resources.LogsLogGroup{
						Type: "AWS::Logs::LogsGroup",
						Properties: resources.LogsLogGroupProperties{
							LogGroupName:    "TestLogGroup",
							RetentionInDays: interface{}(nil)},
						Condition: map[interface{}]interface{}{
							"ConditionName": "ConditionValue",
						},
						Metadata: map[interface{}]interface{}{
							"MetadataName": "MetadataValue",
						},
						DependsOn: []interface{}{"OtherResource"},
					},
				},
				transform: types.TemplateObject{},
			},
		},
		{
			input: Input{
				templateResources: types.ResourceMap{
					"Test": types.CfResource{
						Type: "AWS::CloudFormation::CustomResource",
						Properties: map[string]interface{}{
							"CustomResource": "CustomResource",
						},
						Condition: map[string]interface{}{
							"CustomResource": "CustomResource",
						},
						Metadata: map[string]interface{}{
							"CustomResource": "CustomResource",
						},
						DependsOn: []interface{}{
							"CustomResource",
						},
					},
				},
				parserFuncs: map[string]types.ParserFunc{
					"AWS::Logs::LogsGroup": resources.ParseLogsLogGroup,
				},
			},
			output: Output{
				conditions: types.TemplateObject{},
				metadata:   types.TemplateObject{},
				mappings:   types.TemplateObject{},
				outputs:    types.TemplateObject{},
				parameters: types.TemplateObject{},
				resources: types.TemplateObject{
					"Test": types.CfResource{
						Type: "AWS::CloudFormation::CustomResource",
						Properties: map[string]interface{}{
							"CustomResource": "CustomResource",
						},
						Condition: map[string]interface{}{
							"CustomResource": "CustomResource",
						}, Metadata: map[string]interface{}{
							"CustomResource": "CustomResource",
						},
						DependsOn: []interface{}{"CustomResource"},
					},
				},
				transform: types.TemplateObject{},
			},
		},
		{
			input: Input{
				templateResources: types.ResourceMap{
					"Test": types.CfResource{
						Type: "Custom::",
						Properties: map[string]interface{}{
							"CustomResource": "CustomResource",
						},
						Condition: map[string]interface{}{
							"CustomResource": "CustomResource",
						},
						Metadata: map[string]interface{}{
							"CustomResource": "CustomResource",
						},
						DependsOn: []interface{}{
							"CustomResource",
						},
					},
				},
				parserFuncs: map[string]types.ParserFunc{},
			},
			output: Output{
				conditions: types.TemplateObject{},
				metadata:   types.TemplateObject{},
				mappings:   types.TemplateObject{},
				outputs:    types.TemplateObject{},
				parameters: types.TemplateObject{},
				resources: types.TemplateObject{
					"Test": types.CfResource{
						Type: "Custom::",
						Properties: map[string]interface{}{
							"CustomResource": "CustomResource",
						},
						Condition: map[string]interface{}{
							"CustomResource": "CustomResource",
						}, Metadata: map[string]interface{}{
							"CustomResource": "CustomResource",
						},
						DependsOn: []interface{}{"CustomResource"},
					},
				},
				transform: types.TemplateObject{},
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)

		conditions,
			metadata,
			mappings,
			outputs,
			parameters,
			resources,
			transform := processParsers(
			test.input.templateResources,
			test.input.parserFuncs,
		)

		assert.Equal(
			test.output.conditions,
			conditions,
			fmt.Sprintf("Test: %d", i),
		)

		assert.Equal(
			test.output.metadata,
			metadata,
			fmt.Sprintf("Test: %d", i),
		)

		assert.Equal(
			test.output.mappings,
			mappings,
			fmt.Sprintf("Test: %d", i),
		)

		assert.Equal(
			test.output.outputs,
			outputs,
			fmt.Sprintf("Test: %d", i),
		)

		assert.Equal(
			test.output.parameters,
			parameters,
			fmt.Sprintf("Test: %d", i),
		)

		assert.Equal(
			test.output.resources,
			resources,
			fmt.Sprintf("Test: %d", i),
		)

		assert.Equal(
			test.output.transform,
			transform,
			fmt.Sprintf("Test: %d", i),
		)
	}
}
