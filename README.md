[![GoDoc](https://godoc.org/github.com/Codehardt/go-ahocorasick?status.svg)](https://godoc.org/github.com/Codehardt/go-ahocorasick)
[![Build Status](https://travis-ci.org/Codehardt/go-ahocorasick.svg?branch=master)](https://travis-ci.org/Codehardt/go-ahocorasick)
[![Go Report Card](https://goreportcard.com/badge/github.com/Codehardt/go-ahocorasick)](https://goreportcard.com/report/github.com/Codehardt/go-ahocorasick)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## go-ahocorasick
Aho-Corasick algorithm implemented in Golang

## Usage

```golang
ac := ahocorasick.New([]string{/*your strings here*/})
matches := ac.Match(/*your text here*/)
```
