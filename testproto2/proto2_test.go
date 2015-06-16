package proto2

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestGeneratedJavascript(t *testing.T) {
	data, err := ioutil.ReadFile("./proto2.letmegrpc.go")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), "activeradio(,") {
		t.Fatalf("defaultBool is wrong for proto2")
	}
}
