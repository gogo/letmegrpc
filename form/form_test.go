package form

import (
	"encoding/json"
	"github.com/gogo/protobuf/parser"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"testing"
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
	formStr = CreateCustom("WeirdMethod", "form", "Weird", desc, CustomBuildField)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(handle),
	}
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		t.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()
	go server.Serve(lis)
	resp, err := http.Get("http://localhost:8080/WeirdMethod?json={%22Name%22:%22%22,%22WeirdName%22:%22another%20string%22,%22Number%22:null}")
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

func CustomBuildField(fileDescriptorSet *descriptor.FileDescriptorSet, msg *descriptor.DescriptorProto, f *descriptor.FieldDescriptorProto) string {
	fieldname := f.GetName()
	if fieldname != "WeirdName" {
		return BuildField(fileDescriptorSet, msg, f)
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
