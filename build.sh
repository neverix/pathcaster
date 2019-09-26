#!/bin/sh
echo "Building Windows..."
GOOS=windows go build -o out/win-pathcaster.exe cmd/main.go
echo "Building Linux..."
GOOS=linux go build -o out/lin-pathcaster cmd/main.go
echo "Building Mac..."
GOOS=darwin go build -o out/mac-pathcaster cmd/main.go