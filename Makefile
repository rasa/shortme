version = 1.2.0

ifeq ($(OS),Windows_NT)
	export PATH := "C:/cygwin/bin:$(PATH)"
	EXTENSION=.exe
else
	EXTENSION=""
endif

dep:
	go get -d ./...

test: build
	go test -v ./...

vet:
	go list ./... | grep -v "./vendor*" | xargs go vet

fmt: 
	find . -type f -name "*.go" | grep -v "./vendor*" | xargs gofmt -s -w

build: dep vet fmt
	go build -ldflags="-X github.com/andyxning/shortme/conf.Version=$(version)" -o shortme$(EXTENSION) main.go

clean:
	rm -f shortme

.PHONY: fmt test dep build clean vet
