package properties

	import "fmt"

type Step_HadoopJarStepConfig struct {
	
	
	
	
	Jar interface{} `yaml:"Jar"`
	MainClass interface{} `yaml:"MainClass,omitempty"`
	Args interface{} `yaml:"Args,omitempty"`
	StepProperties interface{} `yaml:"StepProperties,omitempty"`
}

func (resource Step_HadoopJarStepConfig) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.Jar == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Jar'"))
	}
	return errs
}
