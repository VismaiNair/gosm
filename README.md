# wasmate

## What is wasmate?
wasmate is a CLI wrapper for the Go Wasm complilation tools. Instead of remembering build flags and typing in environment variables, you only need to run simple commands.

## Installing wasmate
To learn how to install wasmate, please take a look at [the quickstart guide](QUICKSTART.md). 

## Commands

### `wasmate build`
`wasmate build` builds Go code to WASM binary format. You can specify either a directory or a go file to build.

__Example Usage__
1. Build all Go files in the current directory: ```wasmate build```
2. Build all Go files in a specified directory: ```wasmate build lib```
3. Build a specified Go file: ```wasmate build main.go```

### `wasmate js`
`wasmate js` copies the required JS file from your go installation, `wasm_exec.js`. It is required to run any WASM.

__Example Usage__
1. Copy `wasm_exec.js`: `wasmate js`

### `wasmate run`
`wasmate run` serves the static WASM files and HTML file over a webserver.

__Example Usage__
1. Run on https://localhost:8080: `wasmate run`
2. Run with specified port (we use 8000 as an example): `wasmate run --port 8080`

__Example Usage__
1. Copy the JS file into the current directory: ```wasmate js```

