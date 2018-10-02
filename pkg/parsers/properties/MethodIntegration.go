package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// MethodIntegration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html
type MethodIntegration struct {
	CacheNamespace        interface{} `yaml:"CacheNamespace,omitempty"`
	ConnectionId          interface{} `yaml:"ConnectionId,omitempty"`
	ConnectionType        interface{} `yaml:"ConnectionType,omitempty"`
	ContentHandling       interface{} `yaml:"ContentHandling,omitempty"`
	Credentials           interface{} `yaml:"Credentials,omitempty"`
	IntegrationHttpMethod interface{} `yaml:"IntegrationHttpMethod,omitempty"`
	PassthroughBehavior   interface{} `yaml:"PassthroughBehavior,omitempty"`
	TimeoutInMillis       interface{} `yaml:"TimeoutInMillis,omitempty"`
	Type                  interface{} `yaml:"Type,omitempty"`
	Uri                   interface{} `yaml:"Uri,omitempty"`
	RequestParameters     interface{} `yaml:"RequestParameters,omitempty"`
	RequestTemplates      interface{} `yaml:"RequestTemplates,omitempty"`
	CacheKeyParameters    interface{} `yaml:"CacheKeyParameters,omitempty"`
	IntegrationResponses  interface{} `yaml:"IntegrationResponses,omitempty"`
}

// MethodIntegration validation
func (resource MethodIntegration) Validate() []error {
	errors := []error{}

	return errors
}
