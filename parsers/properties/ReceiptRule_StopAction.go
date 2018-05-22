package properties

	import "fmt"

type ReceiptRule_StopAction struct {
	
	
	Scope interface{} `yaml:"Scope"`
	TopicArn interface{} `yaml:"TopicArn,omitempty"`
}

func (resource ReceiptRule_StopAction) Validate() []error {
	errs := []error{}
	
	
	if resource.Scope == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Scope'"))
	}
	return errs
}
