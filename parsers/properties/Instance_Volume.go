package properties

	import "fmt"

type Instance_Volume struct {
	
	
	Device interface{} `yaml:"Device"`
	VolumeId interface{} `yaml:"VolumeId"`
}

func (resource Instance_Volume) Validate() []error {
	errs := []error{}
	
	
	if resource.Device == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Device'"))
	}
	if resource.VolumeId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VolumeId'"))
	}
	return errs
}
