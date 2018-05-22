package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type InspectorResourceGroup struct {
	Type       string                      `yaml:"Type"`
	Properties InspectorResourceGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type InspectorResourceGroupProperties struct {
	ResourceGroupTags interface{} `yaml:"ResourceGroupTags"`
}

func NewInspectorResourceGroup(properties InspectorResourceGroupProperties, deps ...interface{}) InspectorResourceGroup {
	return InspectorResourceGroup{
		Type:       "AWS::Inspector::ResourceGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseInspectorResourceGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource InspectorResourceGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: InspectorResourceGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource InspectorResourceGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource InspectorResourceGroupProperties) Validate() []error {
	errs := []error{}
	if resource.ResourceGroupTags == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ResourceGroupTags'"))
	}
	return errs
}
