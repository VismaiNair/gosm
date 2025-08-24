# gosm Changelog
These changes are in chronological order.

## v0.1.0
- Initial Version of gosm
- `gosm build` command builds Go code to WASM (no flags)
- `gosm js` command copies the required js file for running WASM, `wasm_exec.js` (no flags)
- `gosm run` starts a server.
    - Default port for `gosm run`: 8080
    - Flags: `--port` or `-p`, followed by port number