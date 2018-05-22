package properties

	import "fmt"

type DataSource_DynamoDBConfig struct {
	
	
	
	AwsRegion interface{} `yaml:"AwsRegion"`
	TableName interface{} `yaml:"TableName"`
	UseCallerCredentials interface{} `yaml:"UseCallerCredentials,omitempty"`
}

func (resource DataSource_DynamoDBConfig) Validate() []error {
	errs := []error{}
	
	
	
	if resource.AwsRegion == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AwsRegion'"))
	}
	if resource.TableName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TableName'"))
	}
	return errs
}
