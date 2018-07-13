package lock

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/yaml"
)

// FindAndLoadLock - Search the current directory for a Lock file, and load it
// If no lock is found, return an empty Lock
func FindAndLoadLock() (lock *Lock) {
	var err error

	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		printer.Fatal(
			err,
			fmt.Sprintf(
				"kombustion.lock may need to be corrupted and needs to be rebuilt. Run `kombustion install` to fix this.",
			),
			"https://www.kombustion.io/api/cli/#install",
		)
	}

	lock, err = findAndLoadLock(path)
	if err != nil {
		printer.Fatal(
			err,
			fmt.Sprintf(
				"kombustion.lock may need to be corrupted and needs to be rebuilt. Run `kombustion install` to fix this.",
			),
			"https://www.kombustion.io/api/cli/#install",
		)
	}

	if lock == nil {
		lock = &Lock{}
		lock.Plugins = make(map[string]Plugin)
	}
	return
}

// findAndLoadLock - Search the given directory for a Lock , and load it
// This is seperated to allow for easy testing
func findAndLoadLock(path string) (lock *Lock, err error) {
	var lockPath string
	foundLock := false

	if _, err := os.Stat(path + "/kombustion.lock"); err == nil {
		lockPath = path + "/kombustion.lock"
		foundLock = true
	}

	if foundLock {
		// Read the Lock file
		data, err := ioutil.ReadFile(lockPath)
		if err != nil {
			return lock, err
		}

		lock, err := loadLockFromString(data)
		if err != nil {
			return lock, err
		}
		return lock, err
	}
	// We didn't find a lock file
	return lock, nil
}

// loadLockFromString - Load a Lock from a string into a Lock struct
func loadLockFromString(lockYaml []byte) (*Lock, error) {
	lock := Lock{}

	err := yaml.Unmarshal([]byte(lockYaml), &lock)
	if err != nil {
		return &lock, err
	}
	return &lock, nil
}
