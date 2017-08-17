
LDFLAGS_STATIC=--ldflags '-s -extldflags "-static"'
LDFLAGS_DYNAMIC=--ldflags '-s'

all: clean build test

memhack:
	@echo Building binaries ...
	@go build $(LDFLAGS_STATIC) -o build/hackme platform/hackme/main.go
	@go build $(LDFLAGS_STATIC) -o build/memhack platform/memhack/main.go
	@go build $(LDFLAGS_STATIC) -o build/memsearch platform/memsearch/main.go
	
build: memhack

clean:
	@echo Cleaning up previous build ...
	@rm -f build/hackme build/memhack build/memsearch

install:
	@echo Installing ptrace settings ...

test:
	@echo Testing ...
	@go test ./...
