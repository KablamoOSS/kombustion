package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// CertificateDomainValidationOption Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-certificate-domainvalidationoption.html
type CertificateDomainValidationOption struct {
	DomainName       interface{} `yaml:"DomainName"`
	ValidationDomain interface{} `yaml:"ValidationDomain"`
}

func (resource CertificateDomainValidationOption) Validate() []error {
	errs := []error{}

	if resource.DomainName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DomainName'"))
	}
	if resource.ValidationDomain == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ValidationDomain'"))
	}
	return errs
}
