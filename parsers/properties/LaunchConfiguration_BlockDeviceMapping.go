package properties

	import "fmt"

type LaunchConfiguration_BlockDeviceMapping struct {
	
	
	
	
	DeviceName interface{} `yaml:"DeviceName"`
	NoDevice interface{} `yaml:"NoDevice,omitempty"`
	VirtualName interface{} `yaml:"VirtualName,omitempty"`
	Ebs *LaunchConfiguration_BlockDevice `yaml:"Ebs,omitempty"`
}

func (resource LaunchConfiguration_BlockDeviceMapping) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.DeviceName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DeviceName'"))
	}
	return errs
}
