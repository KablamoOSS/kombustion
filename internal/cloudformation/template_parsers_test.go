package cloudformation

import (
	"fmt"
	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/pkg/parsers/resources"
	"github.com/KablamoOSS/kombustion/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
