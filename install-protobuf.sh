#!/usr/bin/env bash
set -xe

basename=protobuf-cpp-$PROTOBUF_VERSION

cd /home/travis

wget https://github.com/google/protobuf/releases/download/v$PROTOBUF_VERSION/$basename.tar.gz
tar xzf $basename.tar.gz

cd protobuf-$PROTOBUF_VERSION

./configure --prefix=/home/travis && make -j2 && make install
