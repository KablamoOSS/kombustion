package types

type Validatable interface {
	Validate() []error
}

type ValueMap map[string]interface{}

// ParserFunc - the func definition for resource/output parsers
type ParserFunc func(string, string) (ValueMap, error)

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

// PluginHelp - a set of available documentation fields
type PluginHelp struct {
	Description  string
	TypeMappings []TypeMapping
	Snippets     []string
}

// TypeMapping - recursive list of types with its associated config object
type TypeMapping struct {
	Name        string
	Description string
	Config      interface{}
}
