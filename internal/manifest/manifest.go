package manifest

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/core"

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
			path, err := os.Getwd()
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
func CheckManifestExists(objectStore core.ObjectStore) bool {
	data, err := GetManifestObject(objectStore)
	if err != nil {
		return false
	}
	return data != nil
}

// findAndLoadManifest - Search the given directory for a manifest file, and load it
// This is separated to allow for easy testing
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

// findAndLoadManifest - Search the given directory for a manifest file, and load it
// This is separated to allow for easy testing
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
		manifest, err = loadManifestFromString(ymldata)
	} else if yamldata != nil {
		manifest, err = loadManifestFromString(yamldata)
	} else {
		return &Manifest{}, fmt.Errorf("kombustion.yaml was not found")
	}

	if err != nil {
		return &Manifest{}, err
	}
	return &manifest, nil
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
