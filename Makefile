
build:
	mkdir -p bin
	GOOS=linux GOARCH=386 go build -o bin/deploy-linux-386 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/deploy-linux-amd64 main.go
	GOOS=darwin GOARCH=386 go build -o bin/deploy-darwin-386 main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/deploy-darwin-amd64 main.go

.PHONY: build
