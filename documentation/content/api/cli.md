+++
title = "Kombustion CLI"
description = "kombustion API"
date = "2018-06-14T00:00:00+10:00"
weight = 20
draft = false
bref = ""
toc = true
layout  = "docs"
+++

## Basic Usage

Generate a CloudFormation template (from `./examples/stacks/test.yaml`):

```bash
$ kombustion generate examples/stacks/test.yaml && cat compiled/test.yaml
```

Upsert a CloudFormation template:

```bash
$ kombustion upsert examples/stacks/test.yaml --stackName test-stack
```

Delete a CloudFormation stack:

```bash
$ kombustion delete examples/stacks/test.yaml
```

Print all the events for a stack:

```bash
$ kombustion events examples/stacks/test.yaml
```

## Command

```
kombustion [global options] command [command options] [arguments...]

COMMANDS:
     init      init manifest file
     add       add github.com/organisation/plugin
     install   install all plugins in kombustion.yaml
     generate  parse a cloudformation template from ./config
     upsert    upsert a cloudformation template or a yaml config
     delete    delete a cloudformation stack
     events    print all events for a cloudformation stack
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --verbose                                    output with high verbosity
   --param BucketName=test, -p BucketName=test  cloudformation parameters BucketName=test
   --profile MyProfile                          use a profile from ~/.aws/credentials eg MyProfile
   --load-plugin path/to/plugin.so              load arbitrary plugin path/to/plugin.so
   --help, -h                                   show help
   --version, -v                                print the version
```
