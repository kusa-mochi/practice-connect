#!/bin/bash

rm -r gen/
rm go.mod
rm go.sum
go mod init proto_example
go mod tidy
buf generate
