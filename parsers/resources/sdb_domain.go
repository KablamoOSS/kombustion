package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type SDBDomain struct {
	Type       string                      `yaml:"Type"`
	Properties SDBDomainProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type SDBDomainProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
}

func NewSDBDomain(properties SDBDomainProperties, deps ...interface{}) SDBDomain {
	return SDBDomain{
		Type:       "AWS::SDB::Domain",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSDBDomain(name string, data string) (cf types.ValueMap, err error) {
	var resource SDBDomain
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SDBDomain - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource SDBDomain) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SDBDomainProperties) Validate() []error {
	errs := []error{}
	return errs
}
