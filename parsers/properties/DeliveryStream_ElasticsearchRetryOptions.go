package properties

	import "fmt"

type DeliveryStream_ElasticsearchRetryOptions struct {
	
	DurationInSeconds interface{} `yaml:"DurationInSeconds"`
}

func (resource DeliveryStream_ElasticsearchRetryOptions) Validate() []error {
	errs := []error{}
	
	if resource.DurationInSeconds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DurationInSeconds'"))
	}
	return errs
}
