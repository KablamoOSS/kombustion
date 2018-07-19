package main

import (
	"log"
	"os"
)

/*
	generate
	Auto generates go parser code from the cloudformation spec
*/

const typesDir = "./pkg/parsers/types/"
const parsersDir = "./pkg/parsers/"
const propertiesDir = "./pkg/parsers/properties/"
const outputsDir = "./pkg/parsers/outputs/"
const resourcesDir = "./pkg/parsers/resources/"

const mainPackageName = "parsers"
const sourceDir = "./generate/source/"

func init() {
	log.Println("Creating directories")

	os.RemoveAll(parsersDir)

	os.Mkdir(parsersDir, 0744)
	os.Mkdir(outputsDir, 0744)
	os.Mkdir(propertiesDir, 0744)
	os.Mkdir(resourcesDir, 0744)
	log.Println("Created directories")
}

func main() {

	var cfnSpec CfnSpec

	// Download the latest Cloudformation Specification
	log.Println("Download the latest Cloudformation Specification")
	// getCloudformationSpecification(sourceDir, cfnEndpoints)

	// Load and de-dupe the specification
	log.Println("Load and de-dupe the specification")
	cfnSpec = buildUniqueSet(sourceDir, cfnEndpoints)

	// Build the Yaml Parsers
	log.Println("Build the Yaml Parsers")
	buildYamlParsers(cfnSpec)

}
