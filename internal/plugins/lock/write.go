package lock

import (
	"io/ioutil"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/config"
	"github.com/KablamoOSS/yaml"
)

// WriteLockToDisk - Write the Lock to disk
func WriteLockToDisk(lockFile *Lock) error {

	// Mashall the the struct into yaml
	lockString, err := yaml.Marshal(&lockFile)
	if err != nil {
		return err
	}

	// Write the LockString
	err = ioutil.WriteFile("kombustion.lock", lockString, 0644)
	if err != nil {
		printer.Fatal(err, config.ErrorHelpInfo, "")
	}
	return nil
}
