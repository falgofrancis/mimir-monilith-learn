SHELL = /bin/bash
export PROJECT_ROOT := $(shell dirname $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))
export DOCKER_NETWORK := grafanet

run:
	docker network rm $(DOCKER_NETWORK) || true
	docker network create $(DOCKER_NETWORK)
	PROJECT_ROOT=$(PROJECT_ROOT) DOCKER_NETWORK=$(DOCKER_NETWORK) docker-compose up -d --build --remove-orphans

stop:
	PROJECT_ROOT=$(PROJECT_ROOT) DOCKER_NETWORK=$(DOCKER_NETWORK) docker-compose down

# clean:
# 	@sh $(PROJECT_ROOT)/scripting/cleanup_docker.sh

runapp:
	cd app && go build && ./app