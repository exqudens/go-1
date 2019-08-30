#!/bin/bash

function cleanDependencies() {
    echo 'cleanDependencies:' && \
    rm -rf "src/github.com/golang/example/stringutil/" && \
    rm -rf "src/github.com/golang/example/" && \
    rm -rf "src/github.com/golang/" && \
    rm -rf "src/github.com/" && \
    rm -rf "pkg/" && \
    echo '  done'
}

function dependencies() {
    echo 'dependencies:' && \
    env GOPATH="$(pwd)" go get 'github.com/golang/example/stringutil' && \
    echo '  done'
}

function cleanBuild() {
    echo 'cleanBuild:' && \
    rm -rf 'build/' && \
    echo '  done'
}

function build() {
    echo 'build:' && \
    mkdir -p 'build' && \
    env GOPATH="$(pwd)" go build -o 'build/application' 'src/application.go' && \
    echo '  done'
}

function run() {
    echo 'run:' && \
    build/application && \
    echo '  done'
}

for i in ${@}
do
	${i}
done && \
exit 0
