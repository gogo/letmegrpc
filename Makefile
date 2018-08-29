# Copyright (c) 2015, LetMeGRPCAuthors. All rights reserved.
#
# Redistribution and use in source and binary forms, with or without
# modification, are permitted provided that the following conditions are
# met:
#
#     * Redistributions of source code must retain the above copyright
# notice, this list of conditions and the following disclaimer.
#     * Redistributions in binary form must reproduce the above
# copyright notice, this list of conditions and the following disclaimer
# in the documentation and/or other materials provided with the
# distribution.
#
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
# "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
# LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
# A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
# OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
# SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
# LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
# DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
# THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
# (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
# OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

.PHONY: test

all:
	make nuke
	make install
	make regenerate
	make gofmt
	make install
	make test

install:
	go install github.com/gogo/protobuf/protoc-gen-gogo
	go install -v ./...

clean:
	go clean ./...

nuke:
	go clean -i ./...

test:
	go test -v ./test/...
	go test -v ./form/...

regenerate:
	go install github.com/gogo/letmegrpc/protoc-gen-letmegrpc
	(cd test && protoc --gogo_out=plugins=grpc:. --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/gogo/protobuf/protobuf/ grpc.proto)
	(cd test && protoc --letmegrpc_out=plugins=grpc:. --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/gogo/protobuf/protobuf/ grpc.proto)
	(cd letmetestserver/serve && protoc --gogo_out=plugins=grpc:. --proto_path=. serve.proto)
	(cd testimport && protoc --gogo_out=plugins=grpc:. --proto_path=.:../../../../ import.proto)
	(cd testimport && protoc --letmegrpc_out=plugins=grpc:. --proto_path=.:../../../../ import.proto)
	(cd testproto2 && protoc --gogo_out=plugins=grpc:. proto2.proto)
	(cd testproto2 && protoc --letmegrpc_out=plugins=grpc:. proto2.proto)

gofmt:
	gofmt -l -s -w .

