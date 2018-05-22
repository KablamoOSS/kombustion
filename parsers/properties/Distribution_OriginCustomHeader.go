package properties

	import "fmt"

type Distribution_OriginCustomHeader struct {
	
	
	HeaderName interface{} `yaml:"HeaderName"`
	HeaderValue interface{} `yaml:"HeaderValue"`
}

func (resource Distribution_OriginCustomHeader) Validate() []error {
	errs := []error{}
	
	
	if resource.HeaderName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HeaderName'"))
	}
	if resource.HeaderValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HeaderValue'"))
	}
	return errs
}
