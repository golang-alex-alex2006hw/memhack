
LDFLAGS_STATIC=--ldflags '-s -extldflags "-static"'
LDFLAGS_DYNAMIC=--ldflags '-s'

all: clean build test

memhack:
	@echo Building binaries ...
	@go build -v $(LDFLAGS_STATIC) -o build/hackme platform/hackme/main.go
	@go build -v $(LDFLAGS_STATIC) -o build/memhack platform/memhack/main.go
	@upx -o build/hackme.packed build/hackme >/dev/null
	@upx -o build/memhack.packed build/memhack >/dev/null

build: memhack

clean:
	@echo Cleaning up previous build ...
	@rm -f build/hackme build/hackme.packed build/memhack build/memhack.packed

packages:
	@echo Getting system libraries ...
	sudo apt-get install -y upx

test:
	@echo Testing ...
	@go test ./...
