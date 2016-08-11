DOCKER_REGISTRY ?= quay.io/
DOCKER_IMAGE_PREFIX ?= deis
DOCKER_VERSION ?= git-$(shell git rev-parse --short HEAD)

DOCKER_IMAGE := ${DOCKER_REGISTRY}${DOCKER_IMAGE_PREFIX}/cf-redislabs-broker:${DOCKER_VERSION}

docker-build:
	docker build -t ${DOCKER_IMAGE} rootfs

docker-push:
	docker push ${DOCKER_IMAGE}
