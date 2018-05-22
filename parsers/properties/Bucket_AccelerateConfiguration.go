package properties

	import "fmt"

type Bucket_AccelerateConfiguration struct {
	
	AccelerationStatus interface{} `yaml:"AccelerationStatus"`
}

func (resource Bucket_AccelerateConfiguration) Validate() []error {
	errs := []error{}
	
	if resource.AccelerationStatus == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AccelerationStatus'"))
	}
	return errs
}
