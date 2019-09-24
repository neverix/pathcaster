#!/bin/sh
GOOS=windows go build -o out/win-pathcaster.exe main.go
GOOS=linux go build -o out/lin-pathcaster main.go
GOOS=darwin go build -o out/mac-pathcaster main.go