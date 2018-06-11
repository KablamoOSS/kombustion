package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TableSerdeInfo Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-serdeinfo.html
type TableSerdeInfo struct {
	Name                 interface{} `yaml:"Name,omitempty"`
	Parameters           interface{} `yaml:"Parameters,omitempty"`
	SerializationLibrary interface{} `yaml:"SerializationLibrary,omitempty"`
}

func (resource TableSerdeInfo) Validate() []error {
	errs := []error{}

	return errs
}
