package properties

	import "fmt"

type Project_Source struct {
	
	
	
	
	
	
	BuildSpec interface{} `yaml:"BuildSpec,omitempty"`
	GitCloneDepth interface{} `yaml:"GitCloneDepth,omitempty"`
	InsecureSsl interface{} `yaml:"InsecureSsl,omitempty"`
	Location interface{} `yaml:"Location,omitempty"`
	Type interface{} `yaml:"Type"`
	Auth *Project_SourceAuth `yaml:"Auth,omitempty"`
}

func (resource Project_Source) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
