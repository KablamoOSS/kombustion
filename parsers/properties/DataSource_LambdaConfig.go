package properties

	import "fmt"

type DataSource_LambdaConfig struct {
	
	LambdaFunctionArn interface{} `yaml:"LambdaFunctionArn"`
}

func (resource DataSource_LambdaConfig) Validate() []error {
	errs := []error{}
	
	if resource.LambdaFunctionArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'LambdaFunctionArn'"))
	}
	return errs
}
