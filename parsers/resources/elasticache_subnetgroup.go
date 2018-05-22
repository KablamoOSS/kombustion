package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ElastiCacheSubnetGroup struct {
	Type       string                      `yaml:"Type"`
	Properties ElastiCacheSubnetGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ElastiCacheSubnetGroupProperties struct {
	CacheSubnetGroupName interface{} `yaml:"CacheSubnetGroupName,omitempty"`
	Description interface{} `yaml:"Description"`
	SubnetIds interface{} `yaml:"SubnetIds"`
}

func NewElastiCacheSubnetGroup(properties ElastiCacheSubnetGroupProperties, deps ...interface{}) ElastiCacheSubnetGroup {
	return ElastiCacheSubnetGroup{
		Type:       "AWS::ElastiCache::SubnetGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElastiCacheSubnetGroup(name string, data string) (cf types.ValueMap, err error) {
	var resource ElastiCacheSubnetGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElastiCacheSubnetGroup - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ElastiCacheSubnetGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElastiCacheSubnetGroupProperties) Validate() []error {
	errs := []error{}
	if resource.Description == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Description'"))
	}
	if resource.SubnetIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetIds'"))
	}
	return errs
}
