package tasks

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/KablamoOSS/kombustion/internal/cloudformation"
	yaml "github.com/KablamoOSS/yaml"
)

// GenerateTemplate and save it to disk, without upserting it
func GenerateTemplate(params cloudformation.GenerateParams) {
	output, _ := GenerateYamlTemplate(params)
	writeOutput(params.Filename, output)
}

// GenerateYamlTemplate and return both the raw data as []byte, but also the cloudformation yaml object
func GenerateYamlTemplate(params cloudformation.GenerateParams) ([]byte, cloudformation.YamlCloudformation) {
	cf, err := cloudformation.GenerateYamlTemplate(params)
	checkError(err)
	output, err := yaml.Marshal(cf)
	checkError(err)
	return output, cf
}

func writeOutput(file string, output []byte) {
	filename := filepath.Base(file)
	basename := strings.TrimSuffix(filename, filepath.Ext(filename))
	path := fmt.Sprint("compiled/", basename, ".yaml")
	os.Mkdir("./compiled", 0744)
	err := ioutil.WriteFile(path, output, 0644)
	checkError(err)
}
