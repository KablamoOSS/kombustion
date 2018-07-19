package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BucketCorsRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html
type BucketCorsRule struct {
	Id             interface{} `yaml:"Id,omitempty"`
	MaxAge         interface{} `yaml:"MaxAge,omitempty"`
	AllowedHeaders interface{} `yaml:"AllowedHeaders,omitempty"`
	AllowedMethods interface{} `yaml:"AllowedMethods"`
	AllowedOrigins interface{} `yaml:"AllowedOrigins"`
	ExposedHeaders interface{} `yaml:"ExposedHeaders,omitempty"`
}

// BucketCorsRule validation
func (resource BucketCorsRule) Validate() []error {
	errors := []error{}

	if resource.AllowedMethods == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'AllowedMethods'"))
	}
	if resource.AllowedOrigins == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'AllowedOrigins'"))
	}
	return errors
}
