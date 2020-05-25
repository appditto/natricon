# natricon backend
[![CI](https://github.com/appditto/natricon/workflows/CI/badge.svg)](https://github.com/appditto/natricon/actions?query=workflow%3ACI) [![Twitter Follow](https://img.shields.io/twitter/follow/appditto?style=social)](https://twitter.com/intent/follow?screen_name=appditto)

The backend and business logic for [natricon](https://natricon.com)

natricon is built in [GOLang](http://golang.org/)

## Requirements

The natricon backend requires ImageMagick development libraries to be installed. ImageMagick should be compiled with librsvg, libxml2, libpng, and libwebp.

## Natricon server build setup

```bash
# install dependencies
$ go get -u
# run in debugging mode
$ go run .

# build binary for production
$ go build . -o natricon
# execute natricon in production mode
$ GIN_MODE=release ./natricon

# For all options run
$ ./natricon -help
```

## WebAssembly (wasm) build setup

There is a WebAssembly reference implementation in the [wasm folder](https://github.com/appditto/natricon/tree/master/server/wasm)

This allows you to generate a natricon entirely on client-side from within the browser.

```bash
# To compile wasm
$ cd wasm
$ GOOS=js GOARCH=wasm go build -o main.wasm

# Running the sample
$ go get -u github.com/go-serve/goserve
$ ${GO_PATH}/bin/serve .
```