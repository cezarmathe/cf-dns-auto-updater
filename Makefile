SRC=main.go lib/*.go

all: build

# Build recipes
build: build_linux-amd64 build_linux-arm-v7 build_darwin-amd64

build_linux-amd64: $(SRC)
	GOOS=linux GOARCH=amd64 \
		go build \
		-o build/cf-dns-auto-updater_linux-amd64

build_linux-arm-v7: $(SRC)
	GOOS=linux GOARCH=arm GOARM=7 \
		go build -o build/cf-dns-auto-updater_linux-armv7

build_darwin-amd64: $(SRC)
	GOOS=darwin GOARCH=amd64 \
		go build -o build/cf-dns-auto-updater_darwin-amd64
# ===============

clean:
	rm -r build
