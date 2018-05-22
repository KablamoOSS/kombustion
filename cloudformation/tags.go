package cloudformation

import (
	"reflect"

	yaml "github.com/KablamoOSS/yaml"
)

type intrinsicFn string

func (name intrinsicFn) UnmarshalYAMLTag(t string, out reflect.Value) reflect.Value {
	output := reflect.ValueOf(make(map[string]interface{}))
	output.SetMapIndex(reflect.ValueOf(string(name)), out)
	return output
}

func registerYamlTagUnmarshalers() {
	yaml.RegisterTagUnmarshaler("!Ref", intrinsicFn("Ref"))
	yaml.RegisterTagUnmarshaler("!Base64", intrinsicFn("Fn::Base64"))
	yaml.RegisterTagUnmarshaler("!FindInMap", intrinsicFn("Fn::FindInMap"))
	yaml.RegisterTagUnmarshaler("!Join", intrinsicFn("Fn::Join"))
	yaml.RegisterTagUnmarshaler("!Select", intrinsicFn("Fn::Select"))
	yaml.RegisterTagUnmarshaler("!Split", intrinsicFn("Fn::Split"))
	yaml.RegisterTagUnmarshaler("!Sub", intrinsicFn("Fn::Sub"))
	yaml.RegisterTagUnmarshaler("!And", intrinsicFn("Fn::And"))
	yaml.RegisterTagUnmarshaler("!Equals", intrinsicFn("Fn::Equals"))
	yaml.RegisterTagUnmarshaler("!If", intrinsicFn("Fn::If"))
	yaml.RegisterTagUnmarshaler("!Not", intrinsicFn("Fn::Not"))
	yaml.RegisterTagUnmarshaler("!Or", intrinsicFn("Fn::Or"))
	yaml.RegisterTagUnmarshaler("!GetAtt", intrinsicFn("Fn::GetAtt"))
	yaml.RegisterTagUnmarshaler("!GetAZs", intrinsicFn("Fn::GetAZs"))
	yaml.RegisterTagUnmarshaler("!ImportValue", intrinsicFn("Fn::ImportValue"))
	yaml.RegisterTagUnmarshaler("!Cidr", intrinsicFn("Fn::Cidr"))
}
