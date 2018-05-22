package properties

	import "fmt"

type Certificate_DomainValidationOption struct {
	
	
	DomainName interface{} `yaml:"DomainName"`
	ValidationDomain interface{} `yaml:"ValidationDomain"`
}

func (resource Certificate_DomainValidationOption) Validate() []error {
	errs := []error{}
	
	
	if resource.DomainName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DomainName'"))
	}
	if resource.ValidationDomain == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ValidationDomain'"))
	}
	return errs
}
