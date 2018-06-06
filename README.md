# Gerygone

Quick and Easy Cloud infrastructure validation built with [Go][].

[Overview](#overview) |
[Installation](#installation)

[![GoDoc](https://godoc.org/github.com/torrentalle/gerygone?status.svg)](https://godoc.org/github.com/torrentalle/gerygone)
[![Linux and macOS Build Status](https://api.travis-ci.org/torrentalle/gerygone.svg?branch=master&label=Linux+and+macOS+build "Linux and macOS Build Status")](https://travis-ci.org/torrentalle/gerygone)
[![Windows build status](https://ci.appveyor.com/api/projects/status/kn26cekg25iy5xmb/branch/master?svg=true)](https://ci.appveyor.com/project/torrentalle/gerygone/branch/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/torrentalle/gerygone)](https://goreportcard.com/report/github.com/torrentalle/gerygone)

## Overview

With Gerygone, you can write YAML tests for checking your cloud infrastructure is configured correctly.
It's planned to provide support for multiple cloud providers but in the initial stage only [AWS](https://aws.amazon.com) is supported.

Gerygone tests your actual infrastructure state by executing API calls. So you can use any Infrastructure as Code tool, Terraform, CloudFormation, Heat, Azure Templates and so on.

This package is **at an early stage of development and is not ready for production usage**



## Installation

If you want to use Gerygone to validate your cloud infrastructure, simply install the Gerygone binaries.
The Gerygone binaries have no external dependencies.

Finally, you can install the Gerygone source code with `go`, build the binaries yourself, and run Gerygone that way.
Building the binaries is an easy task for an experienced `go` getter.

### Install Gerygone (Binary Install)

**NOT YET SUPPORTED**

### Build and Install the Binaries from Source (Advanced Install)

#### Prerequisite Tools

* [Git](http://git-scm.com/)
* [Go (latest or previous version)](https://golang.org/dl/)

#### Vendored Dependencies

Hugo uses [dep](https://github.com/golang/dep) to vendor dependencies, but we don't commit the vendored packages themselves to the Hugo git repository. Therefore, a simple `go get` is _not_ supported because the command is not vendor aware.

The simplest way is to use [mage](https://github.com/magefile/mage) (a Make alternative for Go projects.)

#### Fetch from GitHub

```bash
go get github.com/magefile/mage
git clone https://github.com/torrentalle/gerygone --branch v0 \
    ${GOPATH}/src/github.com/torrentalle/gerygone 
cd ${GOPATH}/src/github.com/torrentalle/gerygone
mage vendor
mage install
```

**If you are a Windows user, substitute the `$HOME` environment variable above with `%USERPROFILE%`.**


[Go]: https://golang.org/