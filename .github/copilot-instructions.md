# Copilot Instructions

## Repository layout

- The application source is in `phongdolphin/`; run repository commands from that directory unless a command is explicitly shown otherwise.
- `phongdolphin/main.go` is the Wails desktop entry point. It embeds `frontend/dist`, creates `App`, and binds the app methods to the frontend.
- `phongdolphin/app.go` contains the backend `App` implementation. Exported methods on `App` become callable from the frontend.
- `phongdolphin/frontend/` is a vanilla JavaScript Vite app, not a framework app.
- `phongdolphin/frontend/wailsjs/` contains Wails-generated JavaScript and TypeScript bindings; do not edit these files by hand.
- `phongdolphin/build/` contains Windows/macOS packaging resources and generated installer/build metadata.

## Build and run

From `phongdolphin/`:

```powershell
wails dev
wails build
wails build --target windows/amd64
```

`wails dev` starts the desktop app and Wails/Vite development workflow. `wails build` creates the production desktop application and refreshes the embedded frontend build.

From `phongdolphin/frontend/`:

```powershell
npm install
npm run dev
npm run build
npm run preview
```

The production frontend is written to `frontend/dist/`, which is embedded by `main.go` and therefore should not be edited directly.

## Tests and lint

There is currently no frontend test or lint script, and no Go test files. Do not invent an npm test/lint command.

Use the Go toolchain for backend package checks:

```powershell
cd H:\PhongDolphin\phongdolphin
go test ./...
go test ./... -run TestName
go vet ./...
gofmt -w *.go
```

For a newly named test, replace `TestName` with the exact test function name.

## Architecture and conventions

- Keep desktop lifecycle/window configuration in `main.go`; keep application behavior in methods on `App` in `app.go`.
- Frontend-to-backend calls use the generated modules under `frontend/wailsjs/go/main/`. Import those modules from `frontend/src/main.js` rather than calling runtime internals directly.
- The frontend currently uses direct DOM queries and assigns global functions for inline `onclick` handlers. Preserve this vanilla-JS style unless there is a reason to introduce a framework.
- `startup` stores Wails context on `App`; add startup-dependent backend behavior to this method or methods that receive/use `a.ctx`.
- Add persistent browser-facing assets under `frontend/src/assets/`; CSS entry points are `frontend/src/style.css` and `frontend/src/app.css`.
- Keep `frontend/package-lock.json` synchronized with `frontend/package.json` when dependencies change.
- `frontend/dist/`, `frontend/node_modules/`, and `build/bin/` are ignored build outputs.
