package cloudformation

import (
	"os"
	"testing"

	yaml "github.com/KablamoOSS/yaml"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	registerYamlTagUnmarshalers()
	os.Exit(m.Run())
}

func TestTags_nested(t *testing.T) {
	inputYaml := `Value: !Split [ ",", !Ref MyResource ]`

	expectedValue := map[interface{}]interface{}{
		"Value": map[string]interface{}{
			"Fn::Split": []interface{}{
				",",
				map[string]interface{}{"Ref": "MyResource"},
			},
		},
	}

	var result interface{}
	err := yaml.Unmarshal([]byte(inputYaml), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedValue, result)
}
