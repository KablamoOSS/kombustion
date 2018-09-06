package manifest

import (
	"github.com/KablamoOSS/kombustion/internal/core"
	"github.com/KablamoOSS/yaml"
)

// WriteManifestToDisk - Write the final manifest to disk
func WriteManifestObject(objectStore core.ObjectStore, manifest *Manifest) error {
	// Marshal the the struct into yaml
	manifestString, err := yaml.Marshal(&manifest)
	if err != nil {
		return err
	}

	// Write the manifest
	err = objectStore.Put(manifestString, "kombustion.yaml")
	if err != nil {
		return err
	}
	return nil
}
