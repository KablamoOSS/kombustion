package main

// CfnSpec a struct of the entire Cloudformation Specification
type CfnSpec struct {
	PropertyTypes map[string]CfnType
	ResourceTypes map[string]CfnType
}

// CfnType is an individual Type in the Cloudformation Specification
type CfnType struct {
	Documentation string
	Properties    map[string]CfnProperty
	Attributes    map[string]CfnAttribute
}

// CfnAttribute is an attribute in the Cloudformation Specification
type CfnAttribute struct {
	PrimitiveType string
}

// CfnProperty is a property in the Cloudformation Specification
type CfnProperty struct {
	Documentation     string
	Type              string
	PrimitiveType     string
	ItemType          string
	Required          bool
	DuplicatesAllowed bool
	UpdateType        string
}

// NamedCfnProperty is a named Cloudformation Property
type NamedCfnProperty struct {
	CfnProperty
	name string
}
