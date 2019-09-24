#!/bin/sh
echo "Building Windows..."
GOOS=windows go build -o out/win-pathcaster.exe main.go
echo "Building Linux..."
GOOS=linux go build -o out/lin-pathcaster main.go
echo "Building Mac..."
GOOS=darwin go build -o out/mac-pathcaster main.go