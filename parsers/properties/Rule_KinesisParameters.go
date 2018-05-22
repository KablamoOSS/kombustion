package properties

	import "fmt"

type Rule_KinesisParameters struct {
	
	PartitionKeyPath interface{} `yaml:"PartitionKeyPath"`
}

func (resource Rule_KinesisParameters) Validate() []error {
	errs := []error{}
	
	if resource.PartitionKeyPath == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PartitionKeyPath'"))
	}
	return errs
}
