// +build plugin

package resources

import (
	"log"

	"github.com/KablamoOSS/kombustion/pluginParsers/properties"
	"github.com/KablamoOSS/kombustion/pluginParsers/resources"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

type ApiGatewayConfig struct {
	Properties struct {
		Description    string      `yaml:"Description,omitempty"`
		FailOnWarnings bool        `yaml:"FailOnWarnings"`
		Name           interface{} `yaml:"Name"`
		DeploymentName interface{} `yaml:"DeploymentName"`
		StageName      interface{} `yaml:"StageName"`
		Resources      map[string]struct {
			PathPart string `yaml:"PathPart"`
			Parent   string `yaml:"Description"`
		} `yaml:"Resources"`
		Methods map[string]struct {
			LambdaArn         interface{} `yaml:"LambdaArn"`
			ApiKeyRequired    bool        `yaml:"ApiKeyRequired,omitempty"`
			AuthorizationType interface{} `yaml:"AuthorizationType,omitempty"`
			AuthorizerId      interface{} `yaml:"AuthorizerId,omitempty"`
			HttpMethod        interface{} `yaml:"HttpMethod"`
			ResourceId        interface{} `yaml:"ResourceId"`
		} `yaml:"Methods"`
	} `yaml:"Properties"`
}

func ParseApiGateway(name string, data string) (cf types.ValueMap, err error) {
	var config ApiGatewayConfig
	if err = yaml.Unmarshal([]byte(data), &config); err != nil {
		return
	}

	config.Validate()

	restApiId := map[string]string{"Ref": name}
	methodNames := make([]interface{}, 0, len(config.Properties.Methods))
	for m := range config.Properties.Methods {
		methodNames = append(methodNames, m)
	}

	objects := map[string]types.Validatable{
		(name): resources.NewApiGatewayRestApi(
			resources.ApiGatewayRestApiProperties{
				Description:    config.Properties.Description,
				FailOnWarnings: config.Properties.FailOnWarnings,
				Name:           config.Properties.Name,
			},
		),
		(name + "Deployment"): resources.NewApiGatewayDeployment(
			resources.ApiGatewayDeploymentProperties{
				RestApiId: restApiId,
				StageName: config.Properties.StageName,
			},
			methodNames..., // deps
		),
	}

	for k, v := range config.Properties.Resources {
		parentID := map[string]interface{}{"Fn::GetAtt": []string{name, "RootResourceId"}}
		if len(v.Parent) > 0 {
			parentID = map[string]interface{}{"Ref": v.Parent}
		}

		objects[k] = resources.NewApiGatewayResource(
			resources.ApiGatewayResourceProperties{
				RestApiId: restApiId,
				PathPart:  v.PathPart,
				ParentId:  parentID,
			},
		)
	}

	for k, method := range config.Properties.Methods {
		uri := map[string]interface{}{
			"Fn::Join": []interface{}{
				"", []interface{}{
					"arn:aws:apigateway:",
					map[string]string{"Ref": "AWS::Region"},
					":lambda:path/2015-03-31/functions/",
					method.LambdaArn,
					"/invocations",
				},
			},
		}

		objects[k] = resources.NewApiGatewayMethod(
			resources.ApiGatewayMethodProperties{
				RestApiId:         restApiId,
				ResourceId:        method.ResourceId,
				HttpMethod:        method.HttpMethod,
				ApiKeyRequired:    method.ApiKeyRequired,
				AuthorizationType: method.AuthorizationType,
				AuthorizerId:      method.AuthorizerId,
				Integration: &properties.Method_Integration{
					IntegrationHttpMethod: "POST",
					Type: "AWS_PROXY",
					Uri:  uri,
				},
			},
		)
	}

	cf = make(types.ValueMap)
	for k, resource := range objects {
		// validate resource
		if errs := resource.Validate(); len(errs) > 0 {
			for _, err = range errs {
				log.Println("WARNING: apigateway - ", err)
			}
			return
		}
		// add resource to output
		cf[k] = resource
	}
	return
}

// Validate - input Config validation
func (this ApiGatewayConfig) Validate() {
	if this.Properties.Name == nil {
		log.Println("WARNING: ApiGatewayConfig - Missing required field 'Name'")
	}
	if this.Properties.DeploymentName == nil {
		log.Println("WARNING: ApiGatewayConfig - Missing required field 'DeploymentName'")
	}
	if this.Properties.StageName == nil {
		log.Println("WARNING: ApiGatewayConfig - Missing required field 'StageName'")
	}
}
