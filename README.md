javapropio
==========

Write Java properties files in Go

[![Go Report Card](https://goreportcard.com/badge/go.gophers.dev/pkgs/javapropio)](https://goreportcard.com/report/go.gophers.dev/pkgs/javapropio)
[![Build Status](https://travis-ci.com/shoenig/regexplus.svg?branch=master)](https://travis-ci.com/shoenig/regexplus)
[![GoDoc](https://godoc.org/go.gophers.dev/pkgs/javapropio?status.svg)](https://godoc.org/go.gophers.dev/pkgs/javapropio)
[![NetflixOSS Lifecycle](https://img.shields.io/osslifecycle/shoenig/regexplus.svg)](OSSMETADATA)
[![GitHub](https://img.shields.io/github/license/shoenig/regexplus.svg)](LICENSE)

# Project Overview

Module `go.gophers.dev/pkgs/javapropio` provides a package for writing the content
of Java properties files.

# Getting Started

The `javapropio` package can be installed by running
```
$ go get go.gophers.dev/pkgs/javapropio
```

#### Example usage
```golang
var buf bytes.Buffer
w, _ := NewWriter(&buf)
_ = w.WriteProp("foo", "bar")
_ = w.Close()
s := buf.String() // java properties formatted
```

# Contributing

The `go.gophers.dev/pkgs/javapropio` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file an issue.

# License

The `go.gophers.dev/pkgs/javapropio` module is open source under the [BSD-3-Clause](LICENSE) license.

This is a fork of a deleted package, originally by github.com/rboyer