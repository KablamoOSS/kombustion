package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type EC2PlacementGroup struct {
	Type       string                      `yaml:"Type"`
	Properties EC2PlacementGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2PlacementGroupProperties struct {
	Strategy interface{} `yaml:"Strategy,omitempty"`
}

func NewEC2PlacementGroup(properties EC2PlacementGroupProperties, deps ...interface{}) EC2PlacementGroup {
	return EC2PlacementGroup{
		Type:       "AWS::EC2::PlacementGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2PlacementGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2PlacementGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2PlacementGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2PlacementGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2PlacementGroupProperties) Validate() []error {
	errs := []error{}
	return errs
}
