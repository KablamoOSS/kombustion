package properties

	import "fmt"

type WebACL_Rule struct {
	
	
	
	Priority interface{} `yaml:"Priority"`
	RuleId interface{} `yaml:"RuleId"`
	Action *WebACL_Action `yaml:"Action"`
}

func (resource WebACL_Rule) Validate() []error {
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
