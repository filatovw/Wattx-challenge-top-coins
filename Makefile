.PHONY:all
SUBDIRS := api price rank pricelist 
TOPTARGETS := build clear push pull test check

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