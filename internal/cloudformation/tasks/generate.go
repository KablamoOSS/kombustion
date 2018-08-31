package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/KablamoOSS/kombustion/internal/cloudformation"
	"github.com/KablamoOSS/kombustion/internal/core"
	yaml "github.com/KablamoOSS/yaml"
)

// GenerateTemplate and save it to disk, without upserting it
func GenerateTemplate(params cloudformation.GenerateParams) {
	output, _ := GenerateYamlTemplate(params)
	// prepareOutputDir(params.Directory)
	if params.WriteParams {
		writeParamMap(params.ObjectStore, params.Filename, params.Directory, params.ParamMap)
	}
	writeOutput(params.ObjectStore, params.Filename, params.Directory, output)
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

func writeOutput(store core.ObjectStore, file, directory string, output []byte) {
	filename := filepath.Base(file)
	basename := strings.TrimSuffix(filename, filepath.Ext(filename))
	err := store.Put(output, directory, fmt.Sprint(basename, ".yaml"))
	checkError(err)
}

func writeParamMap(store core.ObjectStore, file, directory string, params map[string]string) {
	outParams := make([]cloudformation.Parameter, 0)

	for key, value := range params {
		cfParam := cloudformation.Parameter{
			ParameterKey:   key,
			ParameterValue: value,
		}
		outParams = append(outParams, cfParam)
	}

	out, err := json.MarshalIndent(outParams, "", "  ")
	checkError(err)

	filename := filepath.Base(file)
	basename := strings.TrimSuffix(filename, filepath.Ext(filename))
	err = store.Put(out, directory, fmt.Sprint(basename, "-params.json"))
	checkError(err)
}
