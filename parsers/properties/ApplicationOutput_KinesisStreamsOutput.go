package properties

	import "fmt"

type ApplicationOutput_KinesisStreamsOutput struct {
	
	
	ResourceARN interface{} `yaml:"ResourceARN"`
	RoleARN interface{} `yaml:"RoleARN"`
}

func (resource ApplicationOutput_KinesisStreamsOutput) Validate() []error {
	errs := []error{}
	
	
	if resource.ResourceARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ResourceARN'"))
	}
	if resource.RoleARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleARN'"))
	}
	return errs
}
