package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type CloudFrontCloudFrontOriginAccessIdentity struct {
	Type       string                      `yaml:"Type"`
	Properties CloudFrontCloudFrontOriginAccessIdentityProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CloudFrontCloudFrontOriginAccessIdentityProperties struct {
	CloudFrontOriginAccessIdentityConfig *properties.CloudFrontOriginAccessIdentity_CloudFrontOriginAccessIdentityConfig `yaml:"CloudFrontOriginAccessIdentityConfig"`
}

func NewCloudFrontCloudFrontOriginAccessIdentity(properties CloudFrontCloudFrontOriginAccessIdentityProperties, deps ...interface{}) CloudFrontCloudFrontOriginAccessIdentity {
	return CloudFrontCloudFrontOriginAccessIdentity{
		Type:       "AWS::CloudFront::CloudFrontOriginAccessIdentity",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCloudFrontCloudFrontOriginAccessIdentity(name string, data string) (cf types.ValueMap, err error) {
	var resource CloudFrontCloudFrontOriginAccessIdentity
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CloudFrontCloudFrontOriginAccessIdentity - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CloudFrontCloudFrontOriginAccessIdentity) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CloudFrontCloudFrontOriginAccessIdentityProperties) Validate() []error {
	errs := []error{}
	if resource.CloudFrontOriginAccessIdentityConfig == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CloudFrontOriginAccessIdentityConfig'"))
	} else {
		errs = append(errs, resource.CloudFrontOriginAccessIdentityConfig.Validate()...)
	}
	return errs
}
