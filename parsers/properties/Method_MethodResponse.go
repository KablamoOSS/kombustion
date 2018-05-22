package properties

	import "fmt"

type Method_MethodResponse struct {
	
	
	
	StatusCode interface{} `yaml:"StatusCode"`
	ResponseModels interface{} `yaml:"ResponseModels,omitempty"`
	ResponseParameters interface{} `yaml:"ResponseParameters,omitempty"`
}

func (resource Method_MethodResponse) Validate() []error {
	errs := []error{}
	
	
	
	if resource.StatusCode == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StatusCode'"))
	}
	return errs
}
