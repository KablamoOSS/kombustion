package properties

	import "fmt"

type StreamingDistribution_StreamingDistributionConfig struct {
	
	
	
	
	
	
	
	Comment interface{} `yaml:"Comment"`
	Enabled interface{} `yaml:"Enabled"`
	PriceClass interface{} `yaml:"PriceClass,omitempty"`
	TrustedSigners *StreamingDistribution_TrustedSigners `yaml:"TrustedSigners"`
	S3Origin *StreamingDistribution_S3Origin `yaml:"S3Origin"`
	Logging *StreamingDistribution_Logging `yaml:"Logging,omitempty"`
	Aliases interface{} `yaml:"Aliases,omitempty"`
}

func (resource StreamingDistribution_StreamingDistributionConfig) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	if resource.Comment == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Comment'"))
	}
	if resource.Enabled == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Enabled'"))
	}
	if resource.TrustedSigners == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TrustedSigners'"))
	} else {
		errs = append(errs, resource.TrustedSigners.Validate()...)
	}
	if resource.S3Origin == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'S3Origin'"))
	} else {
		errs = append(errs, resource.S3Origin.Validate()...)
	}
	return errs
}
