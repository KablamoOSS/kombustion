package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type CloudFrontDistribution struct {
	Type       string                      `yaml:"Type"`
	Properties CloudFrontDistributionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CloudFrontDistributionProperties struct {
	Tags interface{} `yaml:"Tags,omitempty"`
	DistributionConfig *properties.Distribution_DistributionConfig `yaml:"DistributionConfig"`
}

func NewCloudFrontDistribution(properties CloudFrontDistributionProperties, deps ...interface{}) CloudFrontDistribution {
	return CloudFrontDistribution{
		Type:       "AWS::CloudFront::Distribution",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCloudFrontDistribution(name string, data string) (cf types.ValueMap, err error) {
	var resource CloudFrontDistribution
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CloudFrontDistribution - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CloudFrontDistribution) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CloudFrontDistributionProperties) Validate() []error {
	errs := []error{}
	if resource.DistributionConfig == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DistributionConfig'"))
	} else {
		errs = append(errs, resource.DistributionConfig.Validate()...)
	}
	return errs
}
