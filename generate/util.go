package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"text/template"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func valueStringYaml(propPackage, obj, name string, property CfnProperty) string {
	omitempty := ",omitempty"
	if property.Required {
		omitempty = ""
	}
	if len(property.PrimitiveType) > 0 {
		return name + " interface{} `yaml:" + `"` + name + omitempty + `"` + "`"
	} else if len(property.Type) > 0 && property.Type != "List" && property.Type != "Map" {
		subPropertyName := propertyNameFromResourceType(obj, property.Type)
		return name + " *" + propPackage + subPropertyName + " `yaml:" + `"` + name + omitempty + `"` + "`"
	}
	return name + " interface{} `yaml:" + `"` + name + omitempty + `"` + "`"
}

func validatorYaml(obj, name string, property CfnProperty) string {
	if !property.Required {
		return ""
	}

	buf := bytes.NewBufferString("")
	t := template.Must(template.New("").Parse(validatorTemplate))
	err := t.Execute(buf, map[string]interface{}{
		"Name":          name,
		"PrimitiveType": len(property.PrimitiveType) > 0,
		"ListMapType":   len(property.Type) > 0 && property.Type != "List" && property.Type != "Map",
	})
	checkError(err)
	return buf.String()
}

func sortSpecList(specList map[string]string) []string {
	regions := make([]string, len(specList))
	i := 0
	for region := range specList {
		regions[i] = region
		i++
	}
	sort.Strings(regions)
	return regions
}

func sortProperties(properties map[string]CfnProperty) []NamedCfnProperty {
	primitives := []NamedCfnProperty{}
	nonPrimitives := []NamedCfnProperty{}
	for name, property := range properties {
		namedProperty := NamedCfnProperty{CfnProperty: property, name: name}
		if len(property.PrimitiveType) > 0 {
			primitives = append(primitives, namedProperty)
		} else {
			nonPrimitives = append(nonPrimitives, namedProperty)
		}
	}
	sort.Sort(ByName(primitives))
	sort.Sort(ByName(nonPrimitives))
	sort.Sort(ByType(nonPrimitives))
	return append(primitives, nonPrimitives...)
}

func needsFmtImport(cfnType CfnType) bool {
	for _, property := range cfnType.Properties {
		if len(property.PrimitiveType) > 0 {
			if property.Required {
				return true
			}
		} else if len(property.Type) > 0 {
			if property.Required {
				return true
			}
		}
	}
	return false
}

func needsPropertiesImport(cfnType CfnType) bool {
	for _, property := range cfnType.Properties {
		if len(property.Type) > 0 && property.Type != "List" && property.Type != "Map" {
			return true
		}
	}
	return false
}

func sortAttributeNames(attributes map[string]CfnAttribute) []string {
	names := make([]string, len(attributes))
	i := 0
	for name := range attributes {
		names[i] = name
		i++
	}
	sort.Strings(names)
	return names
}

func sortTypeNames(types map[string]CfnType) []string {
	names := make([]string, len(types))
	i := 0
	for name := range types {
		names[i] = name
		i++
	}
	sort.Strings(names)
	return names
}

func isPropertyGlobal(typeName string) bool {
	return !strings.Contains(typeName, "::")
}

func propertyNameFromPropertyType(typeName string) string {
	if isPropertyGlobal(typeName) {
		return typeName
	}
	parts := strings.Split(typeName, "::")
	subParts := strings.Split(parts[len(parts)-1], ".")
	return strings.Join(subParts, "")
}

func propertyNameFromResourceType(typeName, propertyName string) string {
	for _, v := range globalPropertyTypes {
		if v == propertyName {
			return propertyName
		}
	}
	parts := strings.Split(typeName, "::")
	subParts := strings.Split(parts[len(parts)-1], ".")
	return subParts[0] + propertyName
}

func fileNameFromCfnType(typeName string) string {
	parts := strings.Split(typeName, "::")
	return fmt.Sprint(strings.Title(parts[1]), "-", strings.Title(parts[2]))
}

func titleCaseNameFromCfnType(typeName string) string {
	parts := strings.Split(typeName, "::")
	return fmt.Sprint(strings.Title(parts[1]), strings.Title(parts[2]))
}

func generateTypes(types string) {
	err := ioutil.WriteFile("parsers/types.go", []byte(types), 0644)
	checkError(err)

}

// Sort types
type ByType []NamedCfnProperty

func (a ByType) Len() int      { return len(a) }
func (a ByType) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByType) Less(i, j int) bool {
	return a[i].Type > a[j].Type
}

type ByName []NamedCfnProperty

func (a ByName) Len() int      { return len(a) }
func (a ByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool {
	return a[i].name < a[j].name
}
