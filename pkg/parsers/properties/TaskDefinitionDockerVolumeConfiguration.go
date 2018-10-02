package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TaskDefinitionDockerVolumeConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-dockervolumeconfiguration.html
type TaskDefinitionDockerVolumeConfiguration struct {
	Autoprovision interface{} `yaml:"Autoprovision,omitempty"`
	Driver        interface{} `yaml:"Driver,omitempty"`
	Scope         interface{} `yaml:"Scope,omitempty"`
	DriverOpts    interface{} `yaml:"DriverOpts,omitempty"`
	Labels        interface{} `yaml:"Labels,omitempty"`
}

// TaskDefinitionDockerVolumeConfiguration validation
func (resource TaskDefinitionDockerVolumeConfiguration) Validate() []error {
	errors := []error{}

	return errors
}
