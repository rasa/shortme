version = 1.2.0

ifeq ($(OS),Windows_NT)
	FIND=C:/cygwin/bin/find.exe
	EXTENSION=.exe
	NOHUP=cmd.exe /c start .\\shortme.exe
else
	FIND=find
	EXTENSION=
	NOHUP=nohup ./shortme &
endif

all:	build

dep:
	go get -d ./...

test: build
	go test -v ./...

vet:
	go list ./... | grep -v "./vendor*" | xargs go vet

fmt: 
	$(FIND) . -type f -name "*.go" | grep -v "./vendor*" | xargs gofmt -s -w

build: dep vet fmt
	go build -ldflags="-X github.com/rasa/shortme/conf.Version=$(version)" -o shortme$(EXTENSION) main.go

run:
	touch nohup.out
	$(NOHUP)
	tail -f nohup.out

clean:
	rm -f shortme

.PHONY: all fmt test dep build clean vet
