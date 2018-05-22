package properties

	import "fmt"

type LoadBalancer_AppCookieStickinessPolicy struct {
	
	
	CookieName interface{} `yaml:"CookieName"`
	PolicyName interface{} `yaml:"PolicyName"`
}

func (resource LoadBalancer_AppCookieStickinessPolicy) Validate() []error {
	errs := []error{}
	
	
	if resource.CookieName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CookieName'"))
	}
	if resource.PolicyName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyName'"))
	}
	return errs
}
