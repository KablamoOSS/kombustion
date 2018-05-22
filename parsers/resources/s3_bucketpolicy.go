package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type S3BucketPolicy struct {
	Type       string                      `yaml:"Type"`
	Properties S3BucketPolicyProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type S3BucketPolicyProperties struct {
	Bucket interface{} `yaml:"Bucket"`
	PolicyDocument interface{} `yaml:"PolicyDocument"`
}

func NewS3BucketPolicy(properties S3BucketPolicyProperties, deps ...interface{}) S3BucketPolicy {
	return S3BucketPolicy{
		Type:       "AWS::S3::BucketPolicy",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseS3BucketPolicy(name string, data string) (cf types.ValueMap, err error) {
	var resource S3BucketPolicy
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: S3BucketPolicy - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource S3BucketPolicy) Validate() []error {
	return resource.Properties.Validate()
}

func (resource S3BucketPolicyProperties) Validate() []error {
	errs := []error{}
	if resource.Bucket == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Bucket'"))
	}
	if resource.PolicyDocument == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyDocument'"))
	}
	return errs
}
