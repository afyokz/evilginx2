TARGET=evilginx
PACKAGES=core database log parser

.PHONY: all build clean install
all: build

build:
	@go build -o ./build/$(TARGET) -mod=vendor main.go

# install to ~/bin for easy access without adding build/ to PATH
install: build
	@cp ./build/$(TARGET) ~/bin/$(TARGET)
	@echo "Installed $(TARGET) to ~/bin/$(TARGET)"

clean:
	@go clean
	@rm -f ./build/$(TARGET)
