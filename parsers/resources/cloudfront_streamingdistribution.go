package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type CloudFrontStreamingDistribution struct {
	Type       string                      `yaml:"Type"`
	Properties CloudFrontStreamingDistributionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CloudFrontStreamingDistributionProperties struct {
	StreamingDistributionConfig *properties.StreamingDistribution_StreamingDistributionConfig `yaml:"StreamingDistributionConfig"`
	Tags interface{} `yaml:"Tags"`
}

func NewCloudFrontStreamingDistribution(properties CloudFrontStreamingDistributionProperties, deps ...interface{}) CloudFrontStreamingDistribution {
	return CloudFrontStreamingDistribution{
		Type:       "AWS::CloudFront::StreamingDistribution",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCloudFrontStreamingDistribution(name string, data string) (cf types.ValueMap, err error) {
	var resource CloudFrontStreamingDistribution
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CloudFrontStreamingDistribution - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CloudFrontStreamingDistribution) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CloudFrontStreamingDistributionProperties) Validate() []error {
	errs := []error{}
	if resource.StreamingDistributionConfig == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StreamingDistributionConfig'"))
	} else {
		errs = append(errs, resource.StreamingDistributionConfig.Validate()...)
	}
	if resource.Tags == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Tags'"))
	}
	return errs
}
