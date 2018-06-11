package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DistributionDistributionConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html
type DistributionDistributionConfig struct {
	Comment              interface{}                       `yaml:"Comment,omitempty"`
	DefaultRootObject    interface{}                       `yaml:"DefaultRootObject,omitempty"`
	Enabled              interface{}                       `yaml:"Enabled"`
	HttpVersion          interface{}                       `yaml:"HttpVersion,omitempty"`
	IPV6Enabled          interface{}                       `yaml:"IPV6Enabled,omitempty"`
	PriceClass           interface{}                       `yaml:"PriceClass,omitempty"`
	WebACLId             interface{}                       `yaml:"WebACLId,omitempty"`
	ViewerCertificate    *DistributionViewerCertificate    `yaml:"ViewerCertificate,omitempty"`
	Restrictions         *DistributionRestrictions         `yaml:"Restrictions,omitempty"`
	Logging              *DistributionLogging              `yaml:"Logging,omitempty"`
	CustomErrorResponses interface{}                       `yaml:"CustomErrorResponses,omitempty"`
	Origins              interface{}                       `yaml:"Origins,omitempty"`
	Aliases              interface{}                       `yaml:"Aliases,omitempty"`
	CacheBehaviors       interface{}                       `yaml:"CacheBehaviors,omitempty"`
	DefaultCacheBehavior *DistributionDefaultCacheBehavior `yaml:"DefaultCacheBehavior,omitempty"`
}

func (resource DistributionDistributionConfig) Validate() []error {
	errs := []error{}

	if resource.Enabled == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Enabled'"))
	}
	return errs
}
