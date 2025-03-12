ifeq ($(OS),Windows_NT)
    BROWSER = start
    SHELL = C:/Program Files/Git/bin/bash.exe
else
    BROWSER = xdg-open
endif

GOOS ?= js
GOARCH ?= wasm

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o main.wasm
	
serve:
	$(BROWSER) "http://localhost:5000"