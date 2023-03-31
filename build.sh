#!/bin/bash

# Build for the current platform (binary)
go build -o bin/

# Build for windows 64 bit (.exe)
env GOOS=windows GOARCH=amd64 go build -o bin/
