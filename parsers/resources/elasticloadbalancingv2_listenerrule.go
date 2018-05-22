package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ElasticLoadBalancingV2ListenerRule struct {
	Type       string                      `yaml:"Type"`
	Properties ElasticLoadBalancingV2ListenerRuleProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ElasticLoadBalancingV2ListenerRuleProperties struct {
	ListenerArn interface{} `yaml:"ListenerArn"`
	Priority interface{} `yaml:"Priority"`
	Actions interface{} `yaml:"Actions"`
	Conditions interface{} `yaml:"Conditions"`
}

func NewElasticLoadBalancingV2ListenerRule(properties ElasticLoadBalancingV2ListenerRuleProperties, deps ...interface{}) ElasticLoadBalancingV2ListenerRule {
	return ElasticLoadBalancingV2ListenerRule{
		Type:       "AWS::ElasticLoadBalancingV2::ListenerRule",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElasticLoadBalancingV2ListenerRule(name string, data string) (cf types.ValueMap, err error) {
	var resource ElasticLoadBalancingV2ListenerRule
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticLoadBalancingV2ListenerRule - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ElasticLoadBalancingV2ListenerRule) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElasticLoadBalancingV2ListenerRuleProperties) Validate() []error {
	errs := []error{}
	if resource.ListenerArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ListenerArn'"))
	}
	if resource.Priority == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Priority'"))
	}
	if resource.Actions == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Actions'"))
	}
	if resource.Conditions == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Conditions'"))
	}
	return errs
}
