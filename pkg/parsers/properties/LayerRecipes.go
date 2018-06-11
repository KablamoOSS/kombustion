package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// LayerRecipes Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html
type LayerRecipes struct {
	Configure interface{} `yaml:"Configure,omitempty"`
	Deploy    interface{} `yaml:"Deploy,omitempty"`
	Setup     interface{} `yaml:"Setup,omitempty"`
	Shutdown  interface{} `yaml:"Shutdown,omitempty"`
	Undeploy  interface{} `yaml:"Undeploy,omitempty"`
}

func (resource LayerRecipes) Validate() []error {
	errs := []error{}

	return errs
}
