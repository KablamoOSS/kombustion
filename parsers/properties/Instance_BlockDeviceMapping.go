package properties

	import "fmt"

type Instance_BlockDeviceMapping struct {
	
	
	
	
	DeviceName interface{} `yaml:"DeviceName"`
	VirtualName interface{} `yaml:"VirtualName,omitempty"`
	NoDevice *Instance_NoDevice `yaml:"NoDevice,omitempty"`
	Ebs *Instance_Ebs `yaml:"Ebs,omitempty"`
}

func (resource Instance_BlockDeviceMapping) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.DeviceName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DeviceName'"))
	}
	return errs
}
