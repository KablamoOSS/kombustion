package properties

	import "fmt"

type StreamingDistribution_TrustedSigners struct {
	
	
	Enabled interface{} `yaml:"Enabled"`
	AwsAccountNumbers interface{} `yaml:"AwsAccountNumbers,omitempty"`
}

func (resource StreamingDistribution_TrustedSigners) Validate() []error {
	errs := []error{}
	
	
	if resource.Enabled == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Enabled'"))
	}
	return errs
}
