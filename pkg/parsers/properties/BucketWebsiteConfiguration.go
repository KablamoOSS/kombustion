package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// BucketWebsiteConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html
type BucketWebsiteConfiguration struct {
	ErrorDocument         interface{}                  `yaml:"ErrorDocument,omitempty"`
	IndexDocument         interface{}                  `yaml:"IndexDocument,omitempty"`
	RedirectAllRequestsTo *BucketRedirectAllRequestsTo `yaml:"RedirectAllRequestsTo,omitempty"`
	RoutingRules          interface{}                  `yaml:"RoutingRules,omitempty"`
}

func (resource BucketWebsiteConfiguration) Validate() []error {
	errs := []error{}

	return errs
}
