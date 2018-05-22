package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ElastiCacheSecurityGroupIngress struct {
	Type       string                      `yaml:"Type"`
	Properties ElastiCacheSecurityGroupIngressProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ElastiCacheSecurityGroupIngressProperties struct {
	CacheSecurityGroupName interface{} `yaml:"CacheSecurityGroupName"`
	EC2SecurityGroupName interface{} `yaml:"EC2SecurityGroupName"`
	EC2SecurityGroupOwnerId interface{} `yaml:"EC2SecurityGroupOwnerId,omitempty"`
}

func NewElastiCacheSecurityGroupIngress(properties ElastiCacheSecurityGroupIngressProperties, deps ...interface{}) ElastiCacheSecurityGroupIngress {
	return ElastiCacheSecurityGroupIngress{
		Type:       "AWS::ElastiCache::SecurityGroupIngress",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElastiCacheSecurityGroupIngress(name string, data string) (cf types.ValueMap, err error) {
	var resource ElastiCacheSecurityGroupIngress
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElastiCacheSecurityGroupIngress - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ElastiCacheSecurityGroupIngress) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElastiCacheSecurityGroupIngressProperties) Validate() []error {
	errs := []error{}
	if resource.CacheSecurityGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CacheSecurityGroupName'"))
	}
	if resource.EC2SecurityGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'EC2SecurityGroupName'"))
	}
	return errs
}
