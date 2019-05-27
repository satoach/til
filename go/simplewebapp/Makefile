SUBDIRS := $(shell ls -d */)

.PHONY: $(SUBDIRS) build clean run

build: $(SUBDIRS)

clean: $(SUBDIRS)

run:
	@echo "$@ command is not allowed at root dir."

$(SUBDIRS):
	$(MAKE) -C $@ $(MAKECMDGOALS)

targetdir:
	@echo $(SUBDIRS)
