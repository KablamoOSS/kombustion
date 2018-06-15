+++
title = "Writing a plugin"
description = ""
date = "2018-06-14T00:00:00+10:00"
weight = 20
draft = false
bref = "A plugin’s purpose is to transform input YAML into Cloudformation resources, outputs or mappings."
toc = true
layout  = "docs"
+++


## Getting Started

Before you start writing a plugin will need a working Go development environment, see Golang's [getting started](https://golang.org/doc/install).

In addition you need a version of `kombustion` built from source.

```bash
# Get the latest kmobustion
$ go get https://github.com/KablamoOSS/kombustion

# Compile and install it
$ go install github.com/KablamoOSS/kombustion

# Verify you have a version that's BUILT_FROM_SOURCE
$ kombustion -v
> kombustion version BUILT_FROM_SOURCE
```

When `kombustion` is built from source a flag is made available that allows loading any arbirary plugin for the specific purpose of developing plugins. This flag does not exit for official releases of `kombustion`, preventing arbritrary plugins from being loaded.

You can load a plugin using `--load-plugin path/to/plugin.so` before your command. For example, we can load a plugin built in a different folder.

```bash
$ kombustion --load-plugin ../kombustion-plugin-example/kombustion-plugin-example.so generate stacks/MyDemoStack.yaml
```


The easiest way to get started is to start from a copy of the [boilerplate](https://github.com/KablamoOSS/kombustion-plugin-boilerplate) example plugin. This repository has everything you need to get started writing a plugin, including default configuration, an example folder layout, and a build script.

```bash
$ go get https://github.com/KablamoOSS/kombustion-plugin-boilerplate
$ cp $GOPATH/src/github.com/KablamoOSS/kombustion-plugin-boilerplate $GOPATH/src/github.com/{username}/{plugin}
```


## Configuration

You plugin needs to provide some configuration information to `kombustion` when it's loaded. This includes information on the plugin itself, and the resources, mappings, and outputs the plugin provides.

> The return types in `plugin.go` are all `[]byte`, as communication between `kombustion` and plugins is sent as binary. This is a limitation/feature of Go's plugin implementation. Marshalling/unmarshalling of binary is taken care of for you.

This is defined in `plugin.go`. To register the plugin we need to provide a `Register` function that returns `api.RegisterPlugin` with our config.

Just as with `kombustion` version is passed in at compile time (this is done in the build script below), and we have a fallback version `BUILT_FROM_SOURCE`. When not using `--load-plugin` an error will be thrown if the version defined in the plugin does not match the version expected.

The other key configuration value is `Prefix`. This should be set to a pattern of `Organisation::Repository`, so for the boilerplate example we would use `Kablamo::Boilerplate`.

The prefix is then added to the start of any resources you define, seperated with `::`. So if we had a `Function` resource, the final resource export will be `Kablamo::Boilerplate::Function`

> There is no enforcement on what a `Prefix` can be, excepting those in Cloudformation: `AWS::*` and `Custom::*`. If a user has two plugins that have clashing prefixes, they can use the `Alias` paramter in `kombustion.yaml` to add another name in front, for example `MyAlias::Kablamo::Boilerplate::Function`.

```go
package main

import (
  // Import the plugin api functions
  "github.com/KablamoOSS/kombustion/pkg/plugins/api"
  // Import the plugin api types
  "github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
)

var (
  version string
)

func init() {
  // Set a fallback version when not built from a tag
  if version == "" {
    version = "BUILT_FROM_SOURCE"
  }
}

// Register plugin
func Register() []byte {
  return api.RegisterPlugin(types.Config{
    // Name should match the name of your repository
    // so github.com/KablamoOSS/kombustion-plugin-boilerplate
    // becomes kombustion-plugin-boilerplate
    Name:               "kombustion-plugin-boilerplate",
    // Version is set at compile time, to the tag
    Version:            version,
    // The prefix for all resources this plugin exports
    Prefix:             "Kablamo::Boilerplate",
    Help: types.Help{
      // Documentation coming
    },
  })
}

// You also need a main() function to be defined, but it will never be called
// This is required due to how Go's plugins work
func main() {}
```

Now our plugin is registered, we need to provide functions for our resources, mappings and outputs.

To register these you need to provide an exported variable `Resources`, `Mappings`, `Outputs`. They need to be defined exactly as follows, with your acutal Parser function wrapped in `api.RegisterResource`, `api.RegisterMapping` or `api.RegisterOutput`.

```go
// Resources for this plugin
var Resources = map[string]func(
  name string,
  data string,
) []byte{
  // resources.ParseLambdaFunction is explained in the next section
  "Function": api.RegisterResource(resources.ParseLambdaFunction),
}

// Mappings for this plugin
var Mappings = map[string]func(
  name string,
  data string,
) []byte{}

// Outputs for this plugin
var Outputs = map[string]func(
  name string,
  data string,
) []byte{}

```

## Writing a Parser function

A Parser function takes in YAML as a string, and returns one or more of its type (either `Resource`, `Mapping` or `Output`).

We generally store then in a seperate folder for each.

So to use a `Resource` function, as we did above, we need to import our `resources` package in `plugin.go`.

```go
package main

import (
  "github.com/KablamoOSS/kombustion/pkg/plugins/api"
  "github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
  // Import this plugins parser functions for resources
  "github.com/KablamoOSS/kombustion-plugin-boilerplate/resources"
)
```

The process for a Parser function is the same for `Resource`, `Mapping` and `Output`. So in this example we'll cover `Resource`.

In a new file at `resources/lambdaFunction.go` we start by importing the from `kombustion` core the Cloudformation parser functions, some core types, and a YAML library.

```go
package resources

import (
  cfResources "github.com/KablamoOSS/kombustion/pkg/parsers/resources"
  "github.com/KablamoOSS/kombustion/types"
  yaml "github.com/KablamoOSS/yaml"
)

```

First we need to make a `struct` defining the shape of the incoming YAML. This is what the user will put into their template.

In this example the YAML would look like:

```yaml
# Parameters are show as an example here
Parameters:
  LambdaBucket:
    Type: String
  LambdaRole:
    Type: String
Resources:
  # MyLambda is passed to the parser function as the `name` argument
  MyLambda:
    # Type is derived from Prefix::ResourceName
    Type: Kablamo::Boilerplate::Function
    # Properties is the Config struct we are about to define, which is passed to the
    # parser function as `data`
    Properties:
      Code: 
        Bucket: !Ref LambdaBucket
        Key:  "lambda/MyFunction.zip"
      Handler: "main"
      Role: !Ref LambdaRole
      Runtime: go1.x
```

The config `struct` uses [tags](https://medium.com/golangspec/tags-in-golang-3e5db0b8ef3e) to inform the YAML library how to unmarshall the yaml into the struct.

```go
// LambdaFunctionConfig defines the shape of the input YAML this resource takes
type LambdaFunctionConfig struct {
  Properties struct {
    Code           Code        `yaml:"Code"`
    Handler        interface{} `yaml:"Handler"`
    Role           interface{} `yaml:"Role"`
    Runtime        interface{} `yaml:"Runtime"`
  } `yaml:"Properties"`
}
type Code struct {
    Bucket           interface{} `yaml:"Bucket"`
    Key              interface{} `yaml:"Key"`
}
```

Now we have defined the shape of our YAML, we can write the parser function.

A parser function takes two arguments `name` and `data` both strings. Where `name` is the name of the object in the Cloudformation template, and `data` is the yaml of `Properties` of that object. See the yaml example above.

A parser function returns `types.TemplateObject`, and an `error`. When you return an `error`, it's treated as fatal.


```go
// ParseLambdaFunction converts our config into a cloudformation resource
func ParseLambdaFunction(
  name string,
  data string,
) (
  cf types.TemplateObject,
  err error,
) {
  // Setup a variable to load the yaml into
  var config LambdaFunctionConfig

  // Attempt to unmarshall the yaml
  if err = yaml.Unmarshal([]byte(data), &config); err != nil {
    // return an error if there was one
    return cf, err
  }

  // validate the config to ensure we have required fields
  // and apply any other validation logic
  // We'll cover this shortly.
  err = config.Validate()
  if err != nil {
    return cf, err
  }

  // Now we can create resources
  // To do this we need to call a create function from 
  // github.com/KablamoOSS/kombustion/pkg/parsers/resources
  // which we import as cfResource
  cf = types.TemplateObject{
    // We're going to reuse `name` as the name of the resource
    // This is merged in with the rest of the resources,
    // so if you want to make multiple resources, you should
    // derive their name starting from the name provided
    // to avoid clashing with any other resources (either from your own plugin, or another resource)
    (name): cfResources.NewLambdaFunction(
      cfResources.LambdaFunctionProperties{
        Code: properties.{
            S3Bucket:      config.Code.Bucket,
            S3Key:         config.Code.Key,
        }
        Handler: config.Code.Handler,
        Role: config.Code.Role,
        Runtime: config.Code.Runtime,
      }
    )
  }

  return cf, err
}

```


### Validation

In this example our validation function is only ensuring requried fields are provided. However, you can use any logic you want here to validate the input YAML matches what you require.

```go

// Validate is attached to LambdaFunctionConfig
func (config LambdaFunctionConfig) Validate() error {
  if config.Properties.Code == nil {
    return fmt.Errorf("Missing required field 'Code'")
  }
  if config.Properties.Handler == nil {
    return fmt.Errorf("Missing required field 'Handler'")
  }
  if config.Properties.Role == nil {
    return fmt.Errorf("Missing required field 'Role'")
  }
  if config.Properties.Runtime == nil {
    return fmt.Errorf("Missing required field 'Runtime'")
  }
}
```

## Build process


### Release


