# lightenv [![Go Report Card](https://goreportcard.com/badge/github.com/d-tsuji/lightenv)](https://goreportcard.com/report/github.com/d-tsuji/lightenv) ![License MIT](https://img.shields.io/badge/license-MIT-blue.svg) [![GoDoc](https://godoc.org/github.com/d-tsuji/lightenv?status.svg)](https://godoc.org/github.com/d-tsuji/lightenv)

The lightenv is a lightweight library that handles environment variables. The thin wrapper using the standard [os](https://golang.org/pkg/os) library. This library is inspired by [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig).

The purpose is to allow the user to handle environment variables properly by returning to the caller as an error when the environment variable is not set. The feature is that: 

- You can set whether it is required as a struct tag like `required:"true"`.
- You can set a default value when no environment variables are set.

## USAGE
