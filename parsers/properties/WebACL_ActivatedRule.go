package properties

	import "fmt"

type WebACL_ActivatedRule struct {
	
	
	
	Priority interface{} `yaml:"Priority"`
	RuleId interface{} `yaml:"RuleId"`
	Action *WebACL_WafAction `yaml:"Action"`
}

func (resource WebACL_ActivatedRule) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Priority == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Priority'"))
	}
	if resource.RuleId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RuleId'"))
	}
	if resource.Action == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Action'"))
	} else {
		errs = append(errs, resource.Action.Validate()...)
	}
	return errs
}
