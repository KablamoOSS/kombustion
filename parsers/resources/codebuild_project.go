package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type CodeBuildProject struct {
	Type       string                      `yaml:"Type"`
	Properties CodeBuildProjectProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CodeBuildProjectProperties struct {
	BadgeEnabled interface{} `yaml:"BadgeEnabled,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	EncryptionKey interface{} `yaml:"EncryptionKey,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	ServiceRole interface{} `yaml:"ServiceRole"`
	TimeoutInMinutes interface{} `yaml:"TimeoutInMinutes,omitempty"`
	VpcConfig *properties.Project_VpcConfig `yaml:"VpcConfig,omitempty"`
	Source *properties.Project_Source `yaml:"Source"`
	Triggers *properties.Project_ProjectTriggers `yaml:"Triggers,omitempty"`
	Cache *properties.Project_ProjectCache `yaml:"Cache,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	Environment *properties.Project_Environment `yaml:"Environment"`
	Artifacts *properties.Project_Artifacts `yaml:"Artifacts"`
}

func NewCodeBuildProject(properties CodeBuildProjectProperties, deps ...interface{}) CodeBuildProject {
	return CodeBuildProject{
		Type:       "AWS::CodeBuild::Project",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCodeBuildProject(name string, data string) (cf types.ValueMap, err error) {
	var resource CodeBuildProject
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CodeBuildProject - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CodeBuildProject) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CodeBuildProjectProperties) Validate() []error {
	errs := []error{}
	if resource.ServiceRole == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ServiceRole'"))
	}
	if resource.Source == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Source'"))
	} else {
		errs = append(errs, resource.Source.Validate()...)
	}
	if resource.Environment == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Environment'"))
	} else {
		errs = append(errs, resource.Environment.Validate()...)
	}
	if resource.Artifacts == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Artifacts'"))
	} else {
		errs = append(errs, resource.Artifacts.Validate()...)
	}
	return errs
}
