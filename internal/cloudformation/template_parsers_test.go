package cloudformation

import (
	"fmt"
	"testing"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/coretest"
	"github.com/KablamoOSS/kombustion/pkg/parsers/resources"
	"github.com/KablamoOSS/kombustion/types"
	"github.com/stretchr/testify/assert"
)

var testYaml = `AWSTemplateFormatVersion: "2010-09-09"
Description: A Demo Template for testing Kombustion

Parameters:
  Environment:
    Type: String
    Default: UnknownEnvironment

Resources:
  MyDemoLogGroup:
    Type: "AWS::Logs::LogGroup"
    Properties:
      LogGroupName: !Join [ '-', [ "MyDemoLogGroup1",!Ref Environment ] ]
      RetentionInDays: 1
  MyDemoLogGroup2:
    Type: "AWS::Logs::LogGroup"
    Properties:
      LogGroupName: !Join [ '-', [ "MyDemoLogGroup2",!Ref Environment ] ]
      RetentionInDays: 1
  MyDemoLogGroup3:
    Type: "AWS::Logs::LogGroup"
    Properties:
      LogGroupName: !Join [ '-', [ "MyDemoLogGroup3",!Ref Environment ] ]
      RetentionInDays: 1
  MyDemoLogGroup4:
    Type: "AWS::Logs::LogGroup"
    Properties:
      LogGroupName: !Join [ '-', [ "MyDemoLogGroup4",!Ref Environment ] ]
      RetentionInDays: 1
`

func TestGenerateYamlTemplate(t *testing.T) {
	// Prevent the printer from exiting
	printer.Test()

	objectStore := coretest.NewMockObjectStore()
	objectStore.Put([]byte(testYaml), "test.yaml")

	tests := []struct {
		input  GenerateParams
		output YamlCloudformation
		err    error
	}{
		{
			input: GenerateParams{
				ObjectStore: objectStore,
				Filename:    "test.yaml",
			},
			output: YamlCloudformation{
				AWSTemplateFormatVersion: "2010-09-09",
				Description:              "A Demo Template for testing Kombustion",
				Metadata:                 types.TemplateObject{},
				Parameters: types.TemplateObject{
					"Environment": map[interface{}]interface{}{
						"Type":    "String",
						"Default": "UnknownEnvironment",
					},
				},
				Mappings:   types.TemplateObject{},
				Conditions: types.TemplateObject{},
				Transform:  types.TemplateObject{},
				Resources: types.TemplateObject{
					"MyDemoLogGroup4": resources.LogsLogGroup{
						Type: "AWS::Logs::LogGroup",
						Properties: resources.LogsLogGroupProperties{
							LogGroupName: map[interface{}]interface{}{
								"Fn::Join": []interface{}{
									"-",
									[]interface{}{
										"MyDemoLogGroup4",
										map[interface{}]interface{}{
											"Ref": "Environment",
										},
									},
								},
							},
							RetentionInDays: 1,
						},
						Condition: interface{}(nil),
						Metadata:  interface{}(nil),
						DependsOn: interface{}(nil)},
					"MyDemoLogGroup": resources.LogsLogGroup{
						Type: "AWS::Logs::LogGroup",
						Properties: resources.LogsLogGroupProperties{
							LogGroupName: map[interface{}]interface{}{
								"Fn::Join": []interface{}{
									"-", []interface{}{
										"MyDemoLogGroup1",
										map[interface{}]interface{}{
											"Ref": "Environment",
										},
									},
								},
							},
							RetentionInDays: 1,
						},
						Condition: interface{}(nil),
						Metadata:  interface{}(nil),
						DependsOn: interface{}(nil)}, "MyDemoLogGroup2": resources.LogsLogGroup{
						Type: "AWS::Logs::LogGroup", Properties: resources.LogsLogGroupProperties{
							LogGroupName: map[interface{}]interface{}{
								"Fn::Join": []interface{}{
									"-", []interface{}{
										"MyDemoLogGroup2",
										map[interface{}]interface{}{
											"Ref": "Environment",
										},
									},
								},
							},
							RetentionInDays: 1,
						},
						Condition: interface{}(nil),
						Metadata:  interface{}(nil),
						DependsOn: interface{}(nil)}, "MyDemoLogGroup3": resources.LogsLogGroup{
						Type: "AWS::Logs::LogGroup",
						Properties: resources.LogsLogGroupProperties{
							LogGroupName: map[interface{}]interface{}{
								"Fn::Join": []interface{}{
									"-",
									[]interface{}{
										"MyDemoLogGroup3",
										map[interface{}]interface{}{
											"Ref": "Environment",
										},
									},
								},
							},
							RetentionInDays: 1,
						},
						Condition: interface{}(nil),
						Metadata:  interface{}(nil),
						DependsOn: interface{}(nil),
					},
				},
				Outputs: types.TemplateObject{},
			},
			err: nil,
		},
		{
			input: GenerateParams{
				ObjectStore:            objectStore,
				Filename:               "test.yaml",
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
				Resources: types.TemplateObject{
					"MyDemoLogGroup2": resources.LogsLogGroup{
						Type: "AWS::Logs::LogGroup",
						Properties: resources.LogsLogGroupProperties{
							LogGroupName: map[interface{}]interface{}{
								"Fn::Join": []interface{}{"-", []interface{}{
									"MyDemoLogGroup2", map[interface{}]interface{}{
										"Ref": "Environment",
									},
								},
								},
							},
							RetentionInDays: 1,
						},
						Condition: interface{}(nil),
						Metadata:  interface{}(nil),
						DependsOn: interface{}(nil),
					},
					"MyDemoLogGroup3": resources.LogsLogGroup{
						Type: "AWS::Logs::LogGroup",
						Properties: resources.LogsLogGroupProperties{
							LogGroupName: map[interface{}]interface{}{
								"Fn::Join": []interface{}{"-", []interface{}{
									"MyDemoLogGroup3",
									map[interface{}]interface{}{
										"Ref": "Environment",
									},
								},
								},
							},
							RetentionInDays: 1,
						},
						Condition: interface{}(nil),
						Metadata:  interface{}(nil),
						DependsOn: interface{}(nil),
					},
					"MyDemoLogGroup4": resources.LogsLogGroup{
						Type: "AWS::Logs::LogGroup",
						Properties: resources.LogsLogGroupProperties{
							LogGroupName: map[interface{}]interface{}{
								"Fn::Join": []interface{}{
									"-",
									[]interface{}{
										"MyDemoLogGroup4",
										map[interface{}]interface{}{"Ref": "Environment"},
									},
								},
							},
							RetentionInDays: 1,
						},
						Condition: interface{}(nil),
						Metadata:  interface{}(nil),
						DependsOn: interface{}(nil)},
					"MyDemoLogGroup": resources.LogsLogGroup{
						Type: "AWS::Logs::LogGroup",
						Properties: resources.LogsLogGroupProperties{
							LogGroupName: map[interface{}]interface{}{
								"Fn::Join": []interface{}{"-",
									[]interface{}{"MyDemoLogGroup1",
										map[interface{}]interface{}{"Ref": "Environment"},
									},
								},
							},
							RetentionInDays: 1,
						},
						Condition: interface{}(nil),
						Metadata:  interface{}(nil),
						DependsOn: interface{}(nil),
					},
				},
				Outputs: types.TemplateObject{
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
						"Export": map[string]interface{}{
							"Name": map[string]interface{}{
								"Fn::Sub": "${AWS::StackName}-LogsLogGroup-MyDemoLogGroup4",
							},
						}, "Description": "MyDemoLogGroup4 Object",
						"Value": map[string]interface{}{
							"Ref": "MyDemoLogGroup4",
						},
					},
					"MyDemoLogGroup": types.TemplateObject{
						"Description": "MyDemoLogGroup Object",
						"Value": map[string]interface{}{
							"Ref": "MyDemoLogGroup",
						},
						"Export": map[string]interface{}{
							"Name": map[string]interface{}{
								"Fn::Sub": "${AWS::StackName}-LogsLogGroup-MyDemoLogGroup",
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
			false,
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
