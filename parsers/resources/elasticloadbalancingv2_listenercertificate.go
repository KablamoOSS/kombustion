package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type ElasticLoadBalancingV2ListenerCertificate struct {
	Type       string                      `yaml:"Type"`
	Properties ElasticLoadBalancingV2ListenerCertificateProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ElasticLoadBalancingV2ListenerCertificateProperties struct {
	ListenerArn interface{} `yaml:"ListenerArn"`
	Certificates interface{} `yaml:"Certificates"`
}

func NewElasticLoadBalancingV2ListenerCertificate(properties ElasticLoadBalancingV2ListenerCertificateProperties, deps ...interface{}) ElasticLoadBalancingV2ListenerCertificate {
	return ElasticLoadBalancingV2ListenerCertificate{
		Type:       "AWS::ElasticLoadBalancingV2::ListenerCertificate",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElasticLoadBalancingV2ListenerCertificate(name string, data string) (cf types.ValueMap, err error) {
	var resource ElasticLoadBalancingV2ListenerCertificate
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticLoadBalancingV2ListenerCertificate - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ElasticLoadBalancingV2ListenerCertificate) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElasticLoadBalancingV2ListenerCertificateProperties) Validate() []error {
	errs := []error{}
	if resource.ListenerArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ListenerArn'"))
	}
	if resource.Certificates == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Certificates'"))
	}
	return errs
}
