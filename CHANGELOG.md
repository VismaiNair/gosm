# wasmate Changelog

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

__Note that all changes listed in this section are not yet released, and there are no guarantees that they will be released.__
- `wasmate build`: New `--js` flag to automatically copy `wasm_exec.js` (necessary for Go WASM execution) without the need of another command
- New command `wasmate init` to set up a skeleton WASM project
- New command `wasmate html` to add a skeleton HTML file with WASM
- New command `wasmate go` to add a skeleton Go file with wasm

## v0.2.0 - 2025/08/29

### Added

- `wasmate run`: New `--open` flag to automatically open the local server in the default browser.
  - `-o` is usable for this same feature.
- `wasmate build`: New `--output flag` to specify custom build output location
  - `-o` is usable for this same feature.

### Internal

- Created `browser` package for browser-related operations
- Created `envdetect` package for environment detection
- Reduced external dependencies to improve installation performance

## v0.1.0 - 2025/08/20

- **Added**
- Initial Version of wasmate
- `wasmate build` command builds Go code to WASM (no flags)
- `wasmate js` command copies the required js file for running WASM, `wasm_exec.js` (no flags)
- `wasmate run` starts a server.
  - Default port for `wasmate run`: 8080
  - Flags: `--port` or `-p`, followed by port number
