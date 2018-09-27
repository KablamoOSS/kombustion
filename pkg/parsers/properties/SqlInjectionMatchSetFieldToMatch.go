package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// SqlInjectionMatchSetFieldToMatch Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html
type SqlInjectionMatchSetFieldToMatch struct {
	Data interface{} `yaml:"Data,omitempty"`
	Type interface{} `yaml:"Type"`
}

// SqlInjectionMatchSetFieldToMatch validation
func (resource SqlInjectionMatchSetFieldToMatch) Validate() []error {
	errors := []error{}

	if resource.Type == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Type'"))
	}
	return errors
}
