package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// PipelineDeviceShadowEnrich Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceshadowenrich.html
type PipelineDeviceShadowEnrich struct {
	Attribute interface{} `yaml:"Attribute,omitempty"`
	Name      interface{} `yaml:"Name,omitempty"`
	Next      interface{} `yaml:"Next,omitempty"`
	RoleArn   interface{} `yaml:"RoleArn,omitempty"`
	ThingName interface{} `yaml:"ThingName,omitempty"`
}

// PipelineDeviceShadowEnrich validation
func (resource PipelineDeviceShadowEnrich) Validate() []error {
	errors := []error{}

	return errors
}
