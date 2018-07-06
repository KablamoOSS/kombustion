+++
title = "Project"
description = "A project contains your template files, and kombustion files."
date = "2018-06-14T00:00:00+10:00"
weight = 20
draft = false
bref = "A project is a collection of templates, along with the plugins required to use them."
toc = true
layout = "docs"
+++


The `kombustion.yaml` manifest file fits best alongside the concept of a project. Your project may
have a few CloudFormation template files in them. And those template files use resources provided
by Kombustion plugins. Therefore, without the plugins the template files cannot be used.

To ensure you can always modify your CloudFormation stacks, we reccomend committing `kombustion.yaml`,
`kombustion.lock` and the `.kombustion` folder. The latter stores downloaded plugins. With those pieces
you will always be able to `kombustion upsert` and `kombustion delete`. At any time you can generate
the final CloudFormation template with `kombustion generate`.
