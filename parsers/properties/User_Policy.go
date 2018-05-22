package properties

	import "fmt"

type User_Policy struct {
	
	
	PolicyDocument interface{} `yaml:"PolicyDocument"`
	PolicyName interface{} `yaml:"PolicyName"`
}

func (resource User_Policy) Validate() []error {
	errs := []error{}
	
	
	if resource.PolicyDocument == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyDocument'"))
	}
	if resource.PolicyName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyName'"))
	}
	return errs
}
