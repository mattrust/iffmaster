# Copyright (c) 2025 Matthias Rustler
# Licensed under the MIT License - see LICENSE for details

.PHONY: all dist clean

SRC_DIR := cmd/iffmaster
RELEASE_DIR := release
ICON := icon.png

all: linux windows darwin wasm

linux:
	@echo "Building for Linux"
	mkdir -p $(RELEASE_DIR)
	cd $(RELEASE_DIR) && fyne package -os linux -src ../$(SRC_DIR) -name iffmaster-linux -icon $(ICON)

windows:
	@echo "Building for Windows"
	mkdir -p $(RELEASE_DIR)
	- cd $(RELEASE_DIR) && fyne package -os windows -src ../$(SRC_DIR) -name iffmaster-windows -icon $(ICON)

darwin:
	@echo "Building for Darwin"
	mkdir -p $(RELEASE_DIR)
	- cd $(RELEASE_DIR) && fyne package -os darwin -src ../$(SRC_DIR) -name iffmaster-darwin -icon $(ICON)

wasm:
	@echo "Building for WebAssembly"
	mkdir -p $(RELEASE_DIR)
	cd $(RELEASE_DIR) && fyne package -os web -src ../$(SRC_DIR) -name iffmaster-wasm -icon $(ICON)

clean:
	@echo "Cleaning up"
	rm -rf $(RELEASE_DIR)