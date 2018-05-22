package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type GlueJob struct {
	Type       string                      `yaml:"Type"`
	Properties GlueJobProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GlueJobProperties struct {
	AllocatedCapacity interface{} `yaml:"AllocatedCapacity,omitempty"`
	DefaultArguments interface{} `yaml:"DefaultArguments,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	LogUri interface{} `yaml:"LogUri,omitempty"`
	MaxRetries interface{} `yaml:"MaxRetries,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	Role interface{} `yaml:"Role"`
	Command *properties.Job_JobCommand `yaml:"Command"`
	ExecutionProperty *properties.Job_ExecutionProperty `yaml:"ExecutionProperty,omitempty"`
	Connections *properties.Job_ConnectionsList `yaml:"Connections,omitempty"`
}

func NewGlueJob(properties GlueJobProperties, deps ...interface{}) GlueJob {
	return GlueJob{
		Type:       "AWS::Glue::Job",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGlueJob(name string, data string) (cf types.ValueMap, err error) {
	var resource GlueJob
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GlueJob - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource GlueJob) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GlueJobProperties) Validate() []error {
	errs := []error{}
	if resource.Role == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Role'"))
	}
	if resource.Command == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Command'"))
	} else {
		errs = append(errs, resource.Command.Validate()...)
	}
	return errs
}
