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

package form

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gogo/pbparser"
	"github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	plugin "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
)

type Msg struct {
	Name      string
	WeirdName string
}

func TestCreateCustom(t *testing.T) {
	desc, err := parser.ParseFile("form.proto", ".")
	if err != nil {
		t.Fatal(err)
	}
	g := generator.New()
	g.Request = &plugin.CodeGeneratorRequest{ProtoFile: desc.File}
	g.Request.FileToGenerate = []string{"form.proto"}
	g.Request.Parameter = proto.String("plugins=grpc")
	g.CommandLineParameters(g.Request.GetParameter())
	g.WrapTypes()
	g.SetPackageNames()
	g.BuildTypeNameMap()
	g.Reset()
	g.SetFile(desc.File[0])
	formStr = CreateCustom("WeirdMethod", "weird.form", "Weird", g, CustomBuildField)
	testserver := httptest.NewServer(http.HandlerFunc(handle))
	defer testserver.Close()
	resp, err := http.Get(testserver.URL + "/WeirdMethod?json={%22Name%22:%22%22,%22WeirdName%22:%22another%20string%22,%22Number%22:null}")
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
	if !strings.Contains(string(body), `"WeirdName":"another string"`) {
		t.Fatal("could not find json value")
	}
}

func CustomBuildField(fileDescriptorSet *descriptor.FileDescriptorSet, msg *descriptor.DescriptorProto, f *descriptor.FieldDescriptorProto, help string, proto3 bool) string {
	fieldname := f.GetName()
	if fieldname != "WeirdName" {
		return BuildField(fileDescriptorSet, msg, f, help, proto3)
	}
	s := `s += '<div class="field form-group"><label class="col-sm-2 control-label">` + fieldname + `: </label><div class="col-sm-10">';
	s += '<select class="form-control" name="` + fieldname + `">';
	`
	stringOptions := []string{"option1", "another string", "192.168.1.1"}
	for _, stringOption := range stringOptions {
		s += `s += 	'<option value="` + stringOption + `" ' + selected("` + stringOption + `", json["` + fieldname + `"]) + '>` + stringOption + `</option>';
		`
	}
	s += `s += '</select></div></div>';
	`
	return s
}

var formStr = ""

var header = `
	<html>
	<head>
	<title>WeirdTest</title>
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css">
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/js/bootstrap.min.js"></script>
	</head>
	<body>
	<div class="container"><div class="jumbotron">
	<h3>WeirdTest</h3>
	`

var footer string = `
	</div>
	</body>
	</html>
	`

func handle(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(header))
	jsonString := req.FormValue("json")
	someValue := false
	msg := &Msg{}
	if len(jsonString) > 0 {
		err := json.Unmarshal([]byte(jsonString), msg)
		if err != nil {
			if err != io.EOF {
				w.Write([]byte(err.Error()))
				return
			}
			w.Write([]byte(err.Error()))
		}
		someValue = true
	}
	w.Write([]byte(formStr))
	if someValue {
		w.Write([]byte("<pre>" + jsonString + "</pre>"))
	}
	w.Write([]byte(footer))
}
