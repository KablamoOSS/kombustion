+++
title = "Project"
description = "How to setup a project"
date = "2018-06-14T00:00:00+10:00"
weight = 20
draft = false
bref = "A project is a collection of template files, and plugins"
toc = true
layout  = "docs"
+++

## Best Practices

A project should be a version controlled folder, with `kombustion.yaml`, `kombustion.lock` and `.kombution` all committed to version control.

When using a plugin, a template cannot be created without the plugin - to ensure you can always update your stacks, you must commit your download plugins to version control.

## Create your project

In the root directory of your repository containing your CloudFormation templates run the following to create the `kombustion.yaml` file:

```bash
# Create a new kombustion.yaml file in the current directory
$ kombustion init
```
