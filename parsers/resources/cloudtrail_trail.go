package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type CloudTrailTrail struct {
	Type       string                      `yaml:"Type"`
	Properties CloudTrailTrailProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CloudTrailTrailProperties struct {
	CloudWatchLogsLogGroupArn interface{} `yaml:"CloudWatchLogsLogGroupArn,omitempty"`
	CloudWatchLogsRoleArn interface{} `yaml:"CloudWatchLogsRoleArn,omitempty"`
	EnableLogFileValidation interface{} `yaml:"EnableLogFileValidation,omitempty"`
	IncludeGlobalServiceEvents interface{} `yaml:"IncludeGlobalServiceEvents,omitempty"`
	IsLogging interface{} `yaml:"IsLogging"`
	IsMultiRegionTrail interface{} `yaml:"IsMultiRegionTrail,omitempty"`
	KMSKeyId interface{} `yaml:"KMSKeyId,omitempty"`
	S3BucketName interface{} `yaml:"S3BucketName"`
	S3KeyPrefix interface{} `yaml:"S3KeyPrefix,omitempty"`
	SnsTopicName interface{} `yaml:"SnsTopicName,omitempty"`
	TrailName interface{} `yaml:"TrailName,omitempty"`
	EventSelectors interface{} `yaml:"EventSelectors,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewCloudTrailTrail(properties CloudTrailTrailProperties, deps ...interface{}) CloudTrailTrail {
	return CloudTrailTrail{
		Type:       "AWS::CloudTrail::Trail",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCloudTrailTrail(name string, data string) (cf types.ValueMap, err error) {
	var resource CloudTrailTrail
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CloudTrailTrail - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CloudTrailTrail) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CloudTrailTrailProperties) Validate() []error {
	errs := []error{}
	if resource.IsLogging == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IsLogging'"))
	}
	if resource.S3BucketName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'S3BucketName'"))
	}
	return errs
}
