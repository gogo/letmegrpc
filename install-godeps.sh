#!/usr/bin/env bash
set -xe
mkdir -p $GOPATH/src/github.com/gogo
git clone https://github.com/gogo/protobuf $GOPATH/src/github.com/gogo/protobuf
go get google.golang.org/grpc
go get golang.org/x/net/context
