package manifest

import (
	"io/ioutil"

	"github.com/KablamoOSS/yaml"
)

// WriteManifestToDisk - Write the final manifest to disk
func WriteManifestToDisk(manifest *Manifest) error {

	// Mashall the the struct into yaml
	manifestString, err := yaml.Marshal(&manifest)
	if err != nil {
		return err
	}

	// Write the manifest
	err = ioutil.WriteFile("kombustion.yaml", manifestString, 0644)
	if err != nil {
		return err
	}
	return nil
}
