// Copyright (c) 2015, LetMeGRPCAuthors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package com_grpc

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type aServer struct{}

func (this *aServer) UnaryCall(c context.Context, s *MyRequest) (*MyResponse, error) {
	return &MyResponse{Value: s.Value}, nil
}
func (this *aServer) Downstream(m *MyRequest, s MyTest_DownstreamServer) error {
	for i := 0; i < int(m.Value); i++ {
		err := s.Send(&MyMsg{Value: int64(i)})
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *aServer) Upstreamy(s MyTest_UpstreamyServer) error {
	rec, err := s.Recv()
	sum := int64(0)
	for err == nil {
		sum += rec.Value
		rec, err = s.Recv()
	}
	return s.SendAndClose(&MyResponse{Value: sum})
}
func (this *aServer) Bidi(b MyTest_BidiServer) error {
	var err error
	msg := &MyMsg{}
	for {
		msg, err = b.Recv()
		if err != nil {
			break
		}
		err = b.Send(&MyMsg2{Value: msg.Value})
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		return nil
	}
	return err
}

func setup(t testing.TB, mytest MyTestServer) (*grpc.Server, string) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Failed to listen: %v", err)
	}
	_, port, err := net.SplitHostPort(lis.Addr().String())
	if err != nil {
		t.Fatalf("Failed to parse listener address: %v", err)
	}
	s := grpc.NewServer()
	RegisterMyTestServer(s, mytest)
	go s.Serve(lis)
	grpcAddr := "localhost:" + port
	return s, grpcAddr
}

func TestHTML(t *testing.T) {
	server, grpcAddr := setup(t, &aServer{})
	defer server.Stop()
	handler, err := NewHandler(grpcAddr, DefaultHtmlStringer, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(handler)
	defer httpServer.Close()
	time.Sleep(1e9)
	resp, err := http.Get(httpServer.URL + "/MyTest/UnaryCall")
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", string(body))
	if strings.Contains(string(body), "404") {
		t.Fatal("404")
	}
	if !strings.Contains(string(body), "<form") {
		t.Fatal("no form")
	}
	want := int64(5)
	req := &MyRequest{Value: want, Value2: 0}
	data, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = http.Get(fmt.Sprintf(httpServer.URL+"/MyTest/UnaryCall?json=%s", string(data)))
	if err != nil {
		t.Fatal(err)
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", string(body))
	if strings.Contains(string(body), "404") {
		t.Fatal("404")
	}
	if !strings.Contains(string(body), `"Value":`) {
		t.Fatal("could not find json value")
	}
	data = []byte(`{"value":"25921044673987072"}`)
	resp, err = http.Get(fmt.Sprintf("%s/MyTest/UnaryCall?json=%s", httpServer.URL, string(data)))
	if err != nil {
		t.Fatal(err)
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", string(body))
	if strings.Contains(string(body), "404") {
		t.Fatal("404")
	}
	if !strings.Contains(string(body), `"Value":`) {
		t.Fatal("could not find json value")
	}
}
