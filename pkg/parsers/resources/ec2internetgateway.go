package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// EC2InternetGateway Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-internetgateway.html
type EC2InternetGateway struct {
	Type       string                       `yaml:"Type"`
	Properties EC2InternetGatewayProperties `yaml:"Properties"`
	Condition  interface{}                  `yaml:"Condition,omitempty"`
	Metadata   interface{}                  `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                  `yaml:"DependsOn,omitempty"`
}

// EC2InternetGateway Properties
type EC2InternetGatewayProperties struct {
	Tags interface{} `yaml:"Tags,omitempty"`
}

// NewEC2InternetGateway constructor creates a new EC2InternetGateway
func NewEC2InternetGateway(properties EC2InternetGatewayProperties, deps ...interface{}) EC2InternetGateway {
	return EC2InternetGateway{
		Type:       "AWS::EC2::InternetGateway",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2InternetGateway parses EC2InternetGateway
func ParseEC2InternetGateway(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource EC2InternetGateway
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2InternetGateway - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource EC2InternetGateway) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2InternetGatewayProperties) Validate() []error {
	errs := []error{}
	return errs
}
