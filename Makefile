TARGET=evilginx
PACKAGES=core database log parser

.PHONY: all build clean install uninstall
all: build

build:
	@go build -o ./build/$(TARGET) -mod=vendor main.go

# install to ~/bin for easy access without adding build/ to PATH
install: build
	@mkdir -p ~/bin
	@cp ./build/$(TARGET) ~/bin/$(TARGET)
	@echo "Installed $(TARGET) to ~/bin/$(TARGET)"

clean:
	@go clean
	@rm -f ./build/$(TARGET)

# remove binary from ~/bin as well
uninstall:
	@rm -f ~/bin/$(TARGET)
	@echo "Uninstalled $(TARGET) from ~/bin/$(TARGET)"

# quick rebuild: clean then build, useful during active development
rebuild: clean build
