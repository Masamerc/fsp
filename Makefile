build: 
	@go build -o ./target/debug/fsp ./cmd/fsp.go

install:
	@go build -o fsp ./cmd
	@mv fsp $(GOPATH)/bin/
