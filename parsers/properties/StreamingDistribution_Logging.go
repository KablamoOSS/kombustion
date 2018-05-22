package properties

	import "fmt"

type StreamingDistribution_Logging struct {
	
	
	
	Bucket interface{} `yaml:"Bucket"`
	Enabled interface{} `yaml:"Enabled"`
	Prefix interface{} `yaml:"Prefix"`
}

func (resource StreamingDistribution_Logging) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Bucket == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Bucket'"))
	}
	if resource.Enabled == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Enabled'"))
	}
	if resource.Prefix == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Prefix'"))
	}
	return errs
}
