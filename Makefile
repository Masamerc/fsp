ARCH ?= arm64
OS   ?= darwin


install:
	@echo "installing fsp: OS=$(OS) ARCH=$(ARCH)"
	@docker build --build-arg GOARCH=$(ARCH) --build-arg GOOS=$(OS) -t fsp-builder .
	@docker create --name fsp-container fsp-builder
	@docker cp fsp-container:/app/fsp /usr/local/bin/fsp
	@docker rm fsp-container

local_build: 
	@go build -o ./target/debug/fsp ./cmd/fsp.go

local_install:
	@go build -o fsp ./cmd
	@mv fsp $(GOPATH)/bin/
