+++
title = "Plugin API"
description = "Plugin API Specification"
date = "2018-06-14T00:00:00+10:00"
weight = 20
draft = false
bref = ""
toc = true
layout  = "docs"
+++


## `pkg/plugins/api`

### func RegisterMapping

```go
func RegisterMapping(
    mapping func(
        name string,
        data string,
    ) (cf types.TemplateObject),
) func(
    name string,
    data string,
) []byte
```

RegisterMapping for your plugin

### func RegisterOutput

```go
func RegisterOutput(
    output func(
        name string,
        data string,
    ) (cf types.TemplateObject),
) func(
    name string,
    data string,
) []byte
```

RegisterOutput for your plugin

### func RegisterPlugin

```go
func RegisterPlugin(config apiTypes.Config) []byte
```

RegisterPlugin to provide the name, prefix and version, and requiresAWSSession

### func RegisterResource

```go
func RegisterResource(
    resource func(
        name string,
        data string,
    ) (cf types.TemplateObject),
) func(
    name string,
    data string,
) []byte
```

RegisterResource for your plugin

__Usage__

```go
var Resources = map[string]func(
  name string,
  data string,
) []byte{
  // resources.ParseLambdaFunction is explained in the next section
  "Function": api.RegisterResource(resources.ParseLambdaFunction),
}
```
