package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"plugin"
	"reflect"

	"github.com/KablamoOSS/kombustion/pkg/plugins/api/types"

	log "github.com/sirupsen/logrus"
)

type TypeDefinition struct {
	Name        string
	Description string
	Fields      []FieldDefinition
}

type FieldDefinition struct {
	Name         string
	Description  string
	ExampleValue string
	Children     []FieldDefinition
}

func getPluginDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	plugindir := fmt.Sprintf("%s/.kombustion/plugins", usr.HomeDir)
	os.MkdirAll(plugindir, 0744)

	return plugindir
}

func loadPlugins() []TypeDefinition {
	pluginBaseDir := ""
	processedTypes := []TypeDefinition{}

	if len(os.Getenv("PLUGINS")) > 0 {
		pluginBaseDir = os.Getenv("PLUGINS")
	} else {
		pluginBaseDir = getPluginDir()
	}

	pluginDirectories, err := ioutil.ReadDir(pluginBaseDir)
	if err != nil {
		// log.Println("WARNING: ", err)
		return processedTypes
	}

	for _, d := range pluginDirectories {
		if !d.IsDir() {
			// load all plugins in the base dir
			processedTypes = append(processedTypes, loadPlugin(d.Name(), pluginBaseDir)...)

			continue
		}

		pluginDir := filepath.Join(pluginBaseDir, d.Name())
		plugins, err := ioutil.ReadDir(pluginDir)
		if err != nil {
			continue
		}

		// load all plugins in each sub dir
		for _, f := range plugins {
			processedTypes = append(processedTypes, loadPlugin(f.Name(), pluginDir)...)
		}
	}

	return processedTypes
}

func loadPlugin(filename, pluginDir string) (ret []TypeDefinition) {
	if filepath.Ext(filename) != ".so" && filepath.Ext(filename) != ".dll" && filepath.Ext(filename) != ".dylib" {
		return
	}

	pluginPath := filepath.Join(pluginDir, filename)

	p, err := plugin.Open(pluginPath)

	if err != nil {
		log.WithFields(log.Fields{
			"filename": filename,
			"err":      err,
		}).Warn("error reading plugin file")
		return
	}

	help, _ := p.Lookup("Help")

	if help != nil {
		for _, typeMapping := range help.(*types.Help).Types {
			var typeDef TypeDefinition

			typeDef.Name = typeMapping.Name
			typeDef.Description = typeMapping.Description

			props := reflect.ValueOf(typeMapping.Config).FieldByName("Properties").Interface()
			for i := 0; i < reflect.TypeOf(props).NumField(); i++ {
				var fieldDef FieldDefinition

				field := reflect.TypeOf(props).Field(i)
				fieldDef.Name = field.Name

				exampleVal, hasExampleSet := field.Tag.Lookup("example")
				if !hasExampleSet {
					exampleVal = ""
				}
				fieldDef.ExampleValue = exampleVal

				typeDef.Fields = append(typeDef.Fields, fieldDef)
			}

			ret = append(ret, typeDef)
		}
	}

	return
}

func loadPrimaryResources() {
	/*for _, v := range parsers.GetParsers_resources() {
		log.Warn(reflect.TypeOf(v))
	}*/
}

func main() {
	res := loadPlugins()

	loadPrimaryResources()

	bytes, _ := json.Marshal(res)

	fmt.Println(string(bytes))
}
