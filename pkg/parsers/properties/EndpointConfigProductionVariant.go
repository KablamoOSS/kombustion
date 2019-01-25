package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// EndpointConfigProductionVariant Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html
type EndpointConfigProductionVariant struct {
	AcceleratorType      interface{} `yaml:"AcceleratorType,omitempty"`
	InitialInstanceCount interface{} `yaml:"InitialInstanceCount"`
	InitialVariantWeight interface{} `yaml:"InitialVariantWeight"`
	InstanceType         interface{} `yaml:"InstanceType"`
	ModelName            interface{} `yaml:"ModelName"`
	VariantName          interface{} `yaml:"VariantName"`
}

// EndpointConfigProductionVariant validation
func (resource EndpointConfigProductionVariant) Validate() []error {
	errors := []error{}

	return errors
}
