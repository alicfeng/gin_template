GC=go
RELEASE_DIR=bin
SERVICE_BIN=server
INITIAL_BIN=initializate

build :
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GC} build -ldflags -w -o ${RELEASE_DIR}/${SERVICE_BIN}_origin cmd/web/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GC} build -ldflags -w -o ${RELEASE_DIR}/${INITIAL_BIN}_origin cmd/cli/initializate.go

compress :
	upx  ${RELEASE_DIR}/${SERVICE_BIN}_origin -o  ${RELEASE_DIR}/${SERVICE_BIN}
	upx  ${RELEASE_DIR}/${INITIAL_BIN}_origin -o  ${RELEASE_DIR}/${INITIAL_BIN}

clean :
	@if [ 0 -ne `ls -A $1|wc -w` ] ; then rm -rf ${RELEASE_DIR}/*; fi