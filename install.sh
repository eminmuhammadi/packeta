#/bin/env bash

BUILD_TIME=$(date +%Y%m%d%H%M%S)
BUILD_ID=$(git rev-parse HEAD)
VERSION="0.0.1-beta"

go build -tags="github.com/eminmuhammadi/packeta" -ldflags "-w -s -X main.VERSION=$VERSION -X main.BUILD_TIME=$BUILD_TIME -X main.BUILD_ID=$BUILD_ID" -o packeta ./