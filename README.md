# gosm

## What is Gosm?
Gosm is a CLI wrapper for the Go Wasm complilation tools. Instead of remembering build flags and typing in environment variables, you only need to run simple commands.

## Installing Gosm
To learn how to install Gosm, please take a look at [the quickstart guide](QUICKSTART.md). 

## Commands

### `gosm build`
`gosm build` builds Go code to WASM binary format. You can specify either a directory or a go file to build.

__Example Usage__
1. Build all Go files in the current directory: ```gosm build```
2. Build all Go files in a specified directory: ```gosm build lib```
3. Build a specified Go file: ```gosm build main.go```

### `gosm js`
`gosm js` copies the required JS file from your go installation, `wasm_exec.js`. It is required to run any WASM.

__Example Usage__
1. Copy `wasm_exec.js`: `gosm js`

### `gosm run`
`gosm run` serves the static WASM files and HTML file over a webserver.

__Example Usage__
1. Run on https://localhost:8080: `gosm run`
2. Run with specified port (we use 8000 as an example): `gosm run --port 8080`

__Example Usage__
1. Copy the JS file into the current directory: ```gosm js```

