package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type EC2InternetGateway struct {
	Type       string                      `yaml:"Type"`
	Properties EC2InternetGatewayProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2InternetGatewayProperties struct {
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewEC2InternetGateway(properties EC2InternetGatewayProperties, deps ...interface{}) EC2InternetGateway {
	return EC2InternetGateway{
		Type:       "AWS::EC2::InternetGateway",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2InternetGateway(name string, data string) (cf types.ValueMap, err error) {
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
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2InternetGateway) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2InternetGatewayProperties) Validate() []error {
	errs := []error{}
	return errs
}
