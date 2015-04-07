# letmegrpc

[![Build Status](https://drone.io/github.com/gogo/letmegrpc/status.png)](https://drone.io/github.com/gogo/letmegrpc/latest)

## Experimental

letmegrpc generates an html interface for a [grpc.io](http://www.grpc.io/) interface

## Installation

    mkdir -p $GOPATH/src/github.com/gogo/letmegrpc
    git clone https://github.com/gogo/letmegrpc $GOPATH/src/github.com/gogo/letmegrpc
    (cd $GOPATH/src/github.com/gogo/letmegrpc && make install)

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

    http://localhost:8080/OnionSeller/OnlyOnce