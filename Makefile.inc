BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
TAG	?= $(shell git describe --always --tags)

SERVICE	?= $(notdir $(CURDIR))

BUCKET := filatovw/services-$(SERVICE)
BRANCH_IMG := $(BUCKET):$(BRANCH)
BRANCH_TAG_IMG:= $(BUCKET):$(BRANCH)-$(TAG)

PHONY:build
build:
	docker build --rm -f Dockerfile -t $(BRANCH_IMG) -t $(BRANCH_TAG_IMG) $(CURDIR)/../

PHONY:push
push:
	docker push $(BRANCH_IMG)

PHONY:pull
pull:
	docker pull $(BRANCH_IMG)

PHONY:clear
clear:
	-docker images | grep $(BUCKET) | awk '{print $$3}' | xargs docker rmi -f {}

PHONY:check
check:
	golangci-lint run ./...

PHONY:test
test:
	go test ./...