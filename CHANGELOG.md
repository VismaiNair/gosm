# wasmate Changelog
These changes are in chronological order.

## v0.1.0
- Initial Version of wasmate
- `wasmate build` command builds Go code to WASM (no flags)
- `wasmate js` command copies the required js file for running WASM, `wasm_exec.js` (no flags)
- `wasmate run` starts a server.
    - Default port for `wasmate run`: 8080
    - Flags: `--port` or `-p`, followed by port number