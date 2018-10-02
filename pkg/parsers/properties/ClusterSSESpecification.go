package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ClusterSSESpecification Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dax-cluster-ssespecification.html
type ClusterSSESpecification struct {
	SSEEnabled interface{} `yaml:"SSEEnabled,omitempty"`
}

// ClusterSSESpecification validation
func (resource ClusterSSESpecification) Validate() []error {
	errors := []error{}

	return errors
}
