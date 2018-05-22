package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type OpsWorksElasticLoadBalancerAttachment struct {
	Type       string                      `yaml:"Type"`
	Properties OpsWorksElasticLoadBalancerAttachmentProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type OpsWorksElasticLoadBalancerAttachmentProperties struct {
	ElasticLoadBalancerName interface{} `yaml:"ElasticLoadBalancerName"`
	LayerId interface{} `yaml:"LayerId"`
}

func NewOpsWorksElasticLoadBalancerAttachment(properties OpsWorksElasticLoadBalancerAttachmentProperties, deps ...interface{}) OpsWorksElasticLoadBalancerAttachment {
	return OpsWorksElasticLoadBalancerAttachment{
		Type:       "AWS::OpsWorks::ElasticLoadBalancerAttachment",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseOpsWorksElasticLoadBalancerAttachment(name string, data string) (cf types.ValueMap, err error) {
	var resource OpsWorksElasticLoadBalancerAttachment
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: OpsWorksElasticLoadBalancerAttachment - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource OpsWorksElasticLoadBalancerAttachment) Validate() []error {
	return resource.Properties.Validate()
}

func (resource OpsWorksElasticLoadBalancerAttachmentProperties) Validate() []error {
	errs := []error{}
	if resource.ElasticLoadBalancerName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ElasticLoadBalancerName'"))
	}
	if resource.LayerId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'LayerId'"))
	}
	return errs
}
