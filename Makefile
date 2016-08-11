DOCKER_REPO ?= quay.io/
DOCKER_ORG ?= redislabs
DOCKER_VERSION ?= git-$(shell git rev-parse --short HEAD)

DOCKER_IMAGE := ${DOCKER_REPO}${DOCKER_ORG}/cf-redislabs-broker:${DOCKER_VERSION}

build-linux:
	./bin/build-linux-amd64

docker-build: build-linux
	cp ./out/redislabs-service-broker rootfs/bin/broker
	docker build -t ${DOCKER_IMAGE} rootfs

docker-push:
	docker push ${DOCKER_IMAGE}
