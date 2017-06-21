# Nineties Namer

See and contribute to  [github.com/theRemix/go90s](https://github.com/theRemix/go90s)

This is an Implementation of [Nineties Namer Lib](https://github.com/theRemix/go90s)

Example [./routes/router.go](./routes/router.go)

## Usage

### Running
```sh
LISTEN=127.0.0.1 \
PORT=8000 \
go run main.go
```

### Developing

#### UI Dev
```sh
cd ui
yarn install
yarn run dev
```
_webpack proxies specific api requests to https://90s.fun/[random,names]_

#### API Dev

if you want to use the ui

run the server locally, update [./ui/config/index.js](./ui/config/index.js)

change `const apiUrl = liveApiUrl;` to `const apiUrl = localDevApiUrl`

### Curl interface

`curl 127.0.0.1:8080/random`
```
pete meets world
```

`curl 127.0.0.1:8080/names/v1.0.2`
```
trapper pager
```

`curl 127.0.0.1:8080/names/v1.0.3`
```
wu tang girls
```

### Web interface

`open http://localhost:8080`


## Tests

```sh
go test
```

## Building in Docker

```
make go ui docker
```

## Running in Docker

```
make run
```

