* Start Date: (2018-05-22)
* RFC PR:
* Kombustion Issue:

# Summary

> Brief explanation of the feature.

This is the initial proposal to manage plugins in Kombustion. It covers
retreiving plugins, standard locations and filenames for plugins, and a manifest
and lockfile to ensure consistency.

# Basic example

> If the proposal involves a new or changed API, include a basic code example.

## Cli commands

```sh
# Download the plugin at github.com/KablmoOSS/kombustion-bation-host
kombustion add github.com/KablmoOSS/kombustion-bation-host

# Install all plugins in the manifest file
kombustion install
```

## Files

`kombustion.yaml` - **Manifest**

The manifest file needs to store package information to allow for versioning and
retrival.

```yaml
plugins:
  - github.com/KablmoOSS/kombustion-bation-host@=1.x.x
  - github.com/KablmoOSS/kombustion-example-plugin@latest
  - github.com/KablmoOSS/kombustion-another-example-plugin@~1.x.x
architectures: [ "x68", "arm"]
```

`kombustion.lock` - **Lock file**

The lock file is responsible for saving the resolved dependencies and their
versions. It should contain the absolute urls the plugins were downloaded from,
and the actual version that was downloaded.

```yaml
# This file is auto-generated, and it's structure is yet to be defined
```

`.kombustibles` - **Local plugins folder**

A directory located as a sibiling to the `kombustion.yaml` manifest file. This
directory contains the downloaded plugins.

The name here is unresolved, but considerations need to be made regarding other
tools that may be used, and how to avoid colliding with them.

# Motivation

> Why are we doing this? What use cases does it support? What is the expected
> outcome?
>
> Please focus on explaining the motivation so that if this RFC is not accepted,
> the motivation could be used to develop alternative solutions. In other words,
> enumerate the constraints you are trying to solve without coupling them too
> closely to the solution you have in mind.

Kombustion's power comes from it's plugin system, and therefore it's stability
and reliability is dependent on the management of those plugins.

This RFC proposes an initial plugin management solution that allows for
expansion to include a registry and potentially plugins that depend upon other
plugins. In the intial RFC (this one), plugin dependecies and therefore the need
to resolve a dependency graph are considered out of scope.

The goal is to provide a developer, CI, and Docker friendly CLI API to manage
plugins for a repository of stacks.

If Kombustion gains widespread use, and a large catalogue of plugins, there will
be benefit to the community to having a centralised reliable registry. But until
then, it's not desirable to over-invest and over-engineer this aspect of
Kombustion to the detriment of others.

# Detailed design

> This is the bulk of the RFC. Explain the design in enough detail for somebody
> familiar with Kombustion to understand, and for somebody familiar with the
> implementation to implement. This should get into specifics and corner-cases,
> and include examples of how the feature is used. Any new terminology should be
> defined here.

## Cli Commands

`kombustion init`

An interactive CLI command, that initalises the `kombustion.yaml` file. It
should prompt for which architectures to add, and default to the current system
architecture.

`kombustion add {repository}/{respository-name}@{version-condition}{tag}`

The command to retrieve and add a new plugin.

It must add a new plugin to the manifest file, and then call
`kombustion install`

`kombustion install`

It must load `kombustion.lock` into memory.

It must check the GitHub api for information on the repository tag, and retrieve
a list of releases.

It must find the release associated with the tag (that matches the
`version-condition`), and add that to `kombustion.lock`.

It must then run `kombustion install` to install any new plugins added.

It must store these plugins in the `Plugin Directory`.

It must then save the new manifest and lock files.

### Plugin Directory

The layout of the plugin directory will depend on the best method to reconcile
the lock file, and the installed plugins.

**Local plugin directory** The current proposal is to put local packages into
`.kombustibles` as a sibiling directory to the manifest and lock files.

**Global plugin directory** The current proposal is to maintain the
`~/.kombustion` that is `.kombustion` in the `$HOME` directory of the current
user.

### Files

`kombustion.yaml` - **Manifest**

The manifest file needs to store package information to allow for versioning and
retrival.

_OPTIONAL_:The manifest file could be extended to enhance the usability of
`Kombustion`

```yaml
# A list of plugins used in this project
# of the format
#  github.com/{organisation/user}/{repository-name}@{version-condition}{tag}
#
# Where version-condition is one of:
# = : equal to
# > : less than
# < : greater than
# ~ : Anything in the major version

plugins:
  - github.com/KablmoOSS/kombustion-bation-host@=1.x.x
  - github.com/KablmoOSS/kombustion-example-plugin@latest
  - github.com/KablmoOSS/kombustion-another-example-plugin@~1.x.x

# This determines the plugin compilations to download
# leave blank for all
architectures: [ "x68", "arm"]

# OPTIONAL ADDITIONS
# The ideas are optional and not necessarily in the scope of this RFC, but included for discussion
# regarding extra functionality of the manifest file. If any of these are to be considered, they
# may need their own RFC.

# Provide the only accounts stacks in this project will upload to.
# This can help when working with many AWS accounts, where you want to minimise the risk of
# uploading a stack to the wrong account.

# In the first example, provide a list of all allowed accountIds
accounts: [ "123456789", "456789123"]

# In the second example, provide a list of allowed accountId's and the enviroment they correlate to
# In this case the environment would need to be provided at runtime
# kombustion cf upsert stack.yml --env development
accounts:
  - enviroment: development
    accountId: "123456789"
  - environment: production
    accountId: "456789123"
```

`version-condition` is currently unresolved in so much as the best options to
support.

# Drawbacks

> Why should we _not_ do this? Please consider:
>
> * implementation cost, both in term of code size and complexity
> * whether the proposed feature can be implemented in user space
> * the impact on teaching people Kombustion
> * integration of this feature with other existing and planned features
> * cost of migrating existing Kombustion applications (is it a breaking
>   change?)
>
> There are tradeoffs to choosing any path. Attempt to identify them here.

This ads complexity to the tool, it could be argued that `Kombustion` is merely
a pre-processor and should not be concerned with managing plugins.

The decsions made here will cement a direction for plugin management, that may
cause pain to change from. This proposal attempts to minimise this.

# Alternatives

> What other designs have been considered? What is the impact of not doing this?

Plugins could be managed manually, or with a tool outside of Kombustion.

# Adoption strategy

> If we implement this proposal, how will existing Kombustion developers adopt
> it? Is this a breaking change? Can we write a codemod? Should we coordinate
> with other projects or libraries?

As the tool is not currently in widespread use, adoption is a matter of
education.

# How we teach this

> What names and terminology work best for these concepts and why? How is this
> idea best presented? As a continuation of existing Kombustion patterns?
>
> Would the acceptance of this proposal mean the Kombustion documentation must
> be re-organized or altered? Does it change how Kombustion is taught to new
> developers at any level?
>
> How should this feature be taught to existing Kombustion developers?

This should be easy to teach within the current Readme, documentation and within
the cli tool help.

# Unresolved questions

> Optional, but suggested for first drafts. What parts of the design are still
> TBD?

1.  The output of `kombustion.lock` is auto-generated and not yet resolved.
2.  Global plugins, local plugins, or both?
3.  Multiple OS version of plugins.
    * Should one or multple be downloaded?
4.  The name of the `Plugin Directory`.
5.  `version-condition` see `kombustion.yaml`
