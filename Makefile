.PHONY: build
build:
	@mkdir -p ./bin
	go build -o ./bin ./cmd/...

.PHONY: build-release
build-release:
	@mkdir -p ./release
	CGO_ENABLED=0 go build -a -o ./release ./cmd/...

.PHONY: clean
clean:
	@rm -rf ./bin
	@rm -rf ./release

.PHONY: pb-gen
pb-gen:
	@./build/protobuf/pb-gen.sh

.PHONY: build-docker
build-docker:
	docker build -t demo-go-web -f build/docker/Dockerfile .
