// +build plugin

package resources

import (
	"log"

	"github.com/KablamoOSS/kombustion/pluginParsers/resources"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

type LambdaPermissionConfig struct {
	Properties struct {
		Action           interface{} `yaml:"Action,omitempty"`
		EventSourceToken interface{} `yaml:"EventSourceToken,omitempty"`
		FunctionName     interface{} `yaml:"FunctionName"`
		Principal        interface{} `yaml:"Principal,omitempty"`
		SourceAccount    interface{} `yaml:"SourceAccount,omitempty"`
		SourceArn        interface{} `yaml:"SourceArn,omitempty"`
		SourceApiGateway interface{} `yaml:"SourceApiGateway,omitempty"`
	} `yaml:"Properties"`
}

func ParseLambdaPermission(name string, data string) (cf types.ValueMap, err error) {
	// Parse the config data
	var config LambdaPermissionConfig
	if err = yaml.Unmarshal([]byte(data), &config); err != nil {
		return
	}

	// validate the config
	config.Validate()

	action := config.Properties.Action
	if action == nil {
		action = "lambda:InvokeFunction"
	}

	principal := config.Properties.Principal
	if principal == nil {
		principal = "apigateway.amazonaws.com"
	}

	sourceArn := config.Properties.SourceArn
	if sourceArn == nil {
		if config.Properties.SourceApiGateway != nil {
			sourceArn = map[string]interface{}{
				"Fn::Join": []interface{}{
					"", []interface{}{
						"arn:aws:execute-api:",
						map[string]string{"Ref": "AWS::Region"},
						":",
						map[string]string{"Ref": "AWS::AccountId"},
						":",
						config.Properties.SourceApiGateway,
						"/*",
					},
				},
			}
		}
	}

	cf = types.ValueMap{
		(name): resources.NewLambdaPermission(
			resources.LambdaPermissionProperties{
				Action:           action,
				Principal:        principal,
				SourceArn:        sourceArn,
				FunctionName:     config.Properties.FunctionName,
				SourceAccount:    config.Properties.SourceAccount,
				EventSourceToken: config.Properties.EventSourceToken,
			},
		),
	}

	return
}

// Validate - input Config validation
func (this LambdaPermissionConfig) Validate() {
	props := this.Properties
	if props.FunctionName == nil {
		log.Println("WARNING: LambdaPermissionConfig - Missing required field 'FunctionName'")
	}
	if props.SourceArn == nil && props.SourceApiGateway == nil {
		log.Println("WARNING: LambdaPermissionConfig - Must provide one of: 'SourceArn', 'SourceApiGateway'")
	}
}
