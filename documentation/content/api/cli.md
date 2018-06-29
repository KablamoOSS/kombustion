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

### `verbose`

_Output with high verbosity._

```bash
$ kombustion --verbose
```

---

### `param`, `p`

_Specify Cloudformation parameters._

Parameters are also sourced from `kombustion.yaml`, but paramteres passed via the cli have precedence.
So anything you pass via this option, will be used instead of whats in `kombustion.yaml`

```bash
$ kombustion --param BucketName=test

# Or
$ kombustion -p BucketName=test
```

---

### `profile`

_Use a profile from ~/.aws/credentials_

```bash
$ kombustion --profile MyProfile
```

---

### `load-plugin`

_Load arbitrary plugin._

This option is only avaiable when Kombustion is built from source, see [wirting a plugin](/guides/plugins)
for more information.

```bash
$ kombustion --load-plugin path/to/plugin.so
```

---

### `help`, `h`

_Prints help._

```bash
$ kombustion --help, -h
```

---

### `version`, `v`

_Print the version._

```bash
$ kombustion --version
> kombustion version v1.0.0

# Or
$ kombustion -v
> kombustion version v1.0.0
```

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

The following commands manage Cloudformation Stacks.

---

### `generate`

_Generate a Cloudformation template, from a template file_

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

- Takes one positional argument, that is a relative path to the template file.

```bash
# Arguments
$ kombustion upsert [template file]

# Usage
$ kombustion upsert path/to/cloudformation/stack.yaml
```

__Errors__

If the stack is not created successfully for any reason, `kombustion` returns an exit code of `1` (an error).

__No updates to perform__

If there are no updates to perform, `kombustion` will return an exit code of `0` (no error).

---

### `delete`

_Delete a cloudformation stack._

- Takes one positional argument, that is a relative path to the template file.

```bash
# Arguments
$ kombustion delete [template file]

# Usage
$ kombustion delete path/to/cloudformation/stack.yaml
```

__Errors__

If the stack is not deleted for any reason, `kombustion` returns an exit code of `1` (an error).

---

### `events`

_Print all the events for a stack_

- Takes one positional argument, that is a relative path to the template file.

```bash
# Arguments
$ kombustion events [template file]

# Usage
$ kombustion events path/to/cloudformation/stack.yaml
```
