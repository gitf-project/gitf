#!/usr/bin/env bash

echo "Cross compile gltf project start..."
echo ""
# Windows
export GOARCH="amd64" && GOOS="windows" && go build -o releases/gltf.exe && echo "For Windows finished @ releases/gltf.exe!"
# Windows arm64
export GOARCH="arm64" && GOOS="windows" && go build -o releases/gltf-arm64.exe && echo "For Windows arm64 finished @ releases/gltf-arm64.exe!"
# MacOS
export GOARCH="amd64" && GOOS="darwin" && go build -o releases/gltf.darwin && echo "For MacOS finished @ releases/gltf.darwin!"
# MacOS arm64
export GOARCH="arm64" && GOOS="darwin" && go build -o releases/gltf-arm64.darwin && echo "For MacOS arm64 finished @ releases/gltf-arm64.darwin!"
# Linux
export GOARCH="amd64" && GOOS="linux" && go build -o releases/gltf.linux && echo "For Linux finished @ releases/gltf.linux!"
# Linux arm64
export GOARCH="arm64" && GOOS="linux" && go build -o releases/gltf-arm64.linux && echo "For Linux arm64 finished @ releases/gltf-arm64.linux!"
echo ""
echo "Cross compile gltf project finished!"