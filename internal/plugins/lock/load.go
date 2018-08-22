package lock

import (
	"io/ioutil"
	"os"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/yaml"
)

// FindAndLoadLock - Search the current directory for a Lock file, and load it
// If no lock is found, return an empty Lock
func FindAndLoadLock() *Lock {
	var err error

	path, err := os.Getwd()
	if err != nil {
		// os.Getwd failure conditions are likely to be OS dependant
		printer.Fatal(
			err,
			"Check your operating environment has a valid working directory",
			"",
		)
	}

	lock, err := findAndLoadLock(path)
	if err != nil {
		printer.Fatal(
			err,
			"kombustion.lock may be corrupted and needs to be rebuilt. Run `kombustion install` to fix this.",
			"https://www.kombustion.io/api/cli/#install",
		)
	}

	if lock == nil {
		lock = &Lock{}
		lock.Plugins = make(map[string]Plugin)
	}

	return lock
}

// findAndLoadLock - Search the given directory for a Lock , and load it
// This is seperated to allow for easy testing
func findAndLoadLock(path string) (*Lock, error) {
	lockPath := path + "/kombustion.lock"

	if _, err := os.Stat(lockPath); err == nil {
		// Read the Lock file
		data, err := ioutil.ReadFile(lockPath)
		if err != nil {
			return nil, err
		}

		return unmarshalLock(data)
	}

	// We didn't find a lock file
	return nil, nil
}

// unmarshalLock - Load a Lock from a byte array into a Lock struct
func unmarshalLock(lockYaml []byte) (*Lock, error) {
	lock := &Lock{}

	err := yaml.Unmarshal([]byte(lockYaml), lock)
	return lock, err
}
