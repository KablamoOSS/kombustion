package plugins

import (
	"log"

	pluginTypes "github.com/KablamoOSS/kombustion/plugins/api/types"
	kombustionTypes "github.com/KablamoOSS/kombustion/types"

	"github.com/vmihailenco/msgpack"
)

// [ Config ]---------------------------------------------------------------------------------------

// MarshallConfig for your plugin
func MarshallConfig(config pluginTypes.Config) (blob []byte) {
	blob, err := msgpack.Marshal(&config)
	if err != nil {
		log.Fatal("Config marshalling err:", err)
	}
	return
}

// loadHelp from a plugin
func loadConfig(blob []byte) (config pluginTypes.Config, err error) {
	err = msgpack.Unmarshal(blob, &config)
	if err != nil {
		panic(err)
	}
	return
}

// [ Resources ]------------------------------------------------------------------------------------

// MarshallResource for your plugin
func MarshallResource(resource kombustionTypes.TemplateObject) (blob []byte) {
	blob, err := msgpack.Marshal(&resource)
	if err != nil {
		log.Fatal("Resource marshalling err:", err)
	}
	return
}

func loadResource(blob []byte) (resource kombustionTypes.TemplateObject, err error) {
	err = msgpack.Unmarshal(blob, &resource)
	if err != nil {
		panic(err)
	}
	return
}

// [ Mapping ]--------------------------------------------------------------------------------------

// MarshallMapping for your plugin
func MarshallMapping(mapping kombustionTypes.TemplateObject) (blob []byte) {
	blob, err := msgpack.Marshal(&mapping)
	if err != nil {
		log.Fatal("Mapping marshalling err:", err)
	}
	return
}

func loadMapping(blob []byte) (mapping kombustionTypes.TemplateObject, err error) {
	err = msgpack.Unmarshal(blob, &mapping)
	if err != nil {
		panic(err)
	}
	return
}

// [ Outputs ]--------------------------------------------------------------------------------------

// MarshallOutput for your plugin
func MarshallOutput(output kombustionTypes.TemplateObject) (blob []byte) {
	blob, err := msgpack.Marshal(&output)
	if err != nil {
		log.Fatal("Output marshalling err:", err)
	}
	return
}

func loadOutput(blob []byte) (output kombustionTypes.TemplateObject, err error) {
	err = msgpack.Unmarshal(blob, &output)
	if err != nil {
		panic(err)
	}
	return
}
