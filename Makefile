#!/usr/bin/make -f

GB=gb

all: test

build:
	@echo "Building..."
	@$(GB) build all

test: build
	@echo "Running tests..."
	@$(GB) test -test.v=true
