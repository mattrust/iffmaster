#!/bin/bash

cd cmd/iffmaster

echo "Archives will be created in cmd/iffmaster"

echo "Building for Linux"
fyne package -os linux -name iffmaster-linux -icon icon.png

echo "Building for Windows"
fyne package -os windows -name iffmaster-windows -icon icon.png

echo "Building for Darwin"
fyne package -os darwin -name iffmaster-darwin -icon icon.png

echo "Building for WebAssembly"
fyne package -os web -name iffmaster-wasm -icon icon.png
