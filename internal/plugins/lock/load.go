package lock

import (
	"fmt"
	// "io/ioutil"
	// "os"

	// printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/core"
	"github.com/KablamoOSS/yaml"
)

func GetLockObject(objectStore core.ObjectStore, path string) (*Lock, error) {
	data, err := objectStore.Get(path)
	if err != nil {
		return nil, fmt.Errorf("get lock: %v", err)
	}

	return unmarshalLock(data)
}

// unmarshalLock - Load a Lock from a byte array into a Lock struct
func unmarshalLock(lockYaml []byte) (*Lock, error) {
	lock := &Lock{}

	err := yaml.Unmarshal([]byte(lockYaml), lock)
	return lock, err
}
