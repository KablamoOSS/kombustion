package properties

	import "fmt"

type Topic_Subscription struct {
	
	
	Endpoint interface{} `yaml:"Endpoint"`
	Protocol interface{} `yaml:"Protocol"`
}

func (resource Topic_Subscription) Validate() []error {
	errs := []error{}
	
	
	if resource.Endpoint == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Endpoint'"))
	}
	if resource.Protocol == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Protocol'"))
	}
	return errs
}
