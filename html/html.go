// Copyright (c) 2015, Vastech SA (PTY) LTD. All rights reserved.
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

package html

import (
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"strings"
)

type html struct {
	*generator.Generator
	generator.PluginImports
	ioPkg      generator.Single
	reflectPkg generator.Single
	stringsPkg generator.Single
	jsonPkg    generator.Single
	strconvPkg generator.Single
}

func New() *html {
	return &html{}
}

func (p *html) Name() string {
	return "html"
}

func (p *html) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *html) typeName(name string) string {
	return p.TypeName(p.ObjectNamed(name))
}

func (p *html) writeError() {
	p.P(`if err != nil {`)
	p.In()
	p.P(`if err == `, p.ioPkg.Use(), `.EOF {`)
	p.In()
	p.P(`return`)
	p.Out()
	p.P(`}`)
	p.P(`w.Write([]byte(err.Error()))`)
	p.P(`return`)
	p.Out()
	p.P(`}`)
}

func (p *html) w(s string) {
	p.P(`w.Write([]byte("`, s, `"))`)
}

func formable(msg *descriptor.DescriptorProto) bool {
	for _, f := range msg.GetField() {
		if f.IsRepeated() {
			return false
		}
		if f.IsMessage() {
			return false
		}
	}
	return true
}

func (p *html) getInputType(method *descriptor.MethodDescriptorProto) *descriptor.DescriptorProto {
	fileDescriptorSet := p.AllFiles()
	inputs := strings.Split(method.GetInputType(), ".")
	packageName := inputs[1]
	messageName := inputs[2]
	msg := fileDescriptorSet.GetMessage(packageName, messageName)
	if msg == nil {
		p.Fail("could not find message ", method.GetInputType())
	}
	return msg
}

func (p *html) generateSets(servName string, method *descriptor.MethodDescriptorProto) {
	msg := p.getInputType(method)
	if !formable(msg) {
		return
	}
	p.P(`fieldnames := []string{`)
	p.In()
	for _, f := range msg.GetField() {
		p.P(`"`, f.GetName(), `",`)
	}
	p.Out()
	p.P(`}`)
	p.P(`isString := []bool{`)
	p.In()
	for _, f := range msg.GetField() {
		if f.IsString() {
			p.P(`true,`)
		} else {
			p.P(`false,`)
		}
	}
	p.Out()
	p.P(`}`)
	p.P(`isBool := []bool{`)
	p.In()
	for _, f := range msg.GetField() {
		if f.GetType() == descriptor.FieldDescriptorProto_TYPE_BOOL {
			p.P(`true,`)
		} else {
			p.P(`false,`)
		}
	}
	p.Out()
	p.P(`}`)
	p.P(`fields := make([]string, 0, len(fieldnames))`)
	p.P(`for i, name := range fieldnames {`)
	p.In()
	p.P(`v := req.FormValue(name)`)
	p.P(`if len(v) > 0 {`)
	p.In()
	p.P(`someValue = true`)
	p.P(`if isString[i] {`)
	p.In()
	p.P(`fields = append(fields, "\"" + name + "\":" + `, p.strconvPkg.Use(), `.Quote(v))`)
	p.Out()
	p.P(`} else if isBool[i] {`)
	p.In()
	p.P(`if v == "on" {`)
	p.In()
	p.P(`fields = append(fields, "\"" + name + "\":" + "true")`)
	p.Out()
	p.P(`} else {`)
	p.In()
	p.P(`fields = append(fields, "\"" + name + "\":" + "false")`)
	p.Out()
	p.P(`}`)
	p.Out()
	p.P(`} else {`)
	p.In()
	p.P(`fields = append(fields, "\"" + name + "\":" + v)`)
	p.Out()
	p.P(`}`)
	p.Out()
	p.P(`}`)
	p.P(`if someValue {`)
	p.In()
	p.P(`s := "{" + `, p.stringsPkg.Use(), `.Join(fields, ",") + "}"`)
	p.P(`err := `, p.jsonPkg.Use(), `.Unmarshal([]byte(s), msg)`)
	p.writeError()
	p.Out()
	p.P(`}`)
	p.Out()
	p.P(`}`)
}

func (p *html) generateForm(servName string, method *descriptor.MethodDescriptorProto) {
	msg := p.getInputType(method)
	p.w(`<div class=\"container\"><div class=\"jumbotron\">`)
	p.w(`<h3>` + servName + ` - ` + method.GetName() + `</h3>`)
	p.P(`s := "<form action=\"/`, servName, `/`, method.GetName(), `\" method=\"GET\" role=\"form\">"`)
	p.P(`w.Write([]byte(s))`)
	if !formable(msg) {
		panic("I don't think it is complicated")
		p.w(`<div class=\"form-group\">`)
		p.w(`Json for ` + method.GetInputType()[1:] + ` : <input name=\"json\" type=\"text\"><br>`)
		p.w(`</div>`)
	} else {
		for _, f := range msg.GetField() {
			if f.GetType() == descriptor.FieldDescriptorProto_TYPE_BOOL {
				p.w(`<div class=\"checkbox\">`)
				p.w(`<label for=\"` + f.GetName() + `\">`)
				p.w(`<input id=\"` + f.GetName() + `\" name=\"` + f.GetName() + `\" type=\"checkbox\"/>`)
				p.w(f.GetName())
				p.w(`</label>`)
			} else {
				p.w(`<div class=\"form-group\">`)
				p.w(`<label for=\"` + f.GetName() + `\">` + f.GetName() + `</label>`)
				p.w(`<input id=\"` + f.GetName() + `\" name=\"` + f.GetName() + `\" type=\"text\" class=\"form-control\"/><br>`)
			}
			p.w(`</div>`)
		}
	}
	p.w(`<button type=\"submit\" class=\"btn btn-primary\">Submit</button></form></div></div>`)
}

func (p *html) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)
	httpPkg := p.NewImport("net/http")
	p.jsonPkg = p.NewImport("encoding/json")
	p.ioPkg = p.NewImport("io")
	contextPkg := p.NewImport("golang.org/x/net/context")
	p.reflectPkg = p.NewImport("reflect")
	p.stringsPkg = p.NewImport("strings")
	p.strconvPkg = p.NewImport("strconv")
	logPkg := p.NewImport("log")
	grpcPkg := p.NewImport("google.golang.org/grpc")

	p.P(`var htmlstringer = func(v interface{}) ([]byte, error) {`)
	p.In()
	p.P(`header := []byte("<div class=\"container\"><pre>")`)
	p.P(`data, err := `, p.jsonPkg.Use(), `.MarshalIndent(v, "", "\t")`)
	p.P(`if err != nil {`)
	p.In()
	p.P(`return nil, err`)
	p.Out()
	p.P(`}`)
	p.P(`footer := []byte("</pre></div>")`)
	p.P(`return append(append(header, data...), footer...), nil`)
	p.Out()
	p.P(`}`)

	p.P(`func SetHtmlStringer(s func(interface{}) ([]byte, error)) {`)
	p.In()
	p.P(`htmlstringer = s`)
	p.Out()
	p.P(`}`)

	p.P(`func Serve(httpAddr, grpcAddr string, opts ...`, grpcPkg.Use(), `.DialOption) {`)
	p.In()
	p.P(`conn, err := `, grpcPkg.Use(), `.Dial(grpcAddr, opts...)`)
	p.P(`if err != nil {`)
	p.In()
	p.P(logPkg.Use(), `.Fatalf("Dial(%q) = %v", grpcAddr, err)`)
	p.Out()
	p.P(`}`)
	for _, s := range file.GetService() {
		origServName := s.GetName()
		servName := generator.CamelCase(origServName)
		p.P(origServName, `Client := New`, servName, `Client(conn)`)
		p.P(origServName, `Server := NewHTML`, servName, `Server(`, origServName, `Client)`)
		for _, m := range s.GetMethod() {
			p.P(httpPkg.Use(), `.HandleFunc("/`, servName, `/`, m.GetName(), `", `, origServName, `Server.`, m.GetName(), `)`)
		}
	}
	p.P(`if err := `, httpPkg.Use(), `.ListenAndServe(httpAddr, nil); err != nil {`)
	p.In()
	p.P(logPkg.Use(), `.Fatal(err)`)
	p.Out()
	p.P(`}`)
	p.Out()
	p.P(`}`)

	for _, s := range file.GetService() {
		origServName := s.GetName()
		servName := generator.CamelCase(origServName)
		p.P(`type html`, servName, ` struct {`)
		p.In()
		p.P(`client `, servName, `Client`)
		p.Out()
		p.P(`}`)

		p.P(`func NewHTML`, servName, `Server(client `, servName, `Client) *html`, servName, ` {`)
		p.In()
		p.P(`return &html`, servName, `{client}`)
		p.Out()
		p.P(`}`)

		for _, m := range s.GetMethod() {
			p.P(`func (this *html`, servName, `) `, m.GetName(), `(w `, httpPkg.Use(), `.ResponseWriter, req *`, httpPkg.Use(), `.Request) {`)
			p.In()
			p.w("<html>")
			p.w("<head>")
			p.w("<title>" + servName + " - " + m.GetName() + "</title>")
			p.w(`<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css\">`)
			p.w(`<script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/js/bootstrap.min.js\"></script>`)
			p.w(`<script src=\"//code.jquery.com/jquery-1.11.2.min.js\"></script>`)
			p.w("</head>")
			p.P(`jsonString := req.FormValue("json")`)
			p.P(`someValue := false`)
			p.P(`msg := &`, p.typeName(m.GetInputType()), `{}`)
			p.P(`if len(jsonString) > 0 {`)
			p.In()
			p.P(`err := `, p.jsonPkg.Use(), `.Unmarshal([]byte(jsonString), msg)`)
			p.writeError()
			p.P(`someValue = true`)
			p.Out()
			p.P(`} else {`)
			p.In()
			p.generateSets(servName, m)
			p.Out()
			p.P(`}`)
			p.generateForm(servName, m)
			p.P(`if someValue {`)
			p.In()
			if !m.GetClientStreaming() {
				if !m.GetServerStreaming() {
					p.P(`reply, err := this.client.`, m.GetName(), `(`, contextPkg.Use(), `.Background(), msg)`)
					p.writeError()
					p.P(`out, err := htmlstringer(reply)`)
					p.writeError()
					p.P(`w.Write(out)`)
				} else {
					p.P(`down, err := this.client.`, m.GetName(), `(`, contextPkg.Use(), `.Background(), msg)`)
					p.writeError()
					p.P(`for {`)
					p.In()
					p.P(`reply, err := down.Recv()`)
					p.writeError()
					p.P(`out, err := htmlstringer(reply)`)
					p.writeError()
					p.w(`<p>`)
					p.P(`w.Write(out)`)
					p.w(`</p>`)
					p.P(`w.(`, httpPkg.Use(), `.Flusher).Flush()`)
					p.Out()
					p.P(`}`)
				}
			} else {
				if !m.GetServerStreaming() {
					p.P(`up, err := this.client.Upstream(`, contextPkg.Use(), `.Background())`)
					p.writeError()
					p.P(`err = up.Send(msg)`)
					p.writeError()
					p.P(`reply, err := up.CloseAndRecv()`)
					p.writeError()
					p.P(`out, err := htmlstringer(reply)`)
					p.writeError()
					p.P(`w.Write(out)`)
				} else {
					p.P(`bidi, err := this.client.Bidi(`, contextPkg.Use(), `.Background())`)
					p.writeError()
					p.P(`err = bidi.Send(msg)`)
					p.writeError()
					p.P(`reply, err := bidi.Recv()`)
					p.writeError()
					p.P(`out, err := htmlstringer(reply)`)
					p.writeError()
					p.P(`w.Write(out)`)
				}
			}
			p.Out()
			p.P(`}`)
			p.w("</html>")
			p.Out()
			p.P(`}`)
		}
	}
}
