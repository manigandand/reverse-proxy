# Makefile
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_BUILD_RACE=$(GO_CMD) build -race
GO_TEST=$(GO_CMD) test
GO_TEST_VERBOSE=$(GO_CMD) test -v
GO_TEST_COVER=$(GO_CMD) test -cover
GO_INSTALL=$(GO_CMD) install -v
GLIDE_INSTALL=glide install

SERVER_BIN=recipe_proxy_server
SERVER_DIR=.
SERVER_MAIN=main.go

SOURCE_PKG_DIR= ./pkg
SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

all: deps test build-server run

deps:
	@echo "==>Installing dependencies ...";
	$(GLIDE_INSTALL)

test: # run tests
	@echo "==> Running tests ...";
	@$(GO_TEST_COVER) $(SOURCE_PKG_DIR)/...

build-server: # build serevr
	@echo "==> Building server ...";
	@$(GO_BUILD) -o $(SERVER_BIN) $(SERVER_DIR)/$(SERVER_MAIN) || exit 1;
	@chmod 755 $(SERVER_BIN)

run: # run server
	@echo "==> Running server ...";
	bash -c "source .env"
	./$(SERVER_BIN)