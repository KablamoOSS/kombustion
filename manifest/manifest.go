package manifest

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	yaml "github.com/KablamoOSS/yaml"
)

// FindAndLoadManifest - Search the current directory for a manifest file, and load it
func FindAndLoadManifest() (Manifest, error) {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return findAndLoadManifest(path)
}

// findAndLoadManifest - Search the given directory for a manifest file, and load it
// This is seperated to allow for easy testing
func findAndLoadManifest(path string) (Manifest, error) {
	var manifestPath string
	foundManifest := false

	// Support for .yml
	if _, err := os.Stat(path + "/kombustion.yml"); err == nil {
		manifestPath = path + "/kombustion.yml"
		foundManifest = true
	}

	if _, err := os.Stat(path + "/kombustion.yaml"); err == nil {
		if manifestPath == path+"/kombustion.yml" {
			return Manifest{}, fmt.Errorf("there are both kombustion.yaml && kombustion.yml files, please remove one")
		}
		manifestPath = path + "/kombustion.yaml"
		foundManifest = true
	}

	if foundManifest {
		// Read the manifest file
		data, err := ioutil.ReadFile(manifestPath)
		if err != nil {
			return Manifest{}, err
		}

		manifest, err := loadManifestFromString(data)
		if err != nil {
			return Manifest{}, err
		}
		return manifest, nil

	}
	return Manifest{}, fmt.Errorf("no kombustion.yaml manifest file found")
}

// loadManifestFromString - Load a manifest from a string into a Manifest struct
func loadManifestFromString(manifestYaml []byte) (Manifest, error) {
	manifest := Manifest{}

	err := yaml.Unmarshal([]byte(manifestYaml), &manifest)
	if err != nil {
		return Manifest{}, err
	}
	return manifest, nil
}
