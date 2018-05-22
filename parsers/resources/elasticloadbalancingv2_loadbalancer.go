package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type ElasticLoadBalancingV2LoadBalancer struct {
	Type       string                      `yaml:"Type"`
	Properties ElasticLoadBalancingV2LoadBalancerProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ElasticLoadBalancingV2LoadBalancerProperties struct {
	IpAddressType interface{} `yaml:"IpAddressType,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	Scheme interface{} `yaml:"Scheme,omitempty"`
	Type interface{} `yaml:"Type,omitempty"`
	LoadBalancerAttributes interface{} `yaml:"LoadBalancerAttributes,omitempty"`
	SecurityGroups interface{} `yaml:"SecurityGroups,omitempty"`
	SubnetMappings interface{} `yaml:"SubnetMappings,omitempty"`
	Subnets interface{} `yaml:"Subnets,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewElasticLoadBalancingV2LoadBalancer(properties ElasticLoadBalancingV2LoadBalancerProperties, deps ...interface{}) ElasticLoadBalancingV2LoadBalancer {
	return ElasticLoadBalancingV2LoadBalancer{
		Type:       "AWS::ElasticLoadBalancingV2::LoadBalancer",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElasticLoadBalancingV2LoadBalancer(name string, data string) (cf types.ValueMap, err error) {
	var resource ElasticLoadBalancingV2LoadBalancer
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticLoadBalancingV2LoadBalancer - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ElasticLoadBalancingV2LoadBalancer) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElasticLoadBalancingV2LoadBalancerProperties) Validate() []error {
	errs := []error{}
	return errs
}
