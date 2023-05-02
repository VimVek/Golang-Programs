# PHONY: run

all: build

build:
	@go build -o sendm

run:
	@go run main.go

clean:
	@rm sendm