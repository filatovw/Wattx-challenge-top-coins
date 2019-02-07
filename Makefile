DOCKER_PROJECT_PATH=/go/src/github.com/filatovw/Wattx-challenge-top-coins/
.PHONY:all
SUBDIRS := api price rank pricelist 
TOPTARGETS := build clear push pull test

$(TOPTARGETS): $(SUBDIRS)
$(SUBDIRS):
		$(MAKE) -C $@ $(MAKECMDGOALS)

.PHONY: $(TOPTARGETS) $(SUBDIRS)

PHONY:prune
prune:
	docker system prune -f --volumes

PHONY:configure
configure:
	docker-compose run --rm -e APP_ENVIRONMENT=dev configure

PHONY:codegen
codegen:
	protoc -I=. --go_out=plugins=grpc:. ./pricelist/pricelist/*.proto
	protoc -I=. --go_out=plugins=grpc:. ./price/price/*.proto
	protoc -I=. --go_out=plugins=grpc:. ./rank/rank/*.proto

PHONY:docker-codegen
docker-codegen:
	docker run -it --rm -v $(CURDIR):$(DOCKER_PROJECT_PATH) -w=$(DOCKER_PROJECT_PATH) filatovw/go-protobuf:latest bash -c 'go install ./... && make codegen'

PHONY:infra
infra:
	docker-compose stop consul-cluster consul-agent registrator
	docker-compose rm -f consul-cluster consul-agent registrator
	docker-compose up -d consul-cluster consul-agent registrator

PHONY:restart
restart:
	docker-compose stop api price rank pricelist
	docker-compose rm -f api price rank pricelist
	docker-compose up -d api price rank pricelist

PHONY:logs
logs:
	docker-compose logs -f api rank pricelist price

PHONY:stop
stop:
	docker-compose stop