package properties

	import "fmt"

type Application_Input struct {
	
	
	
	
	
	
	NamePrefix interface{} `yaml:"NamePrefix"`
	KinesisStreamsInput *Application_KinesisStreamsInput `yaml:"KinesisStreamsInput,omitempty"`
	KinesisFirehoseInput *Application_KinesisFirehoseInput `yaml:"KinesisFirehoseInput,omitempty"`
	InputSchema *Application_InputSchema `yaml:"InputSchema"`
	InputProcessingConfiguration *Application_InputProcessingConfiguration `yaml:"InputProcessingConfiguration,omitempty"`
	InputParallelism *Application_InputParallelism `yaml:"InputParallelism,omitempty"`
}

func (resource Application_Input) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	if resource.NamePrefix == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'NamePrefix'"))
	}
	if resource.InputSchema == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InputSchema'"))
	} else {
		errs = append(errs, resource.InputSchema.Validate()...)
	}
	return errs
}
