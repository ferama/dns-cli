#! /bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $DIR

build() {
    EXT=""
    [[ $GOOS = "windows" ]] && EXT=".exe"
    echo "Building ${GOOS} ${GOARCH}"
    go build \
        -trimpath \
        -o ./bin/dns-cli .
}

build