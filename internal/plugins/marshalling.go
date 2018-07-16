package plugins

import (
	"fmt"
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

func loadResource(blob []byte) (resource kombustionTypes.TemplateObject, errs []error) {
	var pluginResult pluginTypes.PluginResult
	err := msgpack.Unmarshal(blob, &pluginResult)
	if err != nil {
		panic(err)
	}

	if pluginResult.Data != nil {
		resource = pluginResult.Data
	}
	if pluginResult.Errors != nil {
		errs = pluginResult.Errors
	}
	return
}

func loadMapping(blob []byte) (mapping kombustionTypes.TemplateObject, errs []error) {
	var pluginResult pluginTypes.PluginResult

	err := msgpack.Unmarshal(blob, &pluginResult)
	if err != nil {
		panic(err)
	}
	fmt.Println("pluginResult.Data")
	fmt.Println(pluginResult.Data)
	if pluginResult.Data != nil {
		mapping = pluginResult.Data
	}
	if pluginResult.Errors != nil {
		errs = pluginResult.Errors
	}

	return
}

func loadOutput(blob []byte) (output kombustionTypes.TemplateObject, errs []error) {
	var pluginResult pluginTypes.PluginResult

	err := msgpack.Unmarshal(blob, &pluginResult)
	if err != nil {
		panic(err)
	}

	if pluginResult.Data != nil {
		output = pluginResult.Data
	}
	if pluginResult.Errors != nil {
		errs = pluginResult.Errors
	}

	return
}
