package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type CloudFormationStack struct {
	Type       string                      `yaml:"Type"`
	Properties CloudFormationStackProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CloudFormationStackProperties struct {
	TemplateURL interface{} `yaml:"TemplateURL"`
	TimeoutInMinutes interface{} `yaml:"TimeoutInMinutes,omitempty"`
	Parameters interface{} `yaml:"Parameters,omitempty"`
	NotificationARNs interface{} `yaml:"NotificationARNs,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewCloudFormationStack(properties CloudFormationStackProperties, deps ...interface{}) CloudFormationStack {
	return CloudFormationStack{
		Type:       "AWS::CloudFormation::Stack",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCloudFormationStack(name string, data string) (cf types.ValueMap, err error) {
	var resource CloudFormationStack
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CloudFormationStack - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CloudFormationStack) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CloudFormationStackProperties) Validate() []error {
	errs := []error{}
	if resource.TemplateURL == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TemplateURL'"))
	}
	return errs
}
