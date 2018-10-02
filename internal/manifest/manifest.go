package manifest

import (
	"fmt"

	"github.com/KablamoOSS/kombustion/internal/core"

	yaml "github.com/KablamoOSS/yaml"
)

func CheckManifestExists(objectStore core.ObjectStore, manifestLocation string) bool {
	data, err := GetManifestObject(objectStore, manifestLocation)
	if err != nil {
		return false
	}
	return data != nil
}

func GetManifestObject(objectStore core.ObjectStore, manifestLocation string, path ...string) (*Manifest, error) {
	var manifest Manifest
	var err error

	// Read the manifest file
	yamlpath := append(path, manifestLocation)
	yamldata, err := objectStore.Get(yamlpath[0], yamlpath[1:]...)
	if err != nil {
		return &Manifest{}, fmt.Errorf("%v: %v", manifestLocation, err)
	}

	if yamldata != nil {
		manifest, err = unmarshalManifest(yamldata)
	} else {
		return &Manifest{}, fmt.Errorf("%v was not found", manifestLocation)
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
