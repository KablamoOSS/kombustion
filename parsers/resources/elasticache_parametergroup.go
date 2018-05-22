package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ElastiCacheParameterGroup struct {
	Type       string                      `yaml:"Type"`
	Properties ElastiCacheParameterGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ElastiCacheParameterGroupProperties struct {
	CacheParameterGroupFamily interface{} `yaml:"CacheParameterGroupFamily"`
	Description interface{} `yaml:"Description"`
	Properties interface{} `yaml:"Properties,omitempty"`
}

func NewElastiCacheParameterGroup(properties ElastiCacheParameterGroupProperties, deps ...interface{}) ElastiCacheParameterGroup {
	return ElastiCacheParameterGroup{
		Type:       "AWS::ElastiCache::ParameterGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElastiCacheParameterGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource ElastiCacheParameterGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElastiCacheParameterGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ElastiCacheParameterGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElastiCacheParameterGroupProperties) Validate() []error {
	errs := []error{}
	if resource.CacheParameterGroupFamily == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CacheParameterGroupFamily'"))
	}
	if resource.Description == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Description'"))
	}
	return errs
}
