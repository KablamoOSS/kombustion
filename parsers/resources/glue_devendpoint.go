package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type GlueDevEndpoint struct {
	Type       string                      `yaml:"Type"`
	Properties GlueDevEndpointProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GlueDevEndpointProperties struct {
	EndpointName interface{} `yaml:"EndpointName,omitempty"`
	ExtraJarsS3Path interface{} `yaml:"ExtraJarsS3Path,omitempty"`
	ExtraPythonLibsS3Path interface{} `yaml:"ExtraPythonLibsS3Path,omitempty"`
	NumberOfNodes interface{} `yaml:"NumberOfNodes,omitempty"`
	PublicKey interface{} `yaml:"PublicKey"`
	RoleArn interface{} `yaml:"RoleArn"`
	SubnetId interface{} `yaml:"SubnetId,omitempty"`
	SecurityGroupIds interface{} `yaml:"SecurityGroupIds,omitempty"`
}

func NewGlueDevEndpoint(properties GlueDevEndpointProperties, deps ...interface{}) GlueDevEndpoint {
	return GlueDevEndpoint{
		Type:       "AWS::Glue::DevEndpoint",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGlueDevEndpoint(name string, data string) (cf types.ValueMap, err error) {
	var resource GlueDevEndpoint
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GlueDevEndpoint - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource GlueDevEndpoint) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GlueDevEndpointProperties) Validate() []error {
	errs := []error{}
	if resource.PublicKey == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PublicKey'"))
	}
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	return errs
}
