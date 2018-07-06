+++
title = "kombustion.yaml"
description = "Manifest specification"
date = "2018-06-14T00:00:00+10:00"
weight = 20
draft = false
bref = ""
toc = true
layout = "docs"
+++


## `kombustion.yaml`

You can generate this with `kombustion init`.

```yaml
Name: Test
Region: ap-southeast-2
Plugins:
  github.com/KablamoOSS/kombustion-plugin-serverless@0.1.0:
    Name: github.com/KablamoOSS/kombustion-plugin-serverless
    Version: 0.1.0
    Alias: ""
Environments:
  Production:
    AccountIDs:
    - "13521354"
    Parameters:
      ENVIRONMENT: production
```

### `Name`

_The name of your project._

If `--stack-name` is not provided this is used along with `--environment` and the file name, to make the CloudFormation Stack name.

### `Region`

_Default region to use._

Can be overidden by passing `--region us-east-1` with your desired region.

### `Plugins`

_A list of all plugins_

A plugin has a key formed of it's `Name` and `Version`, under which contains the `Name`, `Version`, and optional `Alias`.

#### `Name`

_URL of the plugin._

Currently only Github is supported.

#### `Version`

_Version constraint._

Using [SemVer](https://semver.org) to describe which version of the plugin you need. This is pinned in `kombustion.lock`.

#### `Alias` _Optional_

_Add an alias to the plugin._

If two plugins use the same namespace for their resource, you can add an `Alias` to one of them to use both.

### `Environments`

_Allows you to provide Parameters to your Stacks based on the target environment_

#### `AccountIDs` _Optional_

_A whitelist of Account ID's this environment can be deployed to._

If the Account ID the stack is going to be deployed into does not match from this list, the operation will fail. This
is a safety to prevent accidentally deploying into the wrong account.

#### `Parameters`

_A map of `Key: Value` parameters that will be supplied to the CloudFormation Stack._

Only the Paramters the stack needs are supplied, so multiple stacks can all use a subset of all the Parameters.
