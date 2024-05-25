current_time = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
git_description = $(shell git describe --always --dirty --tags --long)
linker_flags = '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

run:
	go run ./main.go

.PHONY: build
build:
	go build -ldflags=${linker_flags} -o=./build/main ./main.go

format:
	go fmt ./