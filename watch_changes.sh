#!/bin/sh
find . -name '*.go' | entr -c go run main.go