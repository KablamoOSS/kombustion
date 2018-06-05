package api

import (
	"log"

	pluginTypes "github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
	kombustionTypes "github.com/KablamoOSS/kombustion/types"

	"github.com/vmihailenco/msgpack"
)

func marshallConfig(config pluginTypes.Config) (blob []byte) {
	blob, err := msgpack.Marshal(&config)
	if err != nil {
		log.Fatal("Config marshalling err:", err)
	}
	return
}

func marshallResource(resource kombustionTypes.TemplateObject) (blob []byte) {
	blob, err := msgpack.Marshal(&resource)
	if err != nil {
		log.Fatal("Resource marshalling err:", err)
	}
	return
}

func marshallMapping(mapping kombustionTypes.TemplateObject) (blob []byte) {
	blob, err := msgpack.Marshal(&mapping)
	if err != nil {
		log.Fatal("Mapping marshalling err:", err)
	}
	return
}

func marshallOutput(output kombustionTypes.TemplateObject) (blob []byte) {
	blob, err := msgpack.Marshal(&output)
	if err != nil {
		log.Fatal("Output marshalling err:", err)
	}
	return
}
