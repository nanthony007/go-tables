.DEFAULT_GOAL := run

all: build run-built


test:
	@echo "********************"
	# need go test stuff here
	

build:
	@echo "Building..."
	go build main.py

run:
	@echo "Running..."
	go run main.go

run-built:
	@echo "Running built file..."
	./main