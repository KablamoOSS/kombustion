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

### func RegisterParser

```go
func RegisterParser(
    parser func(
      name string,
      data string,
    ) (
      conditions TemplateObject,
      metadata TemplateObject,
      mappings TemplateObject,
      outputs TemplateObject,
      parameters TemplateObject,
      resources TemplateObject,
      errors []error,
    )
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
