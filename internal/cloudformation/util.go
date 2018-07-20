package cloudformation

import (
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"text/template"

	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ResolveEnvironment loads the correct configs/environment.yaml ValueMap for the specified env
func ResolveEnvironment(envFile string, env string) types.TemplateObject {
	// TODO: this func is deprecated
	if len(envFile) == 0 {
		envFile = "./environment.yml"
	}

	data, err := ioutil.ReadFile(envFile)
	if err != nil {
		return make(types.TemplateObject)
	}

	var envMap map[string]types.TemplateObject
	yaml.Unmarshal(data, &envMap)

	if env, ok := envMap[env]; ok {
		return env
	}

	return types.TemplateObject{}
}

/*
	loadFiles
	Load a map of files from a folder (mapping: <filename>:<filedata> )
*/
func loadFiles(path string) (files map[string][]byte, err error) {
	files = make(map[string][]byte)
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		if file.IsDir() {
			continue
		}
		data, err := ioutil.ReadFile(filepath.Join(path, file.Name()))
		if err != nil {
			return files, err
		}
		files[file.Name()] = data
	}
	return files, nil
}

/*
	executeTemplate
	execute a fresh template from a templateDefinition
	doesn't use the common template.
*/
func executeTemplate(w io.Writer, templateDefinition []byte, data interface{}) error {
	t, err := template.New("cfn").Parse(string(templateDefinition))
	if err != nil {
		return err
	}

	return t.Execute(w, data)
}

/*
	fixYamlKeys
	recursively forces map[interface{}]interface{} types into map[string]interface{}
	to support json.Marshal
*/
func fixYamlKeys(o interface{}) interface{} {
	switch obj := o.(type) {

	case map[interface{}]interface{}:
		fixed := make(map[string]interface{})
		for k, v := range obj {
			fixed[fmt.Sprintf("%v", k)] = fixYamlKeys(v)
		}
		return fixed

	case []interface{}:
		fixed := make([]interface{}, len(obj))
		for i, v := range obj {
			fixed[i] = fixYamlKeys(v)
		}
		return fixed

	default:
		return o
	}
}
