BINARY_NAME=clipf
INSTALL_DIR=/usr/local/bin/

all: build 

install: build
	sudo mv ${BINARY_NAME} ${INSTALL_DIR}

build:
	go build -o ${BINARY_NAME}

clean:
	go clean
	rm -f ${BINARY_NAME}

format:
	go fmt *.go
