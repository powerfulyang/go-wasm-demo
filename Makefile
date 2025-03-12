ifeq ($(OS),Windows_NT)
    BROWSER = start
    SHELL = C:/Program Files/Git/bin/bash.exe
else
    BROWSER = xdg-open
endif

GOOS ?= js
GOARCH ?= wasm

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o main.wasm ./main.go

build-tiny-wasi:
	GOOS=wasip1 GOARCH=$(GOARCH) tinygo build -o tiny-wasi.wasm ./tiny.go

build-tiny-js:
	GOOS=$(GOOS) GOARCH=$(GOARCH) tinygo build -o tiny-js.wasm ./tiny.go
	
run-wasi:
	wasmedge ./tiny-wasi.wasm

serve:
	$(BROWSER) "http://localhost:5000"
	
all: build build-tiny-wasi build-tiny-js