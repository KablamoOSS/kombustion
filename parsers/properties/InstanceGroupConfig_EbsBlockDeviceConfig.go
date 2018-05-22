package properties

	import "fmt"

type InstanceGroupConfig_EbsBlockDeviceConfig struct {
	
	
	VolumesPerInstance interface{} `yaml:"VolumesPerInstance,omitempty"`
	VolumeSpecification *InstanceGroupConfig_VolumeSpecification `yaml:"VolumeSpecification"`
}

func (resource InstanceGroupConfig_EbsBlockDeviceConfig) Validate() []error {
	errs := []error{}
	
	
	if resource.VolumeSpecification == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VolumeSpecification'"))
	} else {
		errs = append(errs, resource.VolumeSpecification.Validate()...)
	}
	return errs
}
