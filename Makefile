IMAGE=theremix/nineties-namer
BIN_PATH=bin/main
PORT=8080

default: run

deps:
	-go get

go: deps
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o ${BIN_PATH}
	@echo "Built Linux bin to ${BIN_PATH}"

docker: go
	docker build -t ${IMAGE} .

run: docker
	docker run -p ${PORT}:80 ${IMAGE}
