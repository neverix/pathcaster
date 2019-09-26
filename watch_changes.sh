#!/bin/sh
find . -name '*.go' | entr -c go run cmd/main.go
