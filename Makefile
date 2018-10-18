# Makefile
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_BUILD_RACE=$(GO_CMD) build -race
GO_TEST=$(GO_CMD) test
GO_TEST_VERBOSE=$(GO_CMD) test -v
GO_TEST_COVER=$(GO_CMD) test -cover
GO_INSTALL=$(GO_CMD) install -v


all: test build-server run

test: # run tests

build-server: # build serevr

run: # run server