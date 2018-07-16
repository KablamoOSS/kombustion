package manifest

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/KablamoOSS/go-cli-printer"

	yaml "github.com/KablamoOSS/yaml"
)

// Once loaded, we keep the manifest in memory
var (
	loadedManifest *Manifest
	once           sync.Once
)

// FindAndLoadManifest - Search the current directory for a manifest file, and load it
func FindAndLoadManifest() *Manifest {
	if loadedManifest == nil {
		once.Do(func() {
			path, err := filepath.Abs(filepath.Dir(os.Args[0]))
			if err != nil {
				printer.Fatal(
					err,
					"If you want to re-initialise your kombustion.yaml file, first remove it.",
					"https://www.kombustion.io/api/manifest/",
				)
			}
			manifest, err := findAndLoadManifest(path)
			if err != nil {
				printer.Fatal(
					err,
					"If you want to re-initialise your kombustion.yaml file, first remove it.",
					"https://www.kombustion.io/api/manifest/",
				)
			}
			loadedManifest = &manifest
		})
	}
	return loadedManifest
}

// CheckManifestExists to determine if there is a manifest in the current dir
func CheckManifestExists() bool {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return false
	}
	_, err = findAndLoadManifest(path)
	if err != nil {
		return false
	}
	return true
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
	return Manifest{}, fmt.Errorf("kombustion.yaml was not found")
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
