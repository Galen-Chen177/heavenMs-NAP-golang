#SHELL := /bin/zsh
BUILD_DATE             = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT             = $(shell git rev-parse HEAD)
GIT_REMOTE             = origin
GIT_BRANCH             = $(shell git rev-parse --abbrev-ref=loose HEAD | sed 's/heads\///')
GIT_TAG                = $(shell git describe --exact-match --tags HEAD 2>/dev/null)
GIT_TREE_STATE         = $(shell if [ -z "`git status --porcelain`" ]; then echo "clean" ; else echo "dirty"; fi)

BINARY = sse
GOOS = linux
GOARCH = amd64

VERSION?=?
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

# Symlink into GOPATH
# BUILD_DIR=${GOPATH}/src/sim
BUILD_DIR=$(shell git rev-parse --show-toplevel)
CURRENT_DIR=$(shell pwd)

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-s -w -X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"


.PHONY: build clean

all: clean build

build:clean
	GOOS=${GOOS} CGO_ENABLED=0 GOARCH=${GOARCH} go build ${LDFLAGS} -installsuffix cgo -o build/heavenMs-NAP-golang ./main.go

clean:
	rm -rf build/
