#!/usr/bin/env bash
set -xe

basename=protoc-$PROTOBUF_VERSION-linux-x86_64

cd /home/travis

wget https://github.com/google/protobuf/releases/download/v$PROTOBUF_VERSION/$basename.zip
unzip $basename.zip
