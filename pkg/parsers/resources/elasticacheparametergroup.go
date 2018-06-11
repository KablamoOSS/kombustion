package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ElastiCacheParameterGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-parameter-group.html
type ElastiCacheParameterGroup struct {
	Type       string                              `yaml:"Type"`
	Properties ElastiCacheParameterGroupProperties `yaml:"Properties"`
	Condition  interface{}                         `yaml:"Condition,omitempty"`
	Metadata   interface{}                         `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                         `yaml:"DependsOn,omitempty"`
}

// ElastiCacheParameterGroup Properties
type ElastiCacheParameterGroupProperties struct {
	CacheParameterGroupFamily interface{} `yaml:"CacheParameterGroupFamily"`
	Description               interface{} `yaml:"Description"`
	Properties                interface{} `yaml:"Properties,omitempty"`
}

// NewElastiCacheParameterGroup constructor creates a new ElastiCacheParameterGroup
func NewElastiCacheParameterGroup(properties ElastiCacheParameterGroupProperties, deps ...interface{}) ElastiCacheParameterGroup {
	return ElastiCacheParameterGroup{
		Type:       "AWS::ElastiCache::ParameterGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseElastiCacheParameterGroup parses ElastiCacheParameterGroup
func ParseElastiCacheParameterGroup(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
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
	cf = types.TemplateObject{name: resource}
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
