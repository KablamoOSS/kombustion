+++
title = "Writing a plugin"
description = "How to write a plugin"
date = "2018-06-14T00:00:00+10:00"
weight = 20
draft = false
bref = "A plugin’s purpose is to transform input YAML into CloudFormation resources, outputs or mappings."
toc = true
layout  = "docs"
+++


## Getting Started

Before you start writing a plugin will need a working Go development environment, see Golang's [getting started](https://golang.org/doc/install).

In addition you need a version of `kombustion` built from source.

```bash
# Get the latest kombustion
$ go get github.com/KablamoOSS/kombustion

# Compile and install it
$ go install github.com/KablamoOSS/kombustion

# Verify you have a version that's BUILT_FROM_SOURCE
$ kombustion -v
> kombustion version BUILT_FROM_SOURCE
```

When `kombustion` is built from source a flag is made available that allows loading any arbirary plugin for the specific purpose of developing plugins. This flag does not exist for official releases of `kombustion`, preventing arbritrary plugins from being loaded.

You can load a plugin using `--load-plugin path/to/plugin.so` before your command. For example, we can load a plugin built in a different folder.

```bash
$ kombustion --load-plugin ../kombustion-plugin-example/kombustion-plugin-example.so generate stacks/MyDemoStack.yaml
```


The easiest way to get started is to start from a copy of the [boilerplate](https://github.com/KablamoOSS/kombustion-plugin-boilerplate) example plugin. This repository has everything you need to get started writing a plugin, including default configuration, an example folder layout, and a build script.

```bash
$ go get github.com/KablamoOSS/kombustion-plugin-boilerplate
$ cp $GOPATH/src/github.com/KablamoOSS/kombustion-plugin-boilerplate \
  $GOPATH/src/github.com/{username}/{plugin}
```


## Configuration

You plugin needs to provide some configuration information to `kombustion` when it's loaded. This includes information on the plugin itself, and the resources, mappings, and outputs the plugin provides.

> The return types in `plugin.go` are all `[]byte`, as communication between `kombustion` and plugins is sent as binary. This is a limitation/feature of Go's plugin implementation. Marshalling/unmarshalling of binary is taken care of for you.

This is defined in `plugin.go`. To register the plugin we need to provide a `Register` function that returns `api.RegisterPlugin` with our config.

Just as with `kombustion` version is passed in at compile time (this is done in the build script below), and we have a fallback version `BUILT_FROM_SOURCE`. When not using `--load-plugin` an error will be thrown if the version defined in the plugin does not match the version expected.

The other key configuration value is `Prefix`. This should be set to a pattern of `Organisation::Repository`, so for the boilerplate example we would use `Kablamo::Boilerplate`.

The prefix is then added to the start of any resources you define, seperated with `::`. So if we had a `Function` resource, the final resource export will be `Kablamo::Boilerplate::Function`

> There is no enforcement on what a `Prefix` can be, excepting those in CloudFormation: `AWS::*` and `Custom::*`. If a user has two plugins that have clashing prefixes, they can use the `Alias` paramter in `kombustion.yaml` to add another name in front for example `MyAlias::Kablamo::Boilerplate::Function`

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
    Name: "kombustion-plugin-boilerplate",
    // Version is set at compile time, to the tag
    Version: version,
    // The prefix for all resources this plugin exports
    Prefix: "Kablamo::Boilerplate",
    Help: types.Help{
      Description: "An Example Plugin",
      TypeMappings: []types.TypeMapping{
        {
          Name:        "Function",
          Description: "Creates a function.",
          Config:      resources.LambdaFunctionConfig{},
        },
      },
    },
  })
}

// You also need a main() function to be defined, but it will never be called
// This is required due to how Go's plugins work
func main() {}
```

Now our plugin is registered, we need to provide functions for our resources, mappings and outputs.

To register these you need to provide an exported variable `Resources`, `Mappings`, `Outputs`. They need to be defined exactly as follows, with your actual Parser function wrapped in `api.RegisterResource`, `api.RegisterMapping` or `api.RegisterOutput`.

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

> A Parser function takes in YAML as a string, and returns one or more of its type (either `Resource`, `Mapping` or `Output`).

We generally store them in a separate folder for each.

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

In a new file at `resources/lambdaFunction.go` we start by importing them from `kombustion` core the CloudFormation parser functions, some core types, and a YAML library.

```go
package resources

import (
  cfResources "github.com/KablamoOSS/kombustion/pkg/parsers/resources"
  kombustionTypes "github.com/KablamoOSS/kombustion/types"
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

A parser function takes two arguments `name` and `data` both strings. Where `name` is the name of the object in the CloudFormation template, and `data` is the yaml of `Properties` of that object. See the yaml example above.

A parser function returns `types.TemplateObject`, and an `error`. When you return an `error`, it's treated as fatal.


```go
// ParseLambdaFunction converts our config into a cloudformation resource
func ParseLambdaFunction(
  name string,
  data string,
) (
  cf types.TemplateObject,
  errs []error,
) {
  // Setup a variable to load the yaml into
  var config LambdaFunctionConfig

  // Attempt to unmarshall the yaml
  err := yaml.Unmarshal([]byte(data), &config)

	if err != nil {
    // Append the error to the errs array
		errs = append(errs, err)
		return
	}


  // validate the config to ensure we have required fields
  // and apply any other validation logic
  // We'll cover this shortly.
  err = config.Validate()

	if err != nil {
    // Append the error to the errs array
		errs = append(errs, err)
		return
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

  return cf, errs
}

```

#### Returning Errors

You must return all errors in the `errs []error` array. These are then printed to the user, with
information about the plugin, and the block in the template that caused it.

**Don't** print errors to `stdout`, as the user won't know where they're coming from.

If `errs []error` contains any errors, the task will fail.

And example of how an error is printed:

```sh
✖  Error: Missing field 'CIDR'
☞  Resolution:
   ├─ Name:    MyNetwork
   ├─ Plugin:  kombustion-plugin-boilerplate
   └─ Type:    Kablamo::Example::VPC
```

### Validation

In this example our validation function is only ensuring requried fields are provided. However, you can use any logic you want here to validate the input YAML matches what you require.

```go

// Validate is attached to LambdaFunctionConfig
func (config LambdaFunctionConfig) Validate() (errors []error) {
  if config.Properties.Code == nil {
    errors = append(errors, fmt.Errorf("Missing required field 'Code'"))
  }
  if config.Properties.Handler == nil {
    errors = append(errors, fmt.Errorf("Missing required field 'Handler'"))
  }
  if config.Properties.Role == nil {
    errors = append(errors, fmt.Errorf("Missing required field 'Role'"))
  }
  if config.Properties.Runtime == nil {
    errors = append(errors, fmt.Errorf("Missing required field 'Runtime'"))
  }

  return
}
```

## Build process

When building locally all you need is to add the plugin buildmode to your normal build command:

```bash
$ go build --buildmode=plugin
```

This will create a `folderName.so` file, which you can then load with `--load-plugin path/to/folderName.so`.

Behind the scenes the plugin and core have been compiled with the C compiler on your machine. Becuase
these match it sould just work.

### Release

For release, you need to ensure you have a build for every operating system and architecture `kombustion` supports.

We use a tool called [`xgo`](github.com/karalabe/xgo) which uses Docker to compile for all possible versions, using the correct C compiler.

The following is a working `.travis.yml` configuration that will compile and attach a release to your Github repository.

`xgo` has a consistent naming convention, which `kombustion` relies on when downloading a plugin to determine what
operating system and architecture the plugin was compiled for.

To make this build script work, you need to add it to your root directory as `.travis.yml` in a public Github repository.

Then you need to setup a token to allow TravisCI to [publish a release](https://docs.travis-ci.com/user/deployment/releases/) on your behalf.

From the root directory of your repository run:

```bash
$ travis setup releases
```

It will then connect to your Github account, create and encrypt a token to allow publishing releases.

```yaml
language: go
os:
- linux
go:
- 1.10.1
sudo: required
script:
- go get -t ./...
- go generate
- go test ./...
- go get github.com/karalabe/xgo
- |
  # Get the full go repo url
  REPO=$(pwd |  rev | cut -d'/' -f-3 | rev)

  # Get the name of the app
  APP="${PWD##*/}"

  # Get this tag as the version
  VERSION=$(git describe --abbrev=0 --tags)

  # Ensure a fresh build folder
  rm -rf build && mkdir build
  # Compile
  xgo \
    -dest build/ \
    -buildmode=plugin  \
    --targets=darwin/amd64,freebsd/386,freebsd/amd64,freebsd/arm,linux/386,linux/amd64,linux/arm64  \
    --ldflags "-X plugin.version=${VERSION}" \
    $REPO

  # Package
  cd build
  # For each compiled binary, we're repackaging it in a zip with the architecture name, and
  # renaming the binary to the app name
  for FILE in $(ls .); do
    mv $FILE $APP.so
    tar cvzf ${FILE}.tgz $APP.so
    rm -f $APP.so
  done
  cd ..

# Deploy to Github release on tags
deploy:
  provider: releases
  api_key:
    secure: XXXXX...YourSecretHere
  file_glob: true
  file: "build/*"
  skip_cleanup: true
  on:
    tags: true
```


### Creating a release

Finally to create a release, you need to create a tag, and push both the commit and tag up.

```bash
# Create a tag
$ git tag v0.1.0

# Ensure the tagged commit has been pushed
$ git push

# Now push your tags
$ git push --tags
```



