package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type EFSFileSystem struct {
	Type       string                      `yaml:"Type"`
	Properties EFSFileSystemProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EFSFileSystemProperties struct {
	Encrypted interface{} `yaml:"Encrypted,omitempty"`
	KmsKeyId interface{} `yaml:"KmsKeyId,omitempty"`
	PerformanceMode interface{} `yaml:"PerformanceMode,omitempty"`
	FileSystemTags interface{} `yaml:"FileSystemTags,omitempty"`
}

func NewEFSFileSystem(properties EFSFileSystemProperties, deps ...interface{}) EFSFileSystem {
	return EFSFileSystem{
		Type:       "AWS::EFS::FileSystem",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEFSFileSystem(name string, data string) (cf types.ValueMap, err error) {
	var resource EFSFileSystem
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EFSFileSystem - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EFSFileSystem) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EFSFileSystemProperties) Validate() []error {
	errs := []error{}
	return errs
}
