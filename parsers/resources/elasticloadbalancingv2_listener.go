package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ElasticLoadBalancingV2Listener struct {
	Type       string                      `yaml:"Type"`
	Properties ElasticLoadBalancingV2ListenerProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ElasticLoadBalancingV2ListenerProperties struct {
	LoadBalancerArn interface{} `yaml:"LoadBalancerArn"`
	Port interface{} `yaml:"Port"`
	Protocol interface{} `yaml:"Protocol"`
	SslPolicy interface{} `yaml:"SslPolicy,omitempty"`
	Certificates interface{} `yaml:"Certificates,omitempty"`
	DefaultActions interface{} `yaml:"DefaultActions"`
}

func NewElasticLoadBalancingV2Listener(properties ElasticLoadBalancingV2ListenerProperties, deps ...interface{}) ElasticLoadBalancingV2Listener {
	return ElasticLoadBalancingV2Listener{
		Type:       "AWS::ElasticLoadBalancingV2::Listener",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElasticLoadBalancingV2Listener(name string, data string) (cf types.ValueMap, err error) {
	var resource ElasticLoadBalancingV2Listener
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticLoadBalancingV2Listener - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ElasticLoadBalancingV2Listener) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElasticLoadBalancingV2ListenerProperties) Validate() []error {
	errs := []error{}
	if resource.LoadBalancerArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'LoadBalancerArn'"))
	}
	if resource.Port == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Port'"))
	}
	if resource.Protocol == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Protocol'"))
	}
	if resource.DefaultActions == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DefaultActions'"))
	}
	return errs
}
