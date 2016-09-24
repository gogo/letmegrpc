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

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func rand16() string {
	s := make([]string, 16)
	for i := range s {
		s[i] = strconv.FormatInt(int64(r.Intn(16)), 16)
	}
	return strings.Join(s, "")
}

func newTempDir() string {
	tmp := os.TempDir()
	ss := rand16()
	tmp = filepath.Join(tmp, "letmegrpc_"+ss)
	return tmp
}

func run(cmd *exec.Cmd) {
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("%s %s\n", string(out), err)
	}
}

var proto_path = flag.String("proto_path", ".", "")
var grpcAddr = flag.String("addr", "127.0.0.1:8081", "grpc address")
var httpAddr = flag.String("httpaddr", "127.0.0.1:8080", "http address")
var httpPort = flag.String("port", "", "http port. if empty, ignored; if non-empty, replaces httpaddr which will then listen on 127.0.0.1.")

func main() {
	flag.Parse()

	tmpDir := newTempDir()
	os.MkdirAll(tmpDir, 0777)
	defer os.RemoveAll(tmpDir)

	outDir := filepath.Join(tmpDir, "src", "tmpprotos")
	os.MkdirAll(outDir, 0777)

	cmdDir := filepath.Join(tmpDir, "src", "tmpprotos", "cmd", "letmegrpc"+rand16())
	os.MkdirAll(cmdDir, 0777)

	if len(flag.Args()) != 1 {
		flag.Usage()
		log.Fatalf("expected an input proto file")
	}
	filename := flag.Args()[0]

	if _, err := exec.LookPath("protoc"); err != nil {
		log.Fatalf("cannot find protoc in PATH")
	}
	run(exec.Command("protoc", "--gogo_out=plugins=grpc:"+outDir, "--proto_path="+*proto_path, filename))
	run(exec.Command("protoc", "--letmegrpc_out=plugins=grpc:"+outDir, "--proto_path="+*proto_path, filename))

	if *httpPort != "" {
		*httpAddr = "127.0.0.1:" + *httpPort
	}

	var mainStr = `package main

import tmpprotos "tmpprotos"
import "google.golang.org/grpc"

func main() {
	tmpprotos.Serve("` + *httpAddr + `", "` + *grpcAddr + `",
		tmpprotos.DefaultHtmlStringer,
		grpc.WithInsecure(), grpc.WithDecompressor(grpc.NewGZIPDecompressor()),
	)
}
`
	if err := ioutil.WriteFile(filepath.Join(cmdDir, "/main.go"), []byte(mainStr), 0777); err != nil {
		log.Fatalf("%s\n", err)
	}
	gorun := exec.Command("go", "run", "main.go")
	envs := os.Environ()
	for i, e := range envs {
		if strings.HasPrefix(e, "GOPATH") {
			envs[i] = envs[i] + ":" + tmpDir
		}
	}
	gorun.Env = envs
	gorun.Dir = cmdDir
	run(gorun)
}
