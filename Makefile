.PHONY: all

GOOS=linux
BUILD_DIR=/home/dev/test/Astraltest/build

all: linux

linux:GOOS=linux
linux:build_all

windows:GOOS=windows
windows:BUILD_DIR=build_windows
windows:build_all

windows_windows:GOOS=windows
windows_windows:BUILD_DIR=build_windows
windows_windows:build_all_windows

build_all:
	@echo "Building the project"
	@python3 build.py ${BUILD_DIR} ${GOOS}

build_all_windows:
	@echo "Building the project"
	@py build.py ${BUILD_DIR} ${GOOS}