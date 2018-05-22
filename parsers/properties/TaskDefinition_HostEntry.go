package properties

	import "fmt"

type TaskDefinition_HostEntry struct {
	
	
	Hostname interface{} `yaml:"Hostname"`
	IpAddress interface{} `yaml:"IpAddress"`
}

func (resource TaskDefinition_HostEntry) Validate() []error {
	errs := []error{}
	
	
	if resource.Hostname == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Hostname'"))
	}
	if resource.IpAddress == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IpAddress'"))
	}
	return errs
}
