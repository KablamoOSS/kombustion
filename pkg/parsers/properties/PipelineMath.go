package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// PipelineMath Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-math.html
type PipelineMath struct {
	Attribute interface{} `yaml:"Attribute,omitempty"`
	Math      interface{} `yaml:"Math,omitempty"`
	Name      interface{} `yaml:"Name,omitempty"`
	Next      interface{} `yaml:"Next,omitempty"`
}

// PipelineMath validation
func (resource PipelineMath) Validate() []error {
	errors := []error{}

	return errors
}
