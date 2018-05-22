package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type S3Bucket struct {
	Type       string                      `yaml:"Type"`
	Properties S3BucketProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type S3BucketProperties struct {
	AccessControl interface{} `yaml:"AccessControl,omitempty"`
	BucketName interface{} `yaml:"BucketName,omitempty"`
	WebsiteConfiguration *properties.Bucket_WebsiteConfiguration `yaml:"WebsiteConfiguration,omitempty"`
	VersioningConfiguration *properties.Bucket_VersioningConfiguration `yaml:"VersioningConfiguration,omitempty"`
	ReplicationConfiguration *properties.Bucket_ReplicationConfiguration `yaml:"ReplicationConfiguration,omitempty"`
	NotificationConfiguration *properties.Bucket_NotificationConfiguration `yaml:"NotificationConfiguration,omitempty"`
	LoggingConfiguration *properties.Bucket_LoggingConfiguration `yaml:"LoggingConfiguration,omitempty"`
	AnalyticsConfigurations interface{} `yaml:"AnalyticsConfigurations,omitempty"`
	InventoryConfigurations interface{} `yaml:"InventoryConfigurations,omitempty"`
	MetricsConfigurations interface{} `yaml:"MetricsConfigurations,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	LifecycleConfiguration *properties.Bucket_LifecycleConfiguration `yaml:"LifecycleConfiguration,omitempty"`
	CorsConfiguration *properties.Bucket_CorsConfiguration `yaml:"CorsConfiguration,omitempty"`
	BucketEncryption *properties.Bucket_BucketEncryption `yaml:"BucketEncryption,omitempty"`
	AccelerateConfiguration *properties.Bucket_AccelerateConfiguration `yaml:"AccelerateConfiguration,omitempty"`
}

func NewS3Bucket(properties S3BucketProperties, deps ...interface{}) S3Bucket {
	return S3Bucket{
		Type:       "AWS::S3::Bucket",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseS3Bucket(name string, data string) (cf types.ValueMap, err error) {
	var resource S3Bucket
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: S3Bucket - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource S3Bucket) Validate() []error {
	return resource.Properties.Validate()
}

func (resource S3BucketProperties) Validate() []error {
	errs := []error{}
	return errs
}
