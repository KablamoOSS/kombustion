package properties

	import "fmt"

type Bucket_ReplicationConfiguration struct {
	
	
	Role interface{} `yaml:"Role"`
	Rules interface{} `yaml:"Rules"`
}

func (resource Bucket_ReplicationConfiguration) Validate() []error {
	errs := []error{}
	
	
	if resource.Role == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Role'"))
	}
	if resource.Rules == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Rules'"))
	}
	return errs
}
