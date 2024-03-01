# Gokr

## Background

Personnal prospective project on [Objective & Key-Results](https://felipecastro.com/en/okr/what-is-okr/) as a goals managment system written in React and Golang.

## Developers

`Gokr` front-end is written in [React](https://react.dev/) and uses the (vitejs)(https://vitejs.dev/) toolchain.

### Golang backend API

Make sure to have golang v1.21+ setup on your environment.

**Run the API from VSCode**

Run using Debug with the default `Server` configuration from `.vscode\launch.json`.

**Run the API from command-line**
```
go run cmd/api
```

**Run the API in docker**
```
TODO
```

**Run the API in Kubernetes**
```
TODO
```

### Running the UI

Make sure to have node v20.11.0+ setup in your environment.

```
cd web/frontend-react
npm run vite dev
```
The `dev` *script* will run the frontend on a local http server while watching changes in the source code letting you develop and test seamlessly.
```
VITE v5.1.0  ready in 204 ms

➜  Local:   http://localhost:5173/
➜  Network: use --host to expose
➜  press h + enter to show help
```