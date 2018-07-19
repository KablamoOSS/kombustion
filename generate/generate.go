package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"text/template"
)

var globalPropertyTypes []string

func buildYamlParsers(cfnSpec CfnSpec) {
	// check for global types
	globalPropertyTypes := []string{}
	for k := range cfnSpec.PropertyTypes {
		if isPropertyGlobal(k) {
			globalPropertyTypes = append(globalPropertyTypes, k)
		}
	}
	var resourceTypes []string

	log.Println("Generate Resource Map")
	resourceParsersObject := buildParserMapping(cfnSpec, "resources")
	filePath := fmt.Sprintf("%vresources.go", parsersDir)
	formatted, err := format.Source([]byte(resourceParsersObject))
	checkError(err)
	err = ioutil.WriteFile(filePath, formatted, 0644)
	checkError(err)

	log.Println("Generate Output Map")
	outputParsersObject := buildParserMapping(cfnSpec, "outputs")
	filePath = fmt.Sprintf("%voutput.go", parsersDir)
	formatted, err = format.Source([]byte(outputParsersObject))
	checkError(err)
	err = ioutil.WriteFile(filePath, formatted, 0644)
	checkError(err)

	log.Println("Generate Properties")
	// properties
	for k, cfnType := range cfnSpec.PropertyTypes {
		propertyObject := buildPropertyYaml(k, cfnType)
		filePath := fmt.Sprintf("%v%v.go", propertiesDir, propertyNameFromPropertyType(k))
		formatted, err := format.Source([]byte(propertyObject))
		checkError(err)
		err = ioutil.WriteFile(filePath, formatted, 0644)

		checkError(err)
	}

	log.Println("Generate Resources")
	// Resource parsers
	for k, cfnType := range cfnSpec.ResourceTypes {
		resourceObject := buildResourceYaml(k, cfnType)
		filePath := fmt.Sprintf("%v%v.go", resourcesDir, fileNameFromCfnType(k))
		formatted, err := format.Source([]byte(resourceObject))
		checkError(err)
		err = ioutil.WriteFile(filePath, formatted, 0644)
		checkError(err)

		resourceTypes = append(resourceTypes, titleCaseNameFromCfnType(k))
	}

	log.Println("Generate Outputs")
	// Output parsers
	for k, cfnType := range cfnSpec.ResourceTypes {
		outputObject := buildOutputYaml(k, cfnType)
		filePath := fmt.Sprintf("%v%v.go", outputsDir, fileNameFromCfnType(k))
		formatted, err := format.Source([]byte(outputObject))
		checkError(err)
		err = ioutil.WriteFile(filePath, formatted, 0644)

		checkError(err)
	}
}

func buildParserMapping(cfnSpec CfnSpec, packageName string) string {
	resourceTypes := make(map[string]string)
	for _, k := range sortTypeNames(cfnSpec.ResourceTypes) {
		name := titleCaseNameFromCfnType(k)
		resourceTypes[k] = name
	}

	buf := bytes.NewBufferString("")
	t := template.Must(template.New("").Parse(parserMapTemplate))
	err := t.Execute(buf, map[string]interface{}{
		"PackageNameTitle": strings.Title(packageName),
		"PackageName":      packageName,
		"ResourceTypes":    resourceTypes,
		"MainPackageName":  mainPackageName,
	})
	checkError(err)
	return buf.String()
}

func buildPropertyYaml(obj string, cfnType CfnType) string {
	propertyStrings := make([]string, len(cfnType.Properties))
	validatorStrings := make([]string, len(cfnType.Properties))
	for _, property := range sortProperties(cfnType.Properties) {
		if str := valueStringYaml("", obj, property.name, property.CfnProperty); len(str) > 0 {
			propertyStrings = append(propertyStrings, str)
		}
	}
	for _, property := range sortProperties(cfnType.Properties) {
		if str := validatorYaml(obj, property.name, property.CfnProperty); len(str) > 0 {
			validatorStrings = append(validatorStrings, str)
		}
	}

	buf := bytes.NewBufferString("")
	t := template.Must(template.New("").Parse(propertyTemplate))
	err := t.Execute(buf, map[string]interface{}{
		"PropertyName":     propertyNameFromPropertyType(obj),
		"Documentation":    cfnType.Documentation,
		"PropertyStrings":  propertyStrings,
		"ValidatorStrings": validatorStrings,
		"NeedsFmtImport":   needsFmtImport(cfnType),
		"MainPackageName":  mainPackageName,
	})
	checkError(err)
	return buf.String()
}

func buildResourceYaml(obj string, cfnType CfnType) string {
	propertyStrings := make([]string, 0, len(cfnType.Properties))
	validatorStrings := make([]string, 0, len(cfnType.Properties))
	for _, property := range sortProperties(cfnType.Properties) {
		if str := valueStringYaml("properties.", obj, property.name, property.CfnProperty); len(str) > 0 {
			propertyStrings = append(propertyStrings, str)
		}
	}
	for _, property := range sortProperties(cfnType.Properties) {
		if str := validatorYaml(obj, property.name, property.CfnProperty); len(str) > 0 {
			validatorStrings = append(validatorStrings, str)
		}
	}

	buf := bytes.NewBufferString("")
	t := template.Must(template.New("").Parse(resourceTemplate))
	err := t.Execute(buf, map[string]interface{}{
		"Type":                  obj,
		"ResourceName":          titleCaseNameFromCfnType(obj),
		"Documentation":         cfnType.Documentation,
		"PropertyStrings":       propertyStrings,
		"ValidatorStrings":      validatorStrings,
		"NeedsFmtImport":        needsFmtImport(cfnType),
		"NeedsPropertiesImport": needsPropertiesImport(cfnType),
		"MainPackageName":       mainPackageName,
	})
	checkError(err)
	return buf.String()
}

func buildOutputYaml(obj string, cfnType CfnType) string {
	alnumRegex, _ := regexp.Compile("[^a-zA-Z0-9]+")
	attributes := make(map[string]string)
	for _, attName := range sortAttributeNames(cfnType.Attributes) {
		attributes[attName] = alnumRegex.ReplaceAllString(attName, "")
	}

	buf := bytes.NewBufferString("")
	t := template.Must(template.New("").Parse(outputTemplate))
	err := t.Execute(buf, map[string]interface{}{
		"ResourceName":    titleCaseNameFromCfnType(obj),
		"Attributes":      attributes,
		"Documentation":   cfnType.Documentation,
		"MainPackageName": mainPackageName,
	})
	checkError(err)
	return buf.String()
}
