package properties

	import "fmt"

type Cluster_BootstrapActionConfig struct {
	
	
	Name interface{} `yaml:"Name"`
	ScriptBootstrapAction *Cluster_ScriptBootstrapActionConfig `yaml:"ScriptBootstrapAction"`
}

func (resource Cluster_BootstrapActionConfig) Validate() []error {
	errs := []error{}
	
	
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.ScriptBootstrapAction == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ScriptBootstrapAction'"))
	} else {
		errs = append(errs, resource.ScriptBootstrapAction.Validate()...)
	}
	return errs
}
