package properties

	import "fmt"

type SpotFleet_BlockDeviceMapping struct {
	
	
	
	
	DeviceName interface{} `yaml:"DeviceName"`
	NoDevice interface{} `yaml:"NoDevice,omitempty"`
	VirtualName interface{} `yaml:"VirtualName,omitempty"`
	Ebs *SpotFleet_EbsBlockDevice `yaml:"Ebs,omitempty"`
}

func (resource SpotFleet_BlockDeviceMapping) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.DeviceName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DeviceName'"))
	}
	return errs
}
