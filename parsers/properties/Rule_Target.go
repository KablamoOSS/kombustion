package properties

	import "fmt"

type Rule_Target struct {
	
	
	
	
	
	
	
	
	
	Arn interface{} `yaml:"Arn"`
	Id interface{} `yaml:"Id"`
	Input interface{} `yaml:"Input,omitempty"`
	InputPath interface{} `yaml:"InputPath,omitempty"`
	RoleArn interface{} `yaml:"RoleArn,omitempty"`
	RunCommandParameters *Rule_RunCommandParameters `yaml:"RunCommandParameters,omitempty"`
	KinesisParameters *Rule_KinesisParameters `yaml:"KinesisParameters,omitempty"`
	InputTransformer *Rule_InputTransformer `yaml:"InputTransformer,omitempty"`
	EcsParameters *Rule_EcsParameters `yaml:"EcsParameters,omitempty"`
}

func (resource Rule_Target) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	if resource.Arn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Arn'"))
	}
	if resource.Id == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Id'"))
	}
	return errs
}
