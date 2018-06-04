package types

type Validatable interface {
	Validate() []error
}

// CfResource - A Cloudformation Resource
type CfResource struct {
	Type       string      `yaml:"Type"`
	Properties interface{} `yaml:"Properties"`
	Condition  interface{} `yaml:"Condition,omitempty"`
	Metadata   interface{} `yaml:"Metadata,omitempty"`
	DependsOn  interface{} `yaml:"DependsOn,omitempty"`
}

// ResourceMap - a map of resouces
type ResourceMap map[string]CfResource

// TemplateObject - the type for outputs, resources, mappings
type TemplateObject map[string]interface{}

// ParserFunc - a definition of the function called for resource/output/mapping parsers
type ParserFunc func(ctx map[string]interface{}, name string, data string) (TemplateObject, error)
