package plugins

import (
	pluginTypes "github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
	kombustionTypes "github.com/KablamoOSS/kombustion/types"

	"github.com/vmihailenco/msgpack"
)

func loadConfig(blob []byte) (config pluginTypes.Config, err error) {
	err = msgpack.Unmarshal(blob, &config)
	if err != nil {
		panic(err)
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

func loadMapping(blob []byte) (mapping kombustionTypes.TemplateObject, err error) {
	err = msgpack.Unmarshal(blob, &mapping)
	if err != nil {
		panic(err)
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
