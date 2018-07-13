+++
title = "Plugins"
description = "How plugins work"
date = "2018-06-14T00:00:00+10:00"
weight = 20
draft = false
bref = ""
toc = true
layout  = "docs"
+++

## How a plugin works

Kombustion's implementation of plugins builds upon the Go [Plugin package][1].

Plugins are loaded according to the `kombustion.lock` file (which is generated from `kombustion.yaml`).

If a template uses a resource registered to a plugin, the relevant parser function is called. The
plugins parser function receives the name of the resource, and the template object as a string. It's
then up to the parser function to transform those inputs into one or more CloudFormation template objects, and
return them.

Due to the nature of Go's Plugin package, however it's not quite that simple. We need to serialise
(marshall) the input into binary before sending it to the plugin, and unserialise (unmarshall) it
on the other side. We also to the same to return the output back to `kombustion` from the plugin.

The reason for this is by design:
> A plugin cannot access the symbols of a host program.

What this means, is a plugin and it's host program must function mostly as two seperate programs. The
only information that can be transimitted between them, are types from the standard library (byte, string, func, bool).

If we try to pass a struct from the plugin to the host, the program will crash.

To get around this, `kombustion` uses an interface that abstracts all the binary marshalling away,
removing the need to write this boilerplate code in your plugin. When you call `api.RegisterResource` you're
calling a helper function, that returns another function that handles the binary marshalling.


### CloudFormation Definitions

Because plugins are compiled the version of CloudFormation definitions is frozen in time for that plugin.
This means, if a new parameter is added to a resource the plugin uses, it will be inaccessible until
the plugin is updated.

This issue is localised to the plugin, and plugins can operate with different definitions of CloudFormation
resources. This will only become problematic if CloudFormation has a breaking change, which happen rarely.

## Best Practices

To build a reliable, safe and secure plugin we have a few best practices.

### Plugins Should Be Pure

A plugin's parser functions should all be pure; they should generate the same output for the same input, and produce no side effects.

### Accessing AWS Resources in a plugin

If your plugin needs to access AWS resources during create or update, the only method you should
use are [Custom Resources][2].

An example of when you may want to do this, is to get the latest AMI for a set of tags. AWS has a [walkthrough for this use case][3].

[1]:	https://golang.org/pkg/plugin/
[2]:	https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources.html
[3]:	https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/walkthrough-custom-resources-lambda-lookup-amiids.html


## Cache

Kombustion caches downloaded plugins in `~/.kombustion/cache/plugins`.
