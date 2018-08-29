> We are looking for a maintainer to take over this repository, file an issue if you are interested.

# letmegrpc

letmegrpc generates a web form gui from a [grpc](http://www.grpc.io/) definition.

[![Build Status](https://travis-ci.org/gogo/letmegrpc.svg?branch=master)](https://travis-ci.org/gogo/letmegrpc)

## Installation

    cd $GOPATH
    mkdir -p ./src/github.com/gogo/letmegrpc
    git clone https://github.com/gogo/letmegrpc ./src/github.com/gogo/letmegrpc
    git clone https://github.com/gogo/protobuf ./src/github.com/gogo/protobuf
    go get google.golang.org/grpc
    go get golang.org/x/net/context
    (cd ./src/github.com/gogo/letmegrpc && make install)

## Usage

Assuming you have a grpc server implementation running on your localhost on port 12345

    letmegrpc --addr=localhost:12345 --port=8080 grpc.proto

Now open your webbrowser and goto

    http://localhost:8080/ServiceName/MethodName

Here you will find a generated html web form.
Clicking Submit will send your newly populated message to your grpc server implementation and display the results.

## Example

    (cd $GOPATH/src/github.com/gogo/letmegrpc && make install)
    letmetestserver --port=12345 &
    (cd $GOPATH/src/github.com/gogo/letmegrpc/testcmd && letmegrpc --addr=localhost:12345 --port=8080 serve.proto

Open your webbrowser at

    http://localhost:8080/Label/Produce

![image](https://github.com/gogo/letmegrpc/blob/master/screenshot.png "ScreenShot")

## More Features

 - Comments that are attached to fields become tooltips.
 - proto2 defaults are filled into the the clean form.

## Customization

letmegrpc is just another protocol buffer code generation plugin.
Simply run:

    protoc -gogo_out=. grpc.proto
    protoc -letmegrpc_out=. grpc.proto

Now you can have the html generated code next to your generated message code.
It contains a:
  - The Serve function that is used to start the server.
  - SetHtmlStringer function that lets you customize your html output for each returned message, this is json by default.  This might be useful to return more links and create an explorable web site.


