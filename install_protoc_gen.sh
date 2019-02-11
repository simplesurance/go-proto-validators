#!/usr/bin/env bash

# This script install protoc-gen-go and protoc-gen-gogo in specific versions.

set -eu -o pipefail

install() {
	local gopath
	local dir
	local url="$1"
	local git_tag="$2"

	echo "installing $url version: $git_tag"

	cd /tmp/  # ensure the tool is not added to the go.mod file

	gopath="$(go env GOPATH | cut -d":" -f1)"
	dir="$gopath/src/$url"
	[ -d "$dir" ] && {
		cd "$dir"
		git checkout master
	}

	go get -d -u "$url"
	cd "$dir"
	git checkout "$git_tag"
	go install -v "$url"
}

install "github.com/golang/protobuf/protoc-gen-go" "v1.2.0"
install "github.com/gogo/protobuf/protoc-gen-gogo" "v1.2.0"
