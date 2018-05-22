package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type CodeCommitRepository struct {
	Type       string                      `yaml:"Type"`
	Properties CodeCommitRepositoryProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CodeCommitRepositoryProperties struct {
	RepositoryDescription interface{} `yaml:"RepositoryDescription,omitempty"`
	RepositoryName interface{} `yaml:"RepositoryName"`
	Triggers interface{} `yaml:"Triggers,omitempty"`
}

func NewCodeCommitRepository(properties CodeCommitRepositoryProperties, deps ...interface{}) CodeCommitRepository {
	return CodeCommitRepository{
		Type:       "AWS::CodeCommit::Repository",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCodeCommitRepository(name string, data string) (cf types.ValueMap, err error) {
	var resource CodeCommitRepository
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CodeCommitRepository - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CodeCommitRepository) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CodeCommitRepositoryProperties) Validate() []error {
	errs := []error{}
	if resource.RepositoryName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RepositoryName'"))
	}
	return errs
}
