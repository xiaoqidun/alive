#!/bin/bash
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o release/alive_amd64.exe -trimpath -ldflags "-H windowsgui -s -w -buildid=" alive.go
GOOS=windows GOARCH=arm64 CGO_ENABLED=0 go build -o release/alive_arm64.exe -trimpath -ldflags "-H windowsgui -s -w -buildid=" alive.go