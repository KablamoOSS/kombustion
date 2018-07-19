package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DistributionCookies Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-cookies.html
type DistributionCookies struct {
	Forward          interface{} `yaml:"Forward"`
	WhitelistedNames interface{} `yaml:"WhitelistedNames,omitempty"`
}

// DistributionCookies validation
func (resource DistributionCookies) Validate() []error {
	errors := []error{}

	if resource.Forward == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Forward'"))
	}
	return errors
}
