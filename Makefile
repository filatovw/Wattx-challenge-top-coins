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