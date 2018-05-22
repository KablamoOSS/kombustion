package properties

	import "fmt"

type ReceiptRule_LambdaAction struct {
	
	
	
	FunctionArn interface{} `yaml:"FunctionArn"`
	InvocationType interface{} `yaml:"InvocationType,omitempty"`
	TopicArn interface{} `yaml:"TopicArn,omitempty"`
}

func (resource ReceiptRule_LambdaAction) Validate() []error {
	errs := []error{}
	
	
	
	if resource.FunctionArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FunctionArn'"))
	}
	return errs
}
