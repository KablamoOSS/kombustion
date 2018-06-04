package cloudformation

import (
	"testing"

	"github.com/KablamoOSS/kombustion/parsers/resources"
	"github.com/KablamoOSS/kombustion/types"
	"github.com/stretchr/testify/assert"
)

func TestTemplateYamlCF_iamRole_success(t *testing.T) {
	rolekey := "testRole"
	assumeRolePolicyDocument := "testAssumeRolePolicyDocument"
	deps := []interface{}{"dep1", "dep2"}
	role := types.CfResource{
		Type: "AWS::IAM::Role",
		Properties: map[string]string{
			"RoleName":                 rolekey,
			"AssumeRolePolicyDocument": assumeRolePolicyDocument,
		},
		DependsOn: deps,
	}
	testResources := types.ResourceMap{
		rolekey: role,
	}

	expectedResources := TemplateObject{
		rolekey: resources.NewIAMRole(
			resources.IAMRoleProperties{
				RoleName:                 rolekey,
				AssumeRolePolicyDocument: assumeRolePolicyDocument,
			},
			deps...,
		),
	}

	populateParsers(false)
	compiledResources, err := yamlTemplateCF(testResources, resourceParsers, true)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedResources, compiledResources)
}
