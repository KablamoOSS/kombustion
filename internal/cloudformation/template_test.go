package cloudformation

import (
	"testing"
	"github.com/KablamoOSS/kombustion/types"
	"github.com/stretchr/testify/assert"
	"github.com/KablamoOSS/kombustion/internal/plugins"
	log "github.com/sirupsen/logrus"
	"bytes"
	"github.com/KablamoOSS/kombustion/pkg/parsers/resources"
	"io/ioutil"
	"os"
	"fmt"
)

type (
	testYamlTemplateCF struct {
		tstRes  types.ResourceMap
		expcRes types.TemplateObject
		isRes   bool
		log     string
		errMsg  string
	}

	testGenerateYamlStack struct {
		genParams GenerateParams
		errMsg    string
		expect    YamlCloudformation
		file      *os.File
	}
)

func Test_GenerateYamlStack(t *testing.T) {

	tests := make([]testGenerateYamlStack, 10)

	// 0 - YAML File path invalid
	tests[0].expect = YamlCloudformation{}
	tests[0].errMsg = "open : no such file or directory"
	tests[0].genParams = GenerateParams{
		Plugins:            []*plugins.PluginLoaded{},
		GenerateDefaultOutputs: true,
		Filename:           "",
	}

	// 1 - Empty YAML file
	tests[1].file = NewTempFile("")
	tests[1].expect = YamlCloudformation{
		Mappings:  types.TemplateObject{},
		Resources: types.TemplateObject{},
		Outputs:   types.TemplateObject{},
	}
	tests[1].errMsg = ""
	tests[1].genParams = GenerateParams{
		Plugins:            []*plugins.PluginLoaded{},
		GenerateDefaultOutputs: false,
		Filename:           tests[1].file.Name(),
	}

	// 2 - YAML File Invalid format
	tests[2].file = NewTempFile("INVALID")
	tests[2].expect = YamlCloudformation{}
	tests[2].errMsg = "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `INVALID` into cloudformation.YamlConfig"
	tests[2].genParams = GenerateParams{
		Plugins:            []*plugins.PluginLoaded{},
		GenerateDefaultOutputs: false,
		Filename:           tests[2].file.Name(),
	}

	// 3 - YAML basic header
	tests[3].file = NewTempFile("AWSTemplateFormatVersion: 2010-09-09\n" +
		"Description: test")
	tests[3].expect = YamlCloudformation{
		AWSTemplateFormatVersion: "2010-09-09",
		Description:              "test",
		Mappings:                 types.TemplateObject{},
		Resources:                types.TemplateObject{},
		Outputs:                  types.TemplateObject{},
	}
	tests[3].errMsg = ""
	tests[3].genParams = GenerateParams{
		Plugins:            []*plugins.PluginLoaded{},
		GenerateDefaultOutputs: false,
		Filename:           tests[3].file.Name(),
		ParamMap:           map[string]string{},
	}

	// 4 - YAML basic header invalid template
	tests[4].file = NewTempFile("AWSTemplateFormatVersion: 2010-09-09\n" +
		"Description: {{INVALID}}")
	tests[4].expect = YamlCloudformation{}
	tests[4].errMsg = "template: cfn:2: function \"INVALID\" not defined"
	tests[4].genParams = GenerateParams{
		Plugins:            []*plugins.PluginLoaded{},
		GenerateDefaultOutputs: false,
		Filename:           tests[4].file.Name(),
		ParamMap:           map[string]string{"p1": "test"},
	}

	// 5 - YAML basic header ok template
	tests[5].file = NewTempFile("AWSTemplateFormatVersion: 2010-09-09\n" +
		"Description: {{.p1}}")
	tests[5].expect = YamlCloudformation{
		AWSTemplateFormatVersion: "2010-09-09",
		Description:              "test",
		Mappings:                 types.TemplateObject{},
		Resources:                types.TemplateObject{},
		Outputs:                  types.TemplateObject{},
	}
	tests[5].errMsg = ""
	tests[5].genParams = GenerateParams{
		Plugins:            []*plugins.PluginLoaded{},
		GenerateDefaultOutputs: false,
		Filename:           tests[5].file.Name(),
		ParamMap:           map[string]string{"p1": "test"},
	}

	// 6 - YAML basic header plus IAM resource missing field
	tests[6].file = NewTempFile("AWSTemplateFormatVersion: 2010-09-09\n" +
		"Description: test\n" +
		"Resources:\n" +
		"  testRole:\n" +
		"    Type: AWS::IAM::Role")
	tests[6].expect = YamlCloudformation{}
	tests[6].errMsg = "Missing required field 'AssumeRolePolicyDocument'"
	tests[6].genParams = GenerateParams{
		Plugins:            []*plugins.PluginLoaded{},
		GenerateDefaultOutputs: false,
		Filename:           tests[6].file.Name(),
	}

	// 7 - YAML basic header plus IAM resource
	tests[7].file = NewTempFile("AWSTemplateFormatVersion: 2010-09-09\n" +
		"Description: test\n" +
		"Resources:\n" +
		"  testRole:\n" +
		"    Type: AWS::IAM::Role\n" +
		"    Properties:\n" +
		"      AssumeRolePolicyDocument:\n" +
		"        Statement:" +
		"")
	tests[7].expect = YamlCloudformation{
		AWSTemplateFormatVersion: "2010-09-09",
		Description:              "test",
		Mappings:                 types.TemplateObject{},
		Resources: types.TemplateObject{
			"testRole": resources.IAMRole{
				Type: "AWS::IAM::Role",
				Properties: resources.IAMRoleProperties{
					RoleName:                 nil,
					AssumeRolePolicyDocument: map[interface{}]interface{}{"Statement": nil},
				},
			},
		},
		Outputs: types.TemplateObject{},
	}
	tests[7].errMsg = ""
	tests[7].genParams = GenerateParams{
		Plugins:            []*plugins.PluginLoaded{},
		GenerateDefaultOutputs: true,
		Filename:           tests[7].file.Name(),
	}

	// 8 - YAML basic header with Mappings
	tests[8].file = NewTempFile("AWSTemplateFormatVersion: 2010-09-09\n" +
		"Description: test\n" +
		"Mappings:\n" +
		"  Mapping01:\n" +
		"   Key01:\n" +
		"     Name01: Value01\n")
	tests[8].expect = YamlCloudformation{
		AWSTemplateFormatVersion: "2010-09-09",
		Description:              "test",
		Mappings: types.TemplateObject{
			"Mapping01": map[interface{}]interface{}{
				"Key01": map[interface{}]interface{}{
					"Name01": "Value01",
				},
			},
		},
		Resources: types.TemplateObject{},
		Outputs:   types.TemplateObject{},
	}
	tests[8].errMsg = ""
	tests[8].genParams = GenerateParams{
		Plugins:            []*plugins.PluginLoaded{},
		GenerateDefaultOutputs: false,
		Filename:           tests[8].file.Name(),
	}

	// 9 - YAML basic header with Outputs
	tests[9].file = NewTempFile("AWSTemplateFormatVersion: 2010-09-09\n" +
		"Outputs:\n" +
		"  Logical ID:\n" +
		"    Description: Information about the value\n" +
		"    Value: Value to return\n" +
		"    Export:\n" +
		"      Name: Value to export\n")
	tests[9].expect = YamlCloudformation{
		AWSTemplateFormatVersion: "2010-09-09",
		Mappings:                 types.TemplateObject{},
		Resources:                types.TemplateObject{},
		Outputs: types.TemplateObject{
			"Logical ID": map[interface{}]interface{}{
				"Description": "Information about the value",
				"Value":       "Value to return",
				"Export": map[interface{}]interface{}{
					"Name": "Value to export",
				},
			},
		},
	}
	tests[9].errMsg = ""
	tests[9].genParams = GenerateParams{
		Plugins:            []*plugins.PluginLoaded{},
		GenerateDefaultOutputs: false,
		Filename:           tests[9].file.Name(),
	}

	for idx, test := range tests {

		yaml, err := GenerateYamlStack(test.genParams)

		if test.errMsg != "" {
			assert.EqualError(t, err, test.errMsg, fmt.Sprintf("Test %d", idx))
		} else {
			assert.Equal(t, nil, err, fmt.Sprintf("Test %d", idx))
		}
		assert.EqualValues(t, test.expect, yaml, fmt.Sprintf("Test %d", idx))

		if test.file != nil {
			test.file.Close()
			os.Remove(test.file.Name())
		}

	}

}

func Test_YamlTemplateCF(t *testing.T) {

	memLog := &bytes.Buffer{}
	log.SetOutput(memLog)

	populateParsers([]*plugins.PluginLoaded{}, false)

	resCustom1 := types.CfResource{Type: "AWS::CloudFormation::CustomResource",}
	resCustom2 := types.CfResource{Type: "Custom::",}
	resUnknown := types.CfResource{Type: "AWS::UNKNOWN::RES",}

	tests := make([]testYamlTemplateCF, 5)

	tests[0].expcRes = types.TemplateObject{"customRole": resCustom1,}
	tests[0].tstRes = types.ResourceMap{"customRole": resCustom1,}
	tests[0].isRes = true
	tests[0].errMsg = ""

	tests[1].expcRes = types.TemplateObject{"customRole": resCustom2,}
	tests[1].tstRes = types.ResourceMap{"customRole": resCustom2,}
	tests[1].isRes = true
	tests[1].errMsg = ""

	tests[2].expcRes = types.TemplateObject{}
	tests[2].tstRes = types.ResourceMap{"unknownRole": resUnknown,}
	tests[2].isRes = true
	tests[2].log = "msg=\"Type not found\" type=\"AWS::UNKNOWN::RES\""
	tests[2].errMsg = ""

	tests[3].tstRes, tests[3].expcRes = FullImRole("fullRole", false)
	tests[3].isRes = true
	tests[3].log = ""
	tests[3].errMsg = ""

	tests[4].tstRes, tests[4].expcRes = FullImRole("conditionRole", true)
	tests[4].isRes = true
	tests[4].log = "msg=\"Condition being applied on resource, this is not yet supported\" resource=conditionRole"
	tests[4].errMsg = ""

	for _, test := range tests {

		compiledResources, err := yamlTemplateCF(test.tstRes, resourceParsers, test.isRes)

		if test.errMsg != "" {
			assert.EqualError(t, err, test.errMsg)
		} else {
			assert.Equal(t, nil, err)
		}
		assert.Contains(t, memLog.String(), test.log)
		assert.EqualValues(t, test.expcRes, compiledResources)

		memLog.Reset()
	}

}

func NewTempFile(content string) (file *os.File) {

	file, err := ioutil.TempFile("", "test")

	if err != nil {
		log.Fatal(err)
	}

	if _, err := file.Write([]byte(content)); err != nil {
		log.Fatal(err)
	}

	return file

}

func FullImRole(roleKey string, conditional bool) (res types.ResourceMap, tmpl types.TemplateObject) {

	assumeRolePolicyDocument := "testAssumeRolePolicyDocument"
	deps := []interface{}{"dep1", "dep2"}

	role := types.CfResource{
		Type: "AWS::IAM::Role",
		Properties: map[string]string{
			"RoleName":                 roleKey,
			"AssumeRolePolicyDocument": assumeRolePolicyDocument,
		},
		DependsOn: deps,
		Condition: conditional,
	}

	iamRole := resources.NewIAMRole(
		resources.IAMRoleProperties{
			RoleName:                 roleKey,
			AssumeRolePolicyDocument: assumeRolePolicyDocument,
		},
		deps...,
	)

	if conditional {
		iamRole.Condition = conditional
	}

	res = types.ResourceMap{
		roleKey: role,
	}

	tmpl = types.TemplateObject{
		roleKey: iamRole,
	}

	return res, tmpl

}
