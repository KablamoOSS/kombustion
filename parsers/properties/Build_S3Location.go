package properties

	import "fmt"

type Build_S3Location struct {
	
	
	
	Bucket interface{} `yaml:"Bucket"`
	Key interface{} `yaml:"Key"`
	RoleArn interface{} `yaml:"RoleArn"`
}

func (resource Build_S3Location) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Bucket == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Bucket'"))
	}
	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	return errs
}
