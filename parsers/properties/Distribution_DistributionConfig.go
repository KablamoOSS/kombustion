package properties

	import "fmt"

type Distribution_DistributionConfig struct {
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	Comment interface{} `yaml:"Comment,omitempty"`
	DefaultRootObject interface{} `yaml:"DefaultRootObject,omitempty"`
	Enabled interface{} `yaml:"Enabled"`
	HttpVersion interface{} `yaml:"HttpVersion,omitempty"`
	IPV6Enabled interface{} `yaml:"IPV6Enabled,omitempty"`
	PriceClass interface{} `yaml:"PriceClass,omitempty"`
	WebACLId interface{} `yaml:"WebACLId,omitempty"`
	ViewerCertificate *Distribution_ViewerCertificate `yaml:"ViewerCertificate,omitempty"`
	Restrictions *Distribution_Restrictions `yaml:"Restrictions,omitempty"`
	Logging *Distribution_Logging `yaml:"Logging,omitempty"`
	CustomErrorResponses interface{} `yaml:"CustomErrorResponses,omitempty"`
	Origins interface{} `yaml:"Origins,omitempty"`
	Aliases interface{} `yaml:"Aliases,omitempty"`
	CacheBehaviors interface{} `yaml:"CacheBehaviors,omitempty"`
	DefaultCacheBehavior *Distribution_DefaultCacheBehavior `yaml:"DefaultCacheBehavior,omitempty"`
}

func (resource Distribution_DistributionConfig) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	if resource.Enabled == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Enabled'"))
	}
	return errs
}
