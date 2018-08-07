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
	prepareOutputDir(params.Directory)
	if params.WriteParams {
		writeParamMap(params.Filename, params.Directory, params.ParamMap)
	}
	writeOutput(params.Filename, params.Directory, output)
}

// GenerateYamlTemplate and return both the raw data as []byte, but also the cloudformation yaml object
func GenerateYamlTemplate(params cloudformation.GenerateParams) ([]byte, cloudformation.YamlCloudformation) {
	cf, err := cloudformation.GenerateYamlTemplate(params)
	checkError(err)
	output, err := yaml.Marshal(cf)
	checkError(err)
	return output, cf
}

func prepareOutputDir(directory string) {
	err := os.Mkdir(directory, 0744)
	if !os.IsExist(err) {
		checkError(err)
	}
}

func writeOutput(file, directory string, output []byte) {
	filename := filepath.Base(file)
	basename := strings.TrimSuffix(filename, filepath.Ext(filename))
	path := filepath.Join(directory, fmt.Sprint(basename, ".yaml"))
	err := ioutil.WriteFile(path, output, 0644)
	checkError(err)
}

func writeParamMap(file, directory string, params map[string]string) {
	filename := filepath.Base(file)
	basename := strings.TrimSuffix(filename, filepath.Ext(filename))
	path := filepath.Join(directory, fmt.Sprint(basename, "-params.yaml"))
	out, err := yaml.Marshal(params)
	checkError(err)

	ioutil.WriteFile(path, out, 0644)
}
