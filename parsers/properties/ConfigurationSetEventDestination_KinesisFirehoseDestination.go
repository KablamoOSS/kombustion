package properties

	import "fmt"

type ConfigurationSetEventDestination_KinesisFirehoseDestination struct {
	
	
	DeliveryStreamARN interface{} `yaml:"DeliveryStreamARN"`
	IAMRoleARN interface{} `yaml:"IAMRoleARN"`
}

func (resource ConfigurationSetEventDestination_KinesisFirehoseDestination) Validate() []error {
	errs := []error{}
	
	
	if resource.DeliveryStreamARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DeliveryStreamARN'"))
	}
	if resource.IAMRoleARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IAMRoleARN'"))
	}
	return errs
}
