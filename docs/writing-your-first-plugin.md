# Writing your first plugin

!> Kombustion plugins are not yet supported on Windows. Please use Docker or WSL in the meantime.

## Using plugins

Kombustion utilizes the `package` plugin (https://godoc.org/plugin). By default, Kombustion
will look for plugins in the `~/.kombustion/plugins` directory. You can also specify custom
plugins directory:

```sh
PLUGINS=/plugins kombustion cf generate stack.yaml
```

## Building from the plugin quickstart

The following documentation uses the plugin name `myplugin` as an example. Please ensure you have completed the [initialization](initialization.md) before starting.

```sh
cd $GOPATH/src/github.com/KablamoOSS/kombustion/plugins/
git clone https://github.com/KablamoOSS/kombustion-plugin-boilerplate
mv kombustion-plugin-boilerplate myplugin
cd myplugin
```

Open the file `plugin.go` in the tutorial directory and amend it to produce the following content:

```plugin.go
package main

import (
	"github.com/KablamoOSS/kombustion/plugins/myplugin/resources"
	"github.com/KablamoOSS/kombustion/types"
)

var Resources = map[string]types.ParserFunc{
	"Tutorial::Example::MultiBucket": resources.ParseExampleMultiBuckets,
}

var Outputs = map[string]types.ParserFunc{}

var Mappings = map[string]types.ParserFunc{}

var Help = types.PluginHelp{
	Description: "My tutorial plugin",
	TypeMappings: []types.TypeMapping{
		{
			Name:        "Tutorial::Example::MultiBucket",
			Description: "Creates a number of S3 buckets.",
			Config:      resources.MultiBucketConfig{},
		},
	},
}

func main() {}
```

We've now added a reference to the `Tutorial::Example::MultiBucket`, which we will now define. Open the `resources/examplemultibucket.go` file and amend it to match the following:

```resources/examplemultibucket.go
package resources

import (
	"log"
	"strconv"

	yaml "github.com/KablamoOSS/go-yaml"
	"github.com/KablamoOSS/kombustion/pluginParsers/resources"
	"github.com/KablamoOSS/kombustion/types"
)

type MultiBucketConfig struct {
	Properties struct {
		Count *int    `yaml:"Count"`
	} `yaml:"Properties"`
}

func ParseExampleMultiBuckets(name string, data string) (cf types.ValueMap, err error) {
	// Parse the config data
	var config MultiBucketConfig
	if err = yaml.Unmarshal([]byte(data), &config); err != nil {
		return
	}

	// validate the config
	config.Validate()

	// create a group of objects (each to be validated)
	cf = make(types.ValueMap)

	// defaults
	count := 1
	if config.Properties.Count != nil {
		count = *config.Properties.Count
	}

	// create the buckets
	for i := 1; i <= count; i++ {
		cf[name+"S3Bucket"+strconv.Itoa(i)] = resources.NewS3Bucket(
			resources.S3BucketProperties{
				BucketName: name+"-"+strconv.Itoa(i),
			},
		)
	}

	return
}

// Validate - input Config validation
func (this MultiBucketConfig) Validate() {
	if this.Properties.Count == nil {
		log.Println("WARNING: MultiBucketConfig - You did not specify a 'Count', defaulting to 1")
	}
}
```

We've just created a `MultiBucketConfig` type, which will be used to pass the values for the `Properties:` section in your template. The defined `Count` property defines how many S3 buckets will be output in the processed template. If that property is not defined, we output a warning and use the default value, 1.

## Compiling the plugin

Now we have the plugin fully defined, let's build it for your system. Execute the following:

```sh
cd $GOPATH/src/github.com/KablamoOSS/kombustion/plugins/myplugin
go build -buildmode plugin
```

You should now see that we have created the `myplugin.so` file in the directory. Kombustion loads this compiled plugin at runtime.

## Test your plugin

You'll now need to test your plugin. Let's define that in a file called `mystack.yaml`:

```mystack.yaml
AWSTemplateFormatVersion: 2010-09-09
Description: Testing my custom plugin
Resources:
  manyBuckets:
    Type: Tutorial::Example::MultiBucket
    Properties:
      Count: 3
```

Now let's generate the config:

```sh
PLUGINS=$GOPATH/src/github.com/KablamoOSS/kombustion/plugins/ kombustion cf generate mystack.yaml
cat compiled/mystack.yaml
```

Congratulations! You have generated the stack from your custom plugin.