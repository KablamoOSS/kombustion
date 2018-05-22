package properties

	import "fmt"

type DataSource_ElasticsearchConfig struct {
	
	
	AwsRegion interface{} `yaml:"AwsRegion"`
	Endpoint interface{} `yaml:"Endpoint"`
}

func (resource DataSource_ElasticsearchConfig) Validate() []error {
	errs := []error{}
	
	
	if resource.AwsRegion == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AwsRegion'"))
	}
	if resource.Endpoint == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Endpoint'"))
	}
	return errs
}
