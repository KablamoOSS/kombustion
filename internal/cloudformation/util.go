package cloudformation

import (
	"fmt"
	"io"
	"text/template"
)

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
