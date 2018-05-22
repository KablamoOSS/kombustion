package properties

	import "fmt"

type Bucket_RedirectAllRequestsTo struct {
	
	
	HostName interface{} `yaml:"HostName"`
	Protocol interface{} `yaml:"Protocol,omitempty"`
}

func (resource Bucket_RedirectAllRequestsTo) Validate() []error {
	errs := []error{}
	
	
	if resource.HostName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HostName'"))
	}
	return errs
}
