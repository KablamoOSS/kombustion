package lock

import (
	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/config"
	"github.com/KablamoOSS/kombustion/internal/core"
	"github.com/KablamoOSS/yaml"
)

func (lockFile *Lock) Save(objectStore core.ObjectStore) error {
	// Marshal the the struct into yaml
	lockString, err := yaml.Marshal(&lockFile)
	if err != nil {
		return err
	}

	// Write the LockString
	err = objectStore.Put(lockString, "kombustion.lock")
	if err != nil {
		printer.Fatal(err, config.ErrorHelpInfo, "")
	}

	return nil
}
