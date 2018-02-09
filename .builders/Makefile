DIRS ?=	$(filter %/,$(wildcard ./*/))

all: build

build clean:
	@set -e; for dir in $(DIRS); do ( \
	  echo "cd $$dir && make $@"; \
	  cd $$dir && make $@ \
	); done

re: clean build
