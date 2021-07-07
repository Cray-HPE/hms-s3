NAME ?= hms-s3
VERSION ?= $(shell cat .version)

all : image unittest

image:
	docker build --pull ${DOCKER_ARGS} --tag '${NAME}:${VERSION}' .

unittest:
	./runUnitTest.sh

