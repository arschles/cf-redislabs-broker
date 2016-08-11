DOCKER_REPO ?= quay.io/
DOCKER_ORG ?= redislabs
DOCKER_VERSION ?= git-$(shell git rev-parse --short HEAD)

DOCKER_IMAGE := ${DOCKER_REPO}${DOCKER_ORG}/cf-redislabs-broker:${DOCKER_VERSION}

build-for-docker:
	GOOS=linux GOARCH=amd64 go build -o rootfs/bin/broker .

docker-build:
	docker build -t ${DOCKER_IMAGE} rootfs

docker-push:
	docker push ${DOCKER_IMAGE}
