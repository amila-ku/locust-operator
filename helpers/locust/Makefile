SHELL = bash
VERSION ?= v0.0.1

build:
	docker build . -t amilaku/locust:${VERSION}

build-slim:
	docker build . -f Dockerfile.slim -t amilaku/locust:${VERSION}-slim

push:
	docker push amilaku/locust:${VERSION}

push-slim:
	docker push amilaku/locust:${VERSION}-slim