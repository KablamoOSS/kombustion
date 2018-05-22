package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type ElasticLoadBalancingLoadBalancer struct {
	Type       string                      `yaml:"Type"`
	Properties ElasticLoadBalancingLoadBalancerProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ElasticLoadBalancingLoadBalancerProperties struct {
	CrossZone interface{} `yaml:"CrossZone,omitempty"`
	LoadBalancerName interface{} `yaml:"LoadBalancerName,omitempty"`
	Scheme interface{} `yaml:"Scheme,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	AppCookieStickinessPolicy interface{} `yaml:"AppCookieStickinessPolicy,omitempty"`
	AvailabilityZones interface{} `yaml:"AvailabilityZones,omitempty"`
	Subnets interface{} `yaml:"Subnets,omitempty"`
	SecurityGroups interface{} `yaml:"SecurityGroups,omitempty"`
	Policies interface{} `yaml:"Policies,omitempty"`
	Instances interface{} `yaml:"Instances,omitempty"`
	LBCookieStickinessPolicy interface{} `yaml:"LBCookieStickinessPolicy,omitempty"`
	Listeners interface{} `yaml:"Listeners"`
	HealthCheck *properties.LoadBalancer_HealthCheck `yaml:"HealthCheck,omitempty"`
	ConnectionSettings *properties.LoadBalancer_ConnectionSettings `yaml:"ConnectionSettings,omitempty"`
	ConnectionDrainingPolicy *properties.LoadBalancer_ConnectionDrainingPolicy `yaml:"ConnectionDrainingPolicy,omitempty"`
	AccessLoggingPolicy *properties.LoadBalancer_AccessLoggingPolicy `yaml:"AccessLoggingPolicy,omitempty"`
}

func NewElasticLoadBalancingLoadBalancer(properties ElasticLoadBalancingLoadBalancerProperties, deps ...interface{}) ElasticLoadBalancingLoadBalancer {
	return ElasticLoadBalancingLoadBalancer{
		Type:       "AWS::ElasticLoadBalancing::LoadBalancer",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElasticLoadBalancingLoadBalancer(name string, data string) (cf types.ValueMap, err error) {
	var resource ElasticLoadBalancingLoadBalancer
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticLoadBalancingLoadBalancer - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ElasticLoadBalancingLoadBalancer) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElasticLoadBalancingLoadBalancerProperties) Validate() []error {
	errs := []error{}
	if resource.Listeners == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Listeners'"))
	}
	return errs
}
