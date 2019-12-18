GOOS=linux
TARGET=amd64

all: vet build test

# Build recipes
build: $(TARGET)

$(TARGET): build/cf-dns-auto-updater_$(TARGET)

build/cf-dns-auto-updater_$(TARGET):
	GOARCH=$(TARGET) go build -o build/cf-dns-auto-updater_$(TARGET)
# ===============

clean:
	rm -r build

vet:
	go vet

test:
	go test
