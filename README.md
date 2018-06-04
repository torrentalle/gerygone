# Gerygone

Quick and Easy infrastructure validation built with [Go][].

[Overview](#overview) |
[Installation](#installation)

[![GoDoc](https://godoc.org/github.com/torrentalle/gerygone?status.svg)](https://godoc.org/github.com/torrentalle/gerygone)
[![Linux and macOS Build Status](https://api.travis-ci.org/torrentalle/gerygone.svg?branch=master&label=Linux+and+macOS+build "Linux and macOS Build Status")](https://travis-ci.org/torrentalle/gerygone)
[![Windows Build Status](https://ci.appveyor.com/api/projects/status/fdfdfdfdfd?svg=true&label=Windows+build "Windows Build Status")](https://ci.appveyor.com/project/bep/gerygone/branch/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/torrentalle/gerygone)](https://goreportcard.com/report/github.com/torrentalle/gerygone)

## Overview

Gerygone is a YAML based tool for validating a infrastructure’s configuration. 
It eases the process of writing tests by allowing the user to generate tests from the current system state. 
Once the test suite is written they can be executed, waited-on, or served as a health endpoint.

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
go get -d github.com/torrentalle/gerygone
cd ${GOPATH:-$HOME/go}/src/github.com/torrentalle/gerygone
mage vendor
mage install
```

**If you are a Windows user, substitute the `$HOME` environment variable above with `%USERPROFILE%`.**


[Go]: https://golang.org/