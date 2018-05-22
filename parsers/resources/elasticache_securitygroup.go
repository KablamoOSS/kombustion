package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ElastiCacheSecurityGroup struct {
	Type       string                      `yaml:"Type"`
	Properties ElastiCacheSecurityGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ElastiCacheSecurityGroupProperties struct {
	Description interface{} `yaml:"Description"`
}

func NewElastiCacheSecurityGroup(properties ElastiCacheSecurityGroupProperties, deps ...interface{}) ElastiCacheSecurityGroup {
	return ElastiCacheSecurityGroup{
		Type:       "AWS::ElastiCache::SecurityGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElastiCacheSecurityGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource ElastiCacheSecurityGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElastiCacheSecurityGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ElastiCacheSecurityGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElastiCacheSecurityGroupProperties) Validate() []error {
	errs := []error{}
	if resource.Description == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Description'"))
	}
	return errs
}
