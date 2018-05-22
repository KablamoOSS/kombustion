package properties

	import "fmt"

type ReceiptRule_WorkmailAction struct {
	
	
	OrganizationArn interface{} `yaml:"OrganizationArn"`
	TopicArn interface{} `yaml:"TopicArn,omitempty"`
}

func (resource ReceiptRule_WorkmailAction) Validate() []error {
	errs := []error{}
	
	
	if resource.OrganizationArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'OrganizationArn'"))
	}
	return errs
}
