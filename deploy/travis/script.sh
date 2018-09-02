#!/bin/bash

set -ex

test() { 
	pushd ${TRAVIS_BUILD_DIR}
	cd src
	go test -cover -v ./...
	popd
}

test
