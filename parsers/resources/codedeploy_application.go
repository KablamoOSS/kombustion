package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type CodeDeployApplication struct {
	Type       string                      `yaml:"Type"`
	Properties CodeDeployApplicationProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CodeDeployApplicationProperties struct {
	ApplicationName interface{} `yaml:"ApplicationName,omitempty"`
	ComputePlatform interface{} `yaml:"ComputePlatform,omitempty"`
}

func NewCodeDeployApplication(properties CodeDeployApplicationProperties, deps ...interface{}) CodeDeployApplication {
	return CodeDeployApplication{
		Type:       "AWS::CodeDeploy::Application",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCodeDeployApplication(name string, data string) (cf types.ValueMap, err error) {
	var resource CodeDeployApplication
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CodeDeployApplication - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CodeDeployApplication) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CodeDeployApplicationProperties) Validate() []error {
	errs := []error{}
	return errs
}
