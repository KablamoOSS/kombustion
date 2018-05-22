package properties

	import "fmt"

type Distribution_CacheBehavior struct {
	
	
	
	
	
	
	
	
	
	
	
	
	
	Compress interface{} `yaml:"Compress,omitempty"`
	DefaultTTL interface{} `yaml:"DefaultTTL,omitempty"`
	MaxTTL interface{} `yaml:"MaxTTL,omitempty"`
	MinTTL interface{} `yaml:"MinTTL,omitempty"`
	PathPattern interface{} `yaml:"PathPattern"`
	SmoothStreaming interface{} `yaml:"SmoothStreaming,omitempty"`
	TargetOriginId interface{} `yaml:"TargetOriginId"`
	ViewerProtocolPolicy interface{} `yaml:"ViewerProtocolPolicy"`
	AllowedMethods interface{} `yaml:"AllowedMethods,omitempty"`
	CachedMethods interface{} `yaml:"CachedMethods,omitempty"`
	LambdaFunctionAssociations interface{} `yaml:"LambdaFunctionAssociations,omitempty"`
	TrustedSigners interface{} `yaml:"TrustedSigners,omitempty"`
	ForwardedValues *Distribution_ForwardedValues `yaml:"ForwardedValues"`
}

func (resource Distribution_CacheBehavior) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	
	
	
	if resource.PathPattern == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PathPattern'"))
	}
	if resource.TargetOriginId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetOriginId'"))
	}
	if resource.ViewerProtocolPolicy == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ViewerProtocolPolicy'"))
	}
	if resource.ForwardedValues == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ForwardedValues'"))
	} else {
		errs = append(errs, resource.ForwardedValues.Validate()...)
	}
	return errs
}
