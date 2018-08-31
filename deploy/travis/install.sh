#!/bin/bash

set -ex

install_go_deps() { 
	pushd ${TRAVIS_BUILD_DIR}
	cd src
	go get -v -t ./...
	popd
}

install_go_deps
