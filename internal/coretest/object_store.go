package coretest

import (
	"io/ioutil"
	"strings"
)

// Implements the ObjectStore interface from internal/core, so that other
// packages can import it for testing.

type MockObjectStore struct {
	Objects map[string][]byte
}

func NewMockObjectStore() *MockObjectStore {
	return &MockObjectStore{
		Objects: make(map[string][]byte),
	}
}

// Get an object
func (objstore *MockObjectStore) Get(path string, subpath ...string) ([]byte, error) {
	fullPath := strings.Join(append([]string{path}, subpath...), "/")
	data, ok := objstore.Objects[fullPath]
	if !ok {
		return nil, nil
	}
	return data, nil
}

// Put a string into the objectStore
func (objstore *MockObjectStore) Put(data []byte, path string, subpath ...string) error {
	fullPath := strings.Join(append([]string{path}, subpath...), "/")
	objstore.Objects[fullPath] = data
	return nil
}

// PutFile into the objectStore by reading `localPath` from disk
func (objstore *MockObjectStore) PutFile(localPath string, path string, subpath ...string) error {
	data, err := ioutil.ReadFile(localPath)
	if err != nil {
		return err
	}
	fullPath := strings.Join(append([]string{path}, subpath...), "/")
	objstore.Objects[fullPath] = data
	return nil
}
