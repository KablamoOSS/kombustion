package properties

	import "fmt"

type Method_IntegrationResponse struct {
	
	
	
	
	
	ContentHandling interface{} `yaml:"ContentHandling,omitempty"`
	SelectionPattern interface{} `yaml:"SelectionPattern,omitempty"`
	StatusCode interface{} `yaml:"StatusCode"`
	ResponseParameters interface{} `yaml:"ResponseParameters,omitempty"`
	ResponseTemplates interface{} `yaml:"ResponseTemplates,omitempty"`
}

func (resource Method_IntegrationResponse) Validate() []error {
	errs := []error{}
	
	
	
	
	
	if resource.StatusCode == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StatusCode'"))
	}
	return errs
}
