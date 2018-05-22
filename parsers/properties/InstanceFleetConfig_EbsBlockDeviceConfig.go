package properties

	import "fmt"

type InstanceFleetConfig_EbsBlockDeviceConfig struct {
	
	
	VolumesPerInstance interface{} `yaml:"VolumesPerInstance,omitempty"`
	VolumeSpecification *InstanceFleetConfig_VolumeSpecification `yaml:"VolumeSpecification"`
}

func (resource InstanceFleetConfig_EbsBlockDeviceConfig) Validate() []error {
	errs := []error{}
	
	
	if resource.VolumeSpecification == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VolumeSpecification'"))
	} else {
		errs = append(errs, resource.VolumeSpecification.Validate()...)
	}
	return errs
}
