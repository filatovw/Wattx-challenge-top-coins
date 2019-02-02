SHELL	:= /bin/bash -o pipefail

BRANCH	?= $(shell git rev-parse --abbrev-ref HEAD)
TAG	?= $(shell git describe --always --tags)

SERVICE	?= $(notdir $(CURDIR))

DSTIMG	:= services/$(SERVICE):$(BRANCH)
TAGIMG	:= services/$(SERVICE):$(BRANCH)-$(TAG)

PHONY: build
build:
    $(info >>> $@)
    docker build --rm -t $(DSTIMG) $(CURDIR)
	; docker tag $(DSTIMG) $(TAGIMG)