#!/usr/bin/env bash
set -xe

basename=protoc-$PROTOBUF_VERSION-linux-x86_64

cd /home/travis

wget https://github.com/google/protobuf/releases/download/v3.2.0/$basename.zip
unzip $basename.zip
cp $basename/bin/protoc ./bin/

