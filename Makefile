IMAGE=theremix/nineties-namer
BIN_PATH=bin/main
PUBLIC_PATH=public
PORT=8080

default: run

clean: goclean uiclean

goclean:
	-@rm ${BIN_PATH}

godeps:
	-go get

go: goclean godeps
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o ${BIN_PATH}
	@echo "Built Linux bin to ${BIN_PATH}"

uideps:
	@(cd ui && yarn install)

uiclean:
	-@rm -rf ${PUBLIC_PATH}/*

ui: uiclean uideps
	@(cd ui && yarn run build)
	@mv ui/dist/* ${PUBLIC_PATH}
	@echo "built ui into ${PUBLIC_PATH}"

docker:
	docker build -t ${IMAGE} .

run: docker
	docker run -p ${PORT}:80 ${IMAGE}
