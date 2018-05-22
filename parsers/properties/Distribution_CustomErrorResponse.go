package properties

	import "fmt"

type Distribution_CustomErrorResponse struct {
	
	
	
	
	ErrorCachingMinTTL interface{} `yaml:"ErrorCachingMinTTL,omitempty"`
	ErrorCode interface{} `yaml:"ErrorCode"`
	ResponseCode interface{} `yaml:"ResponseCode,omitempty"`
	ResponsePagePath interface{} `yaml:"ResponsePagePath,omitempty"`
}

func (resource Distribution_CustomErrorResponse) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.ErrorCode == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ErrorCode'"))
	}
	return errs
}
