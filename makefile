# manual settings
AUTHOR=Olaf Michaelis <o.mic@web.de>

# auto generated
APP_NAME=$(shell pwd | tr '/' '\n' | tail -1)
SOURCE_DIR=src
TARGET_FOLDER=build
BINDATA=${SOURCE_DIR}/server/bindata.go

LOCAL_ARCH_BINARY=${TARGET_FOLDER}/$(shell arch)/${APP_NAME}

all: run_test run_build display_version run_benchmark
benchmark: run_benchmark
build: install_deps run_build
compress: run_compression
quick: run_unittest run_build
version: display_version
unittest: run_unittest
test: run_unittest run_test

install_deps:
	go get github.com/BurntSushi/toml
	go get github.com/pierrec/lz4
	go get github.com/mholt/archiver

run_benchmark:
	hyperfine "${LOCAL_ARCH_BINARY} -h"

run_build:
	maker/build.sh \
		"${SOURCE_DIR}" \
		"${APP_NAME}" \
		"${TARGET_FOLDER}" \
		"${AUTHOR}"

run_compression:
	maker/compress.sh "${TARGET_FOLDER}"

run_unittest:
	go test -cover -bench=. ${SOURCE_DIR}/*.go

run_test:
	testdata/test-keep-last.sh

display_version:
	${LOCAL_ARCH_BINARY} -V
