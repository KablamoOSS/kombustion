package properties

	import "fmt"

type JobDefinition_Ulimit struct {
	
	
	
	HardLimit interface{} `yaml:"HardLimit"`
	Name interface{} `yaml:"Name"`
	SoftLimit interface{} `yaml:"SoftLimit"`
}

func (resource JobDefinition_Ulimit) Validate() []error {
	errs := []error{}
	
	
	
	if resource.HardLimit == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HardLimit'"))
	}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.SoftLimit == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SoftLimit'"))
	}
	return errs
}
