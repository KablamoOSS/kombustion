package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// PipelineAddAttributes Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-addattributes.html
type PipelineAddAttributes struct {
	Attributes interface{} `yaml:"Attributes,omitempty"`
	Name       interface{} `yaml:"Name,omitempty"`
	Next       interface{} `yaml:"Next,omitempty"`
}

// PipelineAddAttributes validation
func (resource PipelineAddAttributes) Validate() []error {
	errors := []error{}

	return errors
}
