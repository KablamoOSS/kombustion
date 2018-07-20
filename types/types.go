package types

type Validatable interface {
	Validate() []error
}

// CfResource - A CloudFormation Resource
type CfResource struct {
	Type       string      `yaml:"Type"`
	Properties interface{} `yaml:"Properties"`
	Condition  interface{} `yaml:"Condition,omitempty"`
	Metadata   interface{} `yaml:"Metadata,omitempty"`
	DependsOn  interface{} `yaml:"DependsOn,omitempty"`
}

// ResourceMap - a map of resouces
type ResourceMap map[string]CfResource

// TemplateObject - the type for objects in a template
type TemplateObject map[string]interface{}

// ParserFunc - a definition of the function called for resource/output/mapping parsers
type ParserFunc func(
	name string,
	data string,
) (
	// Where this parserFunc came from
	source string,

	// Outputs
	conditions TemplateObject,
	metadata TemplateObject,
	mappings TemplateObject,
	outputs TemplateObject,
	parameters TemplateObject,
	resources TemplateObject,
	transform TemplateObject,
	errors []error,
)
