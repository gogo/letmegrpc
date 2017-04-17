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

package testcmd

import (
	"io/ioutil"
	golog "log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func runInBack(cmd *exec.Cmd) {
	go func() {
		err := cmd.Start()
		if err != nil {
			log.Fatalf("%v", err)
		}
		cmd.Wait()
	}()
	time.Sleep(2 * time.Second)
}

var log *golog.Logger = golog.New(os.Stderr, "", golog.Llongfile)

func TestCmd(t *testing.T) {
	port := "12345"
	server := exec.Command("letmetestserver", "--port="+port)
	runInBack(server)
	defer server.Process.Kill()
	log.Printf("server started")
	client := exec.Command("letmegrpc", "--addr=localhost:"+port, "--port=8080", "serve.proto")
	runInBack(client)
	defer client.Process.Kill()
	log.Printf("client started")
	resp, err := http.Get("http://localhost:8080/Label/Produce")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", string(body))
	if strings.Contains(string(body), "404") {
		log.Fatal("404")
	}
	if !strings.Contains(string(body), "<form") {
		log.Fatal("no form")
	}
	resp, err = http.Get(`http://localhost:8080/Label/Produce?json={%22Name%22:%22Walter%22}`)
	if err != nil {
		log.Fatal(err)
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", string(body))
	if strings.Contains(string(body), "404") {
		log.Fatal("404")
	}
	if !strings.Contains(string(body), `Walter`) {
		log.Fatal("could not find json value")
	}
}
