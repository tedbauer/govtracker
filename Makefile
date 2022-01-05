.PHONY: clean website

build:
	mkdir -p build

build/index.html: build
	cp web/index.html build

build/js: build
	babel --plugins transform-react-jsx web/components -d build/js && \
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js build/js

clean:
	rm -rf build

website: build/js build/index.html