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

## Usage

```bash
$ kombustion [global options] command [command options] [arguments...]
```

All of these commands should be called from the same directory where `kombustion.yaml` is location,
which is usually the root directory of your project. If no `kombustion.yaml` can be found, an error is thrown.

## Global Options

### `--verbose`

_Output with high verbosity._

```bash
$ kombustion --verbose
```

---

### `--profile`

_Use a profile from ~/.aws/credentials_

```bash
$ kombustion --profile MyProfile
```

---

### `--load-plugin`

_Load arbitrary plugin._

This option is only avaiable when Kombustion is built from source, see [wirting a plugin](/guides/plugins)
for more information.

```bash
$ kombustion --load-plugin path/to/plugin.so
```

---

### `--help`, `-h`

_Prints help._

```bash
$ kombustion --help, -h
```

---

### `--version`, `-v`

_Print the version._

```bash
$ kombustion --version
> kombustion version v1.0.0

# Or
$ kombustion -v
> kombustion version v1.0.0
```

---

## Core

### `update`

_Update `kombustion` binary to the latest release._

```bash
$ kombustion update
```

__Options__

#### `-y`

_Don't prompt to update, if there is a release._

---

## Plugins

The following commands manage plugins in your project.

Learn more about how to [setup a project](/guides/project).

---

### `init`

_Initialise a new [manifest file](/guides/project) in the current directory._

```bash
$ kombustion init
```

---

### `add`

_Add a [plugin](/concepts/plugin) to your project._

- Takes one positional argument, that must be a Github repository url, with a release.

```bash
# Arguments
> kombustion add [url]

# Usage
$ kombustion add github.com/organisation/plugin
```

---

### `install`

_Install all plugins in kombustion.yaml._

```bash
$ kombustion install
```

---

## Stacks

The following commands manage CloudFormation Stacks.

With most of the following commands the Stack Name can either be supplied directly, or is derived
from `Title` in `kombustion.yaml`, the file name of the stack, and the environment in the
following pattern `{Title}-{FileName}-{Environment}`.

---

### `generate`

_Generate a CloudFormation template, from a template file_

Generate allows you to preview the final template, after plugins. It's the same output that is
generated when calling `upsert`.

- Takes one positional argument, that is a relative path to the template file.

```bash
# Arguments
$ kombustion generate [template file]

# Usage
$ kombustion generate path/to/cloudformation/stack.yaml
```

---

### `upsert`

_Update or insert a cloudformation template._

```bash
$ kombustion upsert [options] path/to/cloudformation/stack.yaml
```

__Arguments__

- Takes one positional argument, that is a relative path to the template file.

__Options__

#### `--region`, `-r`

_Set the region to deploy into._

#### `--stack-name`

_(Optional) Set the stack name._

If not provided a stack name is derived from `Title` in `kombustion.yaml`, the file name of the
stack, and the environment in the following pattern `{Title}-{FileName}-{Environment}`.

If you do not provide `--stack-name` you must provide `--environment`.

#### `--environment`, `-e`

_(Optional) The environment to deploy to._

Environment config to use as defined in ./kombustion.yaml.

If `AccountIDs` are listed under the environment, then you will only be able to deploy into that
account.

If `Parameters` are listed under the environment, they will be added to the stack.

_Required if `--stack-name` is not provided._

#### `--param`, `-p`

_(Optional) Specify CloudFormation parameters._

Parameters are also sourced from `kombustion.yaml`, but paramteres passed via the cli have precedence.
Anything you pass via this option, will be used instead of whats in `kombustion.yaml`.

```bash
$ kombustion upsert \
  --param ParamKeyOne=ParamValueOne, ParamKeyTwo=ParamValueTwo \
  path/to/cloudformation/stack.yaml
```

#### `--generate-default-outputs`, `-b`

_(Optional) Disable generation of outputs for Base AWS types_

#### `--iam`, `-i`

_(Optional) Shorthand for `--capability CAPABILITY_IAM`_

This gives the capability to perform upserts of IAM resources.

CloudFormation requires you to [acknowledge](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities) when you are modifying a stack with IAM resources.

#### `--capability`

_(Optional) Set capabilities for the upsert, for example `CAPABILITY_IAM`_


#### Errors

If the stack is not created successfully for any reason, `kombustion` returns an [exit code](#exit-codes) of `1` (an error).

__No updates to perform__

If there are no updates to perform, `kombustion` will return an [exit code](#exit-codes) of `0` (no error).

---

### `delete`

_Delete a cloudformation stack._

```bash
$ kombustion delete [options] path/to/cloudformation/stack.yaml
```

__Arguments__

- Takes one positional argument, that is a relative path to the template file.

__Options__

#### `--region`, `-r`

_Set the region to deploy into._

#### `--stack-name`

_(Optional) Set the stack name._

If not provided a stack name is derived from `Title` in `kombustion.yaml`, the file name of the
stack, and the environment in the following pattern `{Title}-{FileName}-{Environment}`.

If you do not provide `--stack-name` you must provide `--environment`.

#### `--environment`, `-e`

_(Optional) The environment to deploy to._

Environment config to use as defined in ./kombustion.yaml.

If `AccountIDs` are listed under the environment, then you will only be able to deploy into that
account.

If `Parameters` are listed under the environment, they will be added to the stack.

_Required if `--stack-name` is not provided._


#### Errors

If the stack is not deleted for any reason, `kombustion` returns an [exit code](#exit-codes) of `1` (an error).

---

### `events`

_Print all the events for a stack_


```bash
$ kombustion events [options] path/to/cloudformation/stack.yaml
```

__Arguments__

- Takes one positional argument, that is a relative path to the template file.

__Options__

#### `--region`, `-r`

_Set the region to deploy into._

#### `--stack-name`

_(Optional) Set the stack name._

If not provided a stack name is derived from `Title` in `kombustion.yaml`, the file name of the
stack, and the environment in the following pattern `{Title}-{FileName}-{Environment}`.

If you do not provide `--stack-name` you must provide `--environment`.

#### `--environment`, `-e`

_(Optional) The environment to deploy to._

Environment config to use as defined in ./kombustion.yaml.

If `AccountIDs` are listed under the environment, then you will only be able to deploy into that
account.

If `Parameters` are listed under the environment, they will be added to the stack.

_Required if `--stack-name` is not provided._

---

## Exit Codes

The matrix below describes the exit codes for each CloudFormation status.

In general when calling `upsert` if the changes requested (be they Create Stack, or Update Stack) are
not cleanly applied, an error is returned.

And when calling `delete` if the stack is not fully deleted, and error is returned.

__Legend:__

- `0` - no error
- `1` - error
- `~` - Transitional status, `kombustion` will not exit yet


| Status                                         | **Create Stack** | **Update Stack** | **Delete Stack** |
| ---------------------------------------------- | ---------------- | ---------------- | ---------------- |
| `CREATE_COMPLETE`                              | 0                | 0                | __1__            |
| `CREATE_IN_PROGRESS`                           | ~                | ~                | ~                |
| `CREATE_FAILED`                                | __1__            | __1__            | __1__            |
| `DELETE_COMPLETE`                              | __1__            | __1__            | 0                |
| `DELETE_FAILED`                                | __1__            | __1__            | __1__            |
| `DELETE_IN_PROGRESS`                           | ~                | ~                | ~                |
| `REVIEW_IN_PROGRESS`                           | __1__            | __1__            | __1__            |
| `ROLLBACK_COMPLETE`                            | __1__            | __1__            | __1__            |
| `ROLLBACK_FAILED`                              | __1__            | __1__            | __1__            |
| `ROLLBACK_IN_PROGRESS`                         | ~                | ~                | ~                |
| `UPDATE_COMPLETE`                              | 0                | 0                | __1__            |
| `UPDATE_COMPLETE_CLEANUP_IN_PROGRESS`          | ~                | ~                | ~                |
| `UPDATE_IN_PROGRESS`                           | ~                | ~                | ~                |
| `UPDATE_ROLLBACK_COMPLETE`                     | __1__            | __1__            | __1__            |
| `UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS` | ~                | ~                | ~                |
| `UPDATE_ROLLBACK_FAILED`                       | __1__            | __1__            | __1__            |
| `UPDATE_ROLLBACK_IN_PROGRESS`                  | ~                | ~                | ~                |
