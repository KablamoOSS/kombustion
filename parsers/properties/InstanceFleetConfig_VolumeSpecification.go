package properties

	import "fmt"

type InstanceFleetConfig_VolumeSpecification struct {
	
	
	
	Iops interface{} `yaml:"Iops,omitempty"`
	SizeInGB interface{} `yaml:"SizeInGB"`
	VolumeType interface{} `yaml:"VolumeType"`
}

func (resource InstanceFleetConfig_VolumeSpecification) Validate() []error {
	errs := []error{}
	
	
	
	if resource.SizeInGB == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SizeInGB'"))
	}
	if resource.VolumeType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VolumeType'"))
	}
	return errs
}
