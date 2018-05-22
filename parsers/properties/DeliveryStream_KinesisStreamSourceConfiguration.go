package properties

	import "fmt"

type DeliveryStream_KinesisStreamSourceConfiguration struct {
	
	
	KinesisStreamARN interface{} `yaml:"KinesisStreamARN"`
	RoleARN interface{} `yaml:"RoleARN"`
}

func (resource DeliveryStream_KinesisStreamSourceConfiguration) Validate() []error {
	errs := []error{}
	
	
	if resource.KinesisStreamARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'KinesisStreamARN'"))
	}
	if resource.RoleARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleARN'"))
	}
	return errs
}
