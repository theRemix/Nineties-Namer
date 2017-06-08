# Nineties Namer

Go implementation of Nineties Namer

see [./routes/router.go](./routes/router.go)

## Usage

### Running
```sh
LISTEN=127.0.0.1 \
PORT=8080 \
go run main.go
```
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


