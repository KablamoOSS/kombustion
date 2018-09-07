package manifest

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/internal/core"

	yaml "github.com/KablamoOSS/yaml"
)

func CheckManifestExists(objectStore core.ObjectStore) bool {
	data, err := GetManifestObject(objectStore)
	if err != nil {
		return false
	}
	return data != nil
}

func GetManifestObject(objectStore core.ObjectStore, path ...string) (*Manifest, error) {
	var manifest Manifest
	var err error

	ymlpath := append(path, "kombustion.yml")
	ymldata, err := objectStore.Get(ymlpath[0], ymlpath[1:]...)
	if err != nil {
		return &Manifest{}, fmt.Errorf("kombustion.yml: %v", err)
	}

	// Read the manifest file
	yamlpath := append(path, "kombustion.yaml")
	yamldata, err := objectStore.Get(yamlpath[0], yamlpath[1:]...)
	if err != nil {
		return &Manifest{}, fmt.Errorf("kombustion.yaml: %v", err)
	}

	if ymldata != nil && yamldata != nil {
		return &Manifest{}, fmt.Errorf("there are both kombustion.yaml && kombustion.yml files, please remove one")
	} else if ymldata != nil {
		manifest, err = unmarshalManifest(ymldata)
	} else if yamldata != nil {
		manifest, err = unmarshalManifest(yamldata)
	} else {
		return &Manifest{}, fmt.Errorf("kombustion.yaml was not found")
	}

	if err != nil {
		return &Manifest{}, err
	}
	return &manifest, nil
}

func unmarshalManifest(manifestYaml []byte) (Manifest, error) {
	manifest := Manifest{}

	err := yaml.Unmarshal([]byte(manifestYaml), &manifest)
	if err != nil {
		return Manifest{}, err
	}
	return manifest, nil
}
