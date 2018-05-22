package properties

	import "fmt"

type Distribution_ForwardedValues struct {
	
	
	
	
	QueryString interface{} `yaml:"QueryString"`
	Headers interface{} `yaml:"Headers,omitempty"`
	QueryStringCacheKeys interface{} `yaml:"QueryStringCacheKeys,omitempty"`
	Cookies *Distribution_Cookies `yaml:"Cookies,omitempty"`
}

func (resource Distribution_ForwardedValues) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.QueryString == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'QueryString'"))
	}
	return errs
}
