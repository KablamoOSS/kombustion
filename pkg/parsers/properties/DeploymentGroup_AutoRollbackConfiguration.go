package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type DeploymentGroup_AutoRollbackConfiguration struct {
	Enabled interface{} `yaml:"Enabled,omitempty"`
	Events  interface{} `yaml:"Events,omitempty"`
}

func (resource DeploymentGroup_AutoRollbackConfiguration) Validate() []error {
	errs := []error{}

	return errs
}