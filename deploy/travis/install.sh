#!/bin/bash

set -ex

install_go_deps() { 
	pushd ${TRAVIS_BUILD_DIR}
	cd src
	go get -v ./...
	popd
}

install_go_deps
