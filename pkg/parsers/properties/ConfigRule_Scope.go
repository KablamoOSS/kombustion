package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type ConfigRule_Scope struct {
	ComplianceResourceId    interface{} `yaml:"ComplianceResourceId,omitempty"`
	TagKey                  interface{} `yaml:"TagKey,omitempty"`
	TagValue                interface{} `yaml:"TagValue,omitempty"`
	ComplianceResourceTypes interface{} `yaml:"ComplianceResourceTypes,omitempty"`
}

func (resource ConfigRule_Scope) Validate() []error {
	errs := []error{}

	return errs
}