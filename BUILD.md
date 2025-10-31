### BUILD.md Content

```markdown
# Build and Distribution Guide

This guide details how to compile the `main.go` file into a standalone executable for distribution. Once compiled, end-users will still need Graphviz installed. 

## 1. Local Build

To create the executable for your current operating system:

```bash
# This creates a binary named 'main' (or 'main.exe' on Windows) in the current directory.
go build main.go

2. Cross-Platform Builds

Go makes it easy to compile for different operating systems and architectures. Use environment variables to target specific platforms before running go build.

Build for Windows (x64)

The resulting file will be state-machine-viz.exe.
Bash

GOOS=windows GOARCH=amd64 go build -o state-machine-viz.exe main.go

Build for macOS (Intel x64)

The resulting file will be state-machine-viz.
Bash

GOOS=darwin GOARCH=amd64 go build -o state-machine-viz main.go

Build for Linux (x64)

The resulting file will be state-machine-viz.
Bash

GOOS=linux GOARCH=amd64 go build -o state-machine-viz main.go

