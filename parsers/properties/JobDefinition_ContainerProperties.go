package properties

	import "fmt"

type JobDefinition_ContainerProperties struct {
	
	
	
	
	
	
	
	
	
	
	
	
	Image interface{} `yaml:"Image"`
	JobRoleArn interface{} `yaml:"JobRoleArn,omitempty"`
	Memory interface{} `yaml:"Memory"`
	Privileged interface{} `yaml:"Privileged,omitempty"`
	ReadonlyRootFilesystem interface{} `yaml:"ReadonlyRootFilesystem,omitempty"`
	User interface{} `yaml:"User,omitempty"`
	Vcpus interface{} `yaml:"Vcpus"`
	Command interface{} `yaml:"Command,omitempty"`
	Environment interface{} `yaml:"Environment,omitempty"`
	MountPoints interface{} `yaml:"MountPoints,omitempty"`
	Ulimits interface{} `yaml:"Ulimits,omitempty"`
	Volumes interface{} `yaml:"Volumes,omitempty"`
}

func (resource JobDefinition_ContainerProperties) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	
	
	if resource.Image == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Image'"))
	}
	if resource.Memory == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Memory'"))
	}
	if resource.Vcpus == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Vcpus'"))
	}
	return errs
}
