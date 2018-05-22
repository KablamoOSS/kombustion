package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type ECRRepository struct {
	Type       string                      `yaml:"Type"`
	Properties ECRRepositoryProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ECRRepositoryProperties struct {
	RepositoryName interface{} `yaml:"RepositoryName,omitempty"`
	RepositoryPolicyText interface{} `yaml:"RepositoryPolicyText,omitempty"`
	LifecyclePolicy *properties.Repository_LifecyclePolicy `yaml:"LifecyclePolicy,omitempty"`
}

func NewECRRepository(properties ECRRepositoryProperties, deps ...interface{}) ECRRepository {
	return ECRRepository{
		Type:       "AWS::ECR::Repository",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseECRRepository(name string, data string) (cf types.ValueMap, err error) {
	var resource ECRRepository
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ECRRepository - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ECRRepository) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ECRRepositoryProperties) Validate() []error {
	errs := []error{}
	return errs
}
