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
	prepareOutputDir(params.Filename)
	if params.WriteParams {
		writeParamMap(params.Filename, params.ParamMap)
	}
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

func prepareOutputDir(file string) {
	err := os.Mkdir("./compiled", 0744)
	if !os.IsExist(err) {
		checkError(err)
	}
}

func writeOutput(file string, output []byte) {
	filename := filepath.Base(file)
	basename := strings.TrimSuffix(filename, filepath.Ext(filename))
	path := fmt.Sprint("compiled/", basename, ".yaml")
	err := ioutil.WriteFile(path, output, 0644)
	checkError(err)
}

func writeParamMap(file string, params map[string]string) {
	filename := filepath.Base(file)
	basename := strings.TrimSuffix(filename, filepath.Ext(filename))
	path := fmt.Sprint("compiled/", basename, "-params.yaml")
	out, err := yaml.Marshal(params)
	checkError(err)

	ioutil.WriteFile(path, out, 0644)
}
