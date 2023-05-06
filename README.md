# UI

An immedia mode UI for Go.

[![Go Report Card](https://goreportcard.com/badge/github.com/emad-elsaid/ui)](https://goreportcard.com/report/github.com/emad-elsaid/ui)
[![GoDoc](https://godoc.org/github.com/emad-elsaid/ui?status.svg)](https://godoc.org/github.com/emad-elsaid/ui)



# What

- Set of widgets for Go
- Build on top of https://gioui.org
- Built to allow nesting function calls


# Why

- For learning
- For my Go projects
- For anyone who likes the idea (it's a bit similar to SwiftUI declarative syntax)


# When

- Now it's a work in progress

# How

## Examples

* The package provides runnable examples each is named after the widget function.
* Examples will show widgets in a window
* Clone the package or `go get` it to your system.
* to run an example use. where `RoundedCorners` is the function name and `count=1` is to force running the example regardless of caching.

```
go test . -run RoundedCorners -count=1
```
