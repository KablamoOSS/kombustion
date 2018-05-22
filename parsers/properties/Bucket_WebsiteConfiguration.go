package properties


type Bucket_WebsiteConfiguration struct {
	
	
	
	
	ErrorDocument interface{} `yaml:"ErrorDocument,omitempty"`
	IndexDocument interface{} `yaml:"IndexDocument,omitempty"`
	RedirectAllRequestsTo *Bucket_RedirectAllRequestsTo `yaml:"RedirectAllRequestsTo,omitempty"`
	RoutingRules interface{} `yaml:"RoutingRules,omitempty"`
}

func (resource Bucket_WebsiteConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	
	return errs
}
