package properties

	import "fmt"

type Cluster_ScriptBootstrapActionConfig struct {
	
	
	Path interface{} `yaml:"Path"`
	Args interface{} `yaml:"Args,omitempty"`
}

func (resource Cluster_ScriptBootstrapActionConfig) Validate() []error {
	errs := []error{}
	
	
	if resource.Path == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Path'"))
	}
	return errs
}
