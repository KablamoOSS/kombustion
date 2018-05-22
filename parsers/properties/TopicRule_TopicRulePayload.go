package properties

	import "fmt"

type TopicRule_TopicRulePayload struct {
	
	
	
	
	
	AwsIotSqlVersion interface{} `yaml:"AwsIotSqlVersion,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	RuleDisabled interface{} `yaml:"RuleDisabled"`
	Sql interface{} `yaml:"Sql"`
	Actions interface{} `yaml:"Actions"`
}

func (resource TopicRule_TopicRulePayload) Validate() []error {
	errs := []error{}
	
	
	
	
	
	if resource.RuleDisabled == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RuleDisabled'"))
	}
	if resource.Sql == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Sql'"))
	}
	if resource.Actions == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Actions'"))
	}
	return errs
}
