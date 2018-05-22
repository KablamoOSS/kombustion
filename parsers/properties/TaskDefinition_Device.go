package properties

	import "fmt"

type TaskDefinition_Device struct {
	
	
	
	ContainerPath interface{} `yaml:"ContainerPath,omitempty"`
	HostPath interface{} `yaml:"HostPath"`
	Permissions interface{} `yaml:"Permissions,omitempty"`
}

func (resource TaskDefinition_Device) Validate() []error {
	errs := []error{}
	
	
	
	if resource.HostPath == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HostPath'"))
	}
	return errs
}
