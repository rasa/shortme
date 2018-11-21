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

GIT_IMPORT=github.com/rasa/shortme/conf
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GOLDFLAGS=-X $(GIT_IMPORT).GitCommit=$(GIT_COMMIT)$(GIT_DIRTY) -X $(GIT_IMPORT).Version=$(version)

all:	build ## Build release executable

dep: ## Build dependencies
	go get -d ./...

generate: ## Generate generated code
	go generate ./...

test: build # Run test suite
	go test -v ./...

vet: ## Run go vet
	go list ./... | grep -v "./vendor*" | xargs go vet

fmt: ## Run gofmt
	$(FIND) . -type f -name "*.go" | grep -v "./vendor*" | xargs gofmt -s -w

build: dep vet fmt ## Build release executable
	go build -ldflags="$(GOLDFLAGS)" -o shortme$(EXTENSION) main.go

dev: dep vet fmt ## Build development executable
	go build -ldflags="$(GOLDFLAGS) -X $(GIT_IMPORT).Tags=-dev" -o shortme$(EXTENSION) -tags dev main.go

run: ## Run executable
	touch nohup.out
	$(NOHUP)
	tail -f nohup.out

clean: ## Remove executables
	rm -f shortme shortme$(EXTENSION)

ungenerate: ## Remove generated files
	rm -f static/assets_vfsdata.go
	rm -f template/assets_vfsdata.go
	rm -f www/assets_vfsdata.go

.PHONY: all fmt test dep build clean vet dev generate ungenerate

help: ## Display help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "%-30s %s\n", $$1, $$2}'
