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
	"fmt"
	"strconv"
	"strings"

	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

var Header string = `

function addChildNode(ev) {
	ev.preventDefault();
	var thisNode = $(this).parents(".node:first");
	var myType = $(this).attr("type");
	var child = $(nodeFactory[myType]);
	activateLinks(child);
	$(">.children[type=" + myType + "]", thisNode).append(child);
}

function setChildNode(ev) {
	ev.preventDefault();
	var thisNode = $(this).parents(".node:first");
	var myType = $(this).attr("type");
	var child = $(nodeFactory[myType]);
	activateLinks(child);
	$(">.children[type=" + myType + "]", thisNode).append(child);
$(">.tooltipper", thisNode).hide();
	$(this).hide();
}

function delChildNode(ev) {
	ev.preventDefault();
	var thisNode = $(this).parents(".node:first");
	var parentNode = thisNode.parents(".node:first");
	thisNode.remove();
	var setChildLink = $(">a.set-child[fieldname='" + thisNode.attr('fieldname') + "']", parentNode);
	if (setChildLink.length > 0) {
		setChildLink.show();
                $(">.tooltipper", parentNode).show();
	}
}

function delField(ev) {
	ev.preventDefault();
	var thisField = $(this).parents(".field:first");
	thisField.remove();
}

function addElem(ev) {
	ev.preventDefault();
	var thisNode = $(this).parents(".node:first");
	var myType = $(this).attr("type");
	var myFieldname = $(this).attr("fieldname");
	if (myType == "bool") {
		var input = $('<div class="field form-group"><label class="col-sm-2 control-label">' + myFieldname + ': </label><div class="col-sm-8"><input name="' + myFieldname + '" type="checkbox" repeated="true"/></div><div class="col-sm-2"><a href="#" class="del-field btn btn-warning btn-sm" role="button">Remove</a></div></div>');
		$("a.del-field", input).click(delField);
		$("> .fields[fieldname='" + myFieldname + "']", thisNode).append(input);
	}
	if (myType == "number") {
		var input = $('<div class="field form-group"><label class="col-sm-2 control-label">' + myFieldname + ': </label><div class="col-sm-8"><input class="form-control" name="' + myFieldname + '" type="number" step="1" repeated="true"/></div><div class="col-sm-2"><a href="#"  class="del-field btn btn-warning btn-sm" role="button">Remove</a></div></div>');
		$("a.del-field", input).click(delField);
		$("> .fields[fieldname='" + myFieldname + "']", thisNode).append(input);
	}
	if (myType == "text") {
		var input = $('<div class="field form-group"><label class="col-sm-2 control-label">' + myFieldname + ': </label><div class="col-sm-8"><input class="form-control" name="' + myFieldname + '" type="text" repeated="true"/></div><div class="col-sm-2"><a href="#"  class="del-field btn btn-warning btn-sm" role="button">Remove</a></div></div>');
		$("a.del-field", input).click(delField);
		$("> .fields[fieldname='" + myFieldname + "']", thisNode).append(input);
	}
	if (myType == "float") {
		var input = $('<div class="field form-group"><label class="col-sm-2 control-label">' + myFieldname + ': </label><div class="col-sm-8"><input class="form-control" name="' + myFieldname + '" type="number" step="any" repeated="true"/></div><div class="col-sm-2"><a href="#"  class="del-field btn btn-warning btn-sm" role="button">Remove</a></div></div>');
		$("a.del-field", input).click(delField);
		$("> .fields[fieldname='" + myFieldname + "']", thisNode).append(input);
	}
}

function getUrlParameter(sParam)
{
    var sPageURL = window.location.search.substring(1);
    var sURLVariables = sPageURL.split('&');
    for (var i = 0; i < sURLVariables.length; i++)
    {
        var sParameterName = sURLVariables[i].split('=');
        if (sParameterName[0] == sParam)
        {
            return sParameterName[1];
        }
    }
}

function activateLinks(node) {
 	$("a.add-child", node).click(addChildNode);
	$("a.set-child", node).click(setChildNode);
	$("a.add-elem", node).click(addElem);
	$("a.del-child", node).click(delChildNode);
	$("a.del-field", node).click(delField);
	$('label[type=checkbox]').click(function() {
	    if ($(this).hasClass('active')) {
	        $(this).removeClass('active');
	    } else {
	        $(this).addClass('active');
	    }
	});
	$('[data-toggle="tooltip"]', node).tooltip();
}

function getChildren(el) {
	var json = {};
	$("> .children > .node", el).each(function(idx, node) {
		var nodeJson = getFields($(node));
		var allChildren = getChildren($(node));
		for (childType in allChildren) {
			nodeJson[childType] = allChildren[childType];
		}
		var nodeType = $(node).attr("fieldname");
		var isRepeated = $(node).attr("repeated") == "true";
		if (isRepeated) {
			if (!(nodeType in json)) {
				json[nodeType] = [];
			}
			json[nodeType].push(nodeJson);
		} else {
			json[nodeType] = nodeJson;
		}
	});
	return json
}

function isInt(value) {
  return !isNaN(value) &&
         parseInt(Number(value)) == value &&
         !isNaN(parseInt(value, 10));
}

function replaceAll(str, search, replace) {
	return str.split(search).join(replace);
}

function escapeIllegal(str) {
	return replaceAll(replaceAll(replaceAll(str, "%", "%25"), "&", "%26"), "#", "%23");
}

function getFields(node) {
	var nodeJson = {};
	$("> div.field > div ", $(node)).each(function(idx, field) {
		$("> input[type=text]", $(field)).each(function(idx, input) {
			nodeJson[$(input).attr("name")] = escapeIllegal($(input).val());
		});
		$("> input[type=number][step=any]", $(field)).each(function(idx, input) {
			nodeJson[$(input).attr("name")] = parseFloat($(input).val());
		});
		$("> input[type=number][step=1]", $(field)).each(function(idx, input) {
			nodeJson[$(input).attr("name")] = parseInt($(input).val());
		});
		$("> div > label.active", $(field)).each(function(idx, label) {
                        var input = $("> input[type=radio]", $(label));
			var v = input.val();
			if (v == "true") {
				nodeJson[input.attr("name")] = true;
			} else if (v == "false") {
				nodeJson[input.attr("name")] = false;
			} else {
				nodeJson[input.attr("name")] = parseInt(input.val());
			}
		});
		$("> select", $(field)).each(function(idx, input) {
			var textvalue = $(input).val();
			if (isInt(textvalue)) {
				nodeJson[$(input).attr("name")] = parseInt(textvalue);
			} else {
				nodeJson[$(input).attr("name")] = escapeIllegal(textvalue);
			}
		});
	});
	$("> div.fields > div ", $(node)).each(function(idx, field) {
		$("input[type=text]", $(field)).each(function(idx, input) {
			var fieldname = $(input).attr("name");
			if (!(fieldname in nodeJson)) {
				nodeJson[fieldname] = [];
			}
			nodeJson[fieldname].push(escapeIllegal($(input).val()));
		});
		$("input[type=checkbox]", $(field)).each(function(idx, input) {
			var fieldname = $(input).attr("name");
			if (!(fieldname in nodeJson)) {
				nodeJson[fieldname] = [];
			}
			nodeJson[fieldname].push($(input).is(':checked'));
		});
		$("input[type=number][step=any]", $(field)).each(function(idx, input) {
			var fieldname = $(input).attr("name");
			if (!(fieldname in nodeJson)) {
				nodeJson[fieldname] = [];
			}
			nodeJson[fieldname].push(parseFloat($(input).val()));
		});
		$("input[type=number][step=1]", $(field)).each(function(idx, input) {
			var fieldname = $(input).attr("name");
			if (!(fieldname in nodeJson)) {
				nodeJson[fieldname] = [];
			}
			nodeJson[fieldname].push(parseInt($(input).val()));
		});
		$("label.active", $(field)).each(function(idx, label) {
                        var input = $("> input[type=radio]", $(label));
			var fieldname = $(input).attr("name");
			if (!(fieldname in nodeJson)) {
				nodeJson[fieldname] = [];
			}
			nodeJson[fieldname].push(parseInt(input.val()));
		});
		$("select", $(field)).each(function(idx, input) {
			var fieldname = $(input).attr("name");
			if (!(fieldname in nodeJson)) {
				nodeJson[fieldname] = [];
			}
			var textvalue = $(input).val();
			if (isInt(textvalue)) {
				nodeJson[fieldname].push(parseInt(textvalue));
			} else {
				nodeJson[fieldname].push(escapeIllegal(textvalue));
			}
		});
	});

	return nodeJson;
}

function radioed(def, index, value) {
	if (value == undefined) {
		if (def == index) {
			return "checked"
		}
		return ""
	}
	if (index == parseInt(value)) {
		return "checked"
	}
	if (index == value) {
		return "checked"
	}
	return ""
}

function activeradio(def, index, value) {
	if (value == undefined) {
		if (def == index) {
			return "active"
		}
		return ""
	}
	if (index == parseInt(value)) {
		return "active"
	}
	if (index == value) {
		return "active"
	}
	return ""
}

function checked(value) {
	if (value == undefined) {
		return ""
	}
	if (value == true) {
		return "checked='checked'"
	}
	return ""
}

function selected(def, index, value) {
	if (value == undefined) {
		if (def == index) {
			return "selected='selected'"
		}
		return ""
	}
	if (index == parseInt(value)) {
		return "selected='selected'"
	}
	if (index == value) {
		return "selected='selected'"
	}
	return ""
}

function emptyIfNull(json) {
	if (json == undefined || json == null) {
		return JSON.parse("{}");
	}
	return json;
}

function getValue(json, name) {
	var value = json[name];
	if (value == undefined) {
		return JSON.parse("{}");
	}
	return value;
}

function getList(json, name) {
	var value = json[name];
	if (value == undefined) {
		return JSON.parse("[]");
	}
	return value;
}

function setLink(json, typ, fieldname, help) {
var display = "";
	if (json[fieldname] != undefined) {
display = 'style="display:none"';
}
        var tooltip = "";
        if (help.length > 0) {
		tooltip = '<a href="#" data-toggle="tooltip" ' + display + ' title="' + help + '" class="tooltipper"><span class="glyphicon glyphicon-question-sign" aria-hidden="true"></span></a>';
        }
	if (json[fieldname] == undefined) {
		return '<a href="#" type="' + typ + '" class="set-child btn btn-success btn-sm" role="button" fieldname="' + fieldname + '">Set ' + fieldname + '</a>' + tooltip;
	}
	return '<a href="#" type="' + typ + '" class="set-child btn btn-success btn-sm" role="button" fieldname="' + fieldname + '" style="display: none;">Set ' + fieldname + '</a>';
}

function setValue(def, value) {
	if (value == undefined) {
		if (def.length == 0) {
			return ""
		}
		return 'value="' + def + '"'
	}
	return 'value="' + value + '"'
}

function setRepValue(value) {
	if (value == undefined) {
		return ""
	}
	return 'value="' + value + '"'
}

function encode_utf8(s) {
  return unescape(encodeURIComponent(s));
}

function decode_utf8(s) {
  return decodeURIComponent(escape(s));
}

function HTMLEncode(str){
  var i = str.length,
      aRet = [];

  while (i--) {
    var iC = str[i].charCodeAt();
    if (iC < 65 || iC > 127 || (iC>90 && iC<97)) {
      aRet[i] = '&#'+iC+';';
    } else {
      aRet[i] = str[i];
    }
   }
  return aRet.join('');
}


function setStrValue(def, value) {
	if (value == undefined) {
		if (def == undefined) {
			return ""
		}
		return "value=" + JSON.stringify(HTMLEncode(decode_utf8(def)));
	}
	return "value=" + JSON.stringify(HTMLEncode(decode_utf8(value)));
}

function setRepStrValue(value) {
	if (value == undefined) {
		return ""
	}
	return "value=" + JSON.stringify(HTMLEncode(decode_utf8(value)));
}

`

func isBool(f *descriptor.FieldDescriptorProto) bool {
	return f.GetType() == descriptor.FieldDescriptorProto_TYPE_BOOL
}

func isString(f *descriptor.FieldDescriptorProto) bool {
	return f.GetType() == descriptor.FieldDescriptorProto_TYPE_STRING
}

func isEnum(f *descriptor.FieldDescriptorProto) bool {
	return f.GetType() == descriptor.FieldDescriptorProto_TYPE_ENUM
}

func getEnum(fileDescriptorSet *descriptor.FileDescriptorSet, f *descriptor.FieldDescriptorProto) *descriptor.EnumDescriptorProto {
	typeNames := strings.Split(f.GetTypeName(), ".")
	return fileDescriptorSet.GetEnum(typeNames[1], typeNames[2])
}

func isFloat(f *descriptor.FieldDescriptorProto) bool {
	switch f.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
		descriptor.FieldDescriptorProto_TYPE_FLOAT:
		return true
	}
	return false
}

func isNumber(f *descriptor.FieldDescriptorProto) bool {
	switch f.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_SINT32,
		descriptor.FieldDescriptorProto_TYPE_SINT64,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_UINT64:
		return true
	}
	return false
}

func getMessage(f *descriptor.FieldDescriptorProto, fileDescriptorSet *descriptor.FileDescriptorSet) *descriptor.DescriptorProto {
	typeNames := strings.Split(f.GetTypeName(), ".")
	messageName := typeNames[len(typeNames)-1]
	packageName := strings.Join(typeNames[1:len(typeNames)-1], ".")
	return fileDescriptorSet.GetMessage(packageName, messageName)
}

func BuilderMap(visited map[string]struct{}, fieldname string, repeated bool, msg *descriptor.DescriptorProto, fileDescriptorSet *descriptor.FileDescriptorSet) []string {
	s := []string{`"` + typ(fieldname, repeated, msg) + `": build` + typ(fieldname, repeated, msg) + `(emptyIfNull(null)),`}
	for _, f := range msg.GetField() {
		if !f.IsMessage() {
			continue
		}
		fieldMsg := getMessage(f, fileDescriptorSet)
		if _, ok := visited[msg.GetName()+"."+f.GetName()]; !ok {
			visited[msg.GetName()+"."+f.GetName()] = struct{}{}
			s = append(s, BuilderMap(visited, f.GetName(), f.IsRepeated(), fieldMsg, fileDescriptorSet)...)
		}
	}
	return s
}

func Init(methodName string, fieldname string, repeated bool, msg *descriptor.DescriptorProto) string {
	return `function init() {
	var root = $(nodeFactory["` + typ(fieldname, repeated, msg) + `"]);
	var jsonText = getUrlParameter("json");
	if (jsonText == undefined) {
		var json = emptyIfNull(null);
	} else {
		var json = JSON.parse(unescape(jsonText));
	}
	$("#form > .children").html(build` + typ(fieldname, repeated, msg) + `(json));
	activateLinks(root);
	$("a[id=submit]").click(function(ev) {
		ev.preventDefault();
		c = getChildren($("#form"));
		j = JSON.stringify(c["` + fieldname + `"]);
		window.location.assign("./` + methodName + `?json="+j);
	});
}
`
}

func typ(fieldname string, repeated bool, msg *descriptor.DescriptorProto) string {
	if repeated {
		return "RepeatedKeyword_" + msg.GetName() + "_" + fieldname
	}
	return msg.GetName() + "_" + fieldname
}

type FieldBuilder func(fileDescriptorSet *descriptor.FileDescriptorSet, msg *descriptor.DescriptorProto, f *descriptor.FieldDescriptorProto, help string, proto3 bool) string

func BuildField(fileDescriptorSet *descriptor.FileDescriptorSet, msg *descriptor.DescriptorProto, f *descriptor.FieldDescriptorProto, help string, proto3 bool) string {
	tooltip := ""
	colon := ":"
	if len(help) > 0 {
		help = strings.Replace(help, "`", "'", -1)
		help = strings.Replace(help, "'", "\\'", -1)
		help = strings.Replace(help, "\n", " ", -1)
		tooltip = ` <a href="#" data-toggle="tooltip" title="` + help + `"><span class="glyphicon glyphicon-question-sign" aria-hidden="true"></span></a>`
		colon = ""
	}
	fieldname := f.GetName()
	if f.IsMessage() {
		typName := typ(fieldname, f.IsRepeated(), getMessage(f, fileDescriptorSet))
		if !f.IsRepeated() {
			return `s += '<div class="children" type="` + typName + `">' + build` + typName + `(json["` + f.GetName() + `"]);
			s += '</div>';
		s += setLink(json, "` + typName + `", "` + fieldname + `", "` + strings.Replace(help, "\n", "\\n", -1) + `");
		`
		} else {
			return `s += '<div class="children" type="` + typName + `">';
			var ` + fieldname + ` = getList(json, "` + fieldname + `");
			for (var i = 0; i < ` + fieldname + `.length; i++) {
				s += build` + typName + `(` + fieldname + `[i]);
			}
			s += '</div>';
			s += '<a href="#" class="add-child btn btn-success btn-sm" role="button" type="` + typName + `">add ` + fieldname + `</a>` + tooltip + `';
			s += '<div class="field form-group"></div>';
			`
		}
	} else {
		if !f.IsRepeated() {
			if isBool(f) {
				defaultBool := "\"nothing\""
				if proto3 {
					defaultBool = "false"
				}
				if f.DefaultValue != nil {
					defaultBool = f.GetDefaultValue()
				}
				s := `s += '<div class="field form-group"><label class="col-sm-2 control-label">` + fieldname + tooltip + colon + ` </label>';
					`
				s += `s += '<div class="col-sm-10"><div class="btn-group" data-toggle="buttons">';
					`
				s += `s += 	'<label class="btn btn-primary ' + activeradio(` + defaultBool + `, false, json["` + fieldname + `"]) + '"><input type="radio" name="` + fieldname + `" value="false" ' + radioed(` + defaultBool + `, false, json["` + fieldname + `"]) + '/>No</label>';
					`
				s += `s += 	'<label class="btn btn-primary ' + activeradio(` + defaultBool + `, true, json["` + fieldname + `"]) + '"><input type="radio" name="` + fieldname + `" value="true" ' + radioed(` + defaultBool + `, true, json["` + fieldname + `"]) + '/>Yes</label>';
					`
				s += `s += '</div></div></div>';
					`
				return s
			} else if isEnum(f) {
				enum := getEnum(fileDescriptorSet, f)
				defaultEnum := "\"nothing\""
				if proto3 {
					defaultEnum = "0"
				}
				if f.DefaultValue != nil {
					for _, v := range enum.GetValue() {
						if v.GetName() == f.GetDefaultValue() {
							defaultEnum = strconv.Itoa(int(v.GetNumber()))
							break
						}
					}
				}
				if len(enum.GetValue()) <= 4 {
					s := `s += '<div class="field form-group"><label class="col-sm-2 control-label">` + fieldname + tooltip + colon + ` </label>';
					`
					s += `s += '<div class="col-sm-10"><div class="btn-group" data-toggle="buttons">';
					`
					for _, v := range enum.GetValue() {
						num := strconv.Itoa(int(v.GetNumber()))
						s += `s += 	'<label class="btn btn-primary ' + activeradio(` + defaultEnum + `, ` + num + `, json["` + fieldname + `"]) + '"><input type="radio" name="` + fieldname + `" value="` + num + `" ' + radioed(` + defaultEnum + `, ` + num + `, json["` + fieldname + `"]) + '/> ` +
							v.GetName() + `</label>';
						`
					}
					s += `s += '</div></div></div>';
					`
					return s
				} else {
					s := `s += '<div class="field form-group"><label class="col-sm-2 control-label">` + fieldname + tooltip + colon + ` </label><div class="col-sm-10">';
					s += '<select class="form-control" name="` + fieldname + `">';
					`
					for _, v := range enum.GetValue() {
						num := strconv.Itoa(int(v.GetNumber()))
						s += `s += 	'<option value="` + num + `" ' + selected(` + defaultEnum + `, ` + num + `, json["` + fieldname + `"]) + '>` + v.GetName() + `</option>';
						`
					}
					s += `s += '</select></div></div>';
					`
					return s
				}
			} else if isNumber(f) {
				def := "\"\""
				if proto3 {
					def = "0"
				}
				if f.DefaultValue != nil {
					def = f.GetDefaultValue()
				}
				return `s += '<div class="field form-group"><label class="col-sm-2 control-label">` + fieldname + tooltip + colon + ` </label><div class="col-sm-10"><input class="form-control" name="` + f.GetName() + `" type="number" step="1" '+setValue(` + def + `, json["` + f.GetName() + `"])+'/></div></div>';
				`
			} else if isFloat(f) {
				def := "\"\""
				if proto3 {
					def = "0"
				}
				if f.DefaultValue != nil {
					def = f.GetDefaultValue()
				}
				return `s += '<div class="field form-group"><label class="col-sm-2 control-label">` + fieldname + tooltip + colon + ` </label><div class="col-sm-10"><input class="form-control" name="` + f.GetName() + `" type="number" step="any" '+setValue(` + def + `, json["` + f.GetName() + `"])+'/></div></div>';
				`
			} else {
				def := "undefined"
				if proto3 {
					def = "\"\""
				}
				if f.DefaultValue != nil {
					def = strconv.Quote(f.GetDefaultValue())
				}
				return `s += '<div class="field form-group"><label class="col-sm-2 control-label">` + fieldname + tooltip + colon + ` </label><div class="col-sm-10"><input class="form-control" name="` + f.GetName() + `" type="text" '+setStrValue(` + def + `, json["` + f.GetName() + `"])+'/></div></div>';
				`
			}
		} else {
			if isBool(f) {
				s := `
				s += '<div class="fields" fieldname="` + fieldname + `">';
				var ` + fieldname + ` = getList(json, "` + fieldname + `");
				for (var i = 0; i < ` + fieldname + `.length; i++) {
					s += '<div class="field form-group"><label class="col-sm-2 control-label">` + fieldname + `: </label><div class="col-sm-8"><input name="` + fieldname + `" type="checkbox" repeated="true" ' + checked(` + fieldname + `[i]) + '/></div><div class="col-sm-2"><a href="#" class="del-field btn btn-warning btn-sm" role="button">Remove</a></div></div>';
				}
				s += '</div>';
				s += '<a href="#" fieldname="` + fieldname + `" class="add-elem btn btn-info btn-sm" role="button" type="bool">add ` + fieldname + `</a>` + tooltip + `';
				s += '<div class="field form-group"></div>';
				`
				return s
			} else if isNumber(f) || isEnum(f) {
				s :=
					`s += '<div class="fields" fieldname="` + fieldname + `">';
				var ` + fieldname + ` = getList(json, "` + fieldname + `");
				for (var i = 0; i < ` + fieldname + `.length; i++) {
					s += '<div class="field form-group"><label class="col-sm-2 control-label">` + fieldname + `: </label><div class="col-sm-8"><input class="form-control" name="` + fieldname + `" type="number" step="any" repeated="true" '+setRepValue(json["` + f.GetName() + `"][i])+'/></div><div class="col-sm-2"><a href="#" class="del-field btn btn-warning btn-sm" role="button">Remove</a></div></div>';
				}
				s += '</div>';
				s += '<a href="#" fieldname="` + fieldname + `" class="add-elem btn btn-info btn-sm" role="button" type="number">add ` + fieldname + `</a>` + tooltip + `';
				s += '<div class="field form-group"></div>';
				`
				return s
			} else if isFloat(f) {
				s :=
					`s += '<div class="fields" fieldname="` + fieldname + `">';
				var ` + fieldname + ` = getList(json, "` + fieldname + `");
				for (var i = 0; i < ` + fieldname + `.length; i++) {
					s += '<div class="field form-group"><label class="col-sm-2 control-label">` + fieldname + `: </label><div class="col-sm-8"><input class="form-control" name="` + fieldname + `" type="number" step="1" repeated="true" '+setRepValue(json["` + f.GetName() + `"][i])+'/></div><div class="col-sm-2"><a href="#" class="del-field btn btn-warning btn-sm" role="button">Remove</a></div></div>';
				}
				s += '</div>';
				s += '<a href="#" fieldname="` + fieldname + `" class="add-elem btn btn-info btn-sm" role="button" type="float">add ` + fieldname + `</a>` + tooltip + `';
				s += '<div class="field form-group"></div>';
				`
				return s
			} else {
				s :=
					`s += '<div class="fields" fieldname="` + fieldname + `">';
				var ` + fieldname + ` = getList(json, "` + fieldname + `");
				for (var i = 0; i < ` + fieldname + `.length; i++) {
					s += '<div class="field form-group"><label class="col-sm-2 control-label">` + fieldname + `: </label><div class="col-sm-8"><input class="form-control" name="` + fieldname + `" type="text" repeated="true" '+setRepStrValue(json["` + f.GetName() + `"][i])+'/></div><div class="col-sm-2"><a href="#" class="del-field btn btn-warning btn-sm" role="button">Remove</a></div></div>';
				}
				s += '</div>';
				s += '<a href="#" fieldname="` + fieldname + `" class="add-elem btn btn-info btn-sm" role="button" type="text">add ` + fieldname + `</a>` + tooltip + `';
				s += '<div class="field form-group"></div>';
				`
				return s
			}
		}
	}
	panic("unreachable")
}

func Builder(visited map[string]struct{}, root bool, fieldname string, help string, repeated bool, msg *generator.Descriptor, g *generator.Generator, proto3 bool, buildField FieldBuilder) string {
	s := []string{`function build` + typ(fieldname, repeated, msg.DescriptorProto) + `(json) {`}
	if repeated {
		s = append(s, `var s = '<div class="node" type="`+typ(fieldname, repeated, msg.DescriptorProto)+`" fieldname="`+fieldname+`" repeated="true">';`)
	} else {
		s = append(s, `if (json == undefined) {
		return "";
	}
	`)
		s = append(s, `var s = '<div class="node" type="`+typ(fieldname, repeated, msg.DescriptorProto)+`" fieldname="`+fieldname+`" repeated="false">';`)
	}
	if !root {
		s = append(s, `s += '<div class="row"><div class="col-sm-2">'`)
		s = append(s, `s += '<a href="#" class="del-child btn btn-danger btn-xs" role="button" fieldname="`+fieldname+`">Remove</a>'`)
		s = append(s, `s += '</div><div class="col-sm-10">'`)
		s = append(s, `s += '<label class="heading">`+fieldname+`</label>'`)
		s = append(s, `s += '</div></div>'`)
	}
	ms := []string{}
	for i, f := range msg.GetField() {
		help := g.Comments(fmt.Sprintf("%s,%d,%d", msg.Path(), 2, i))
		if f.IsMessage() {
			fieldMsg := g.ObjectNamed(f.GetTypeName()).(*generator.Descriptor)
			if _, ok := visited[msg.GetName()+"."+f.GetName()]; !ok {
				visited[msg.GetName()+"."+f.GetName()] = struct{}{}
				ms = append(ms, Builder(visited, false, f.GetName(), help, f.IsRepeated(), fieldMsg, g, proto3, buildField))
			}
		}
		s = append(s, buildField(g.AllFiles(), msg.DescriptorProto, f, help, proto3))
	}
	if root {
		s = append(s, `
			s += '</div>';
			var node = $(s);
			activateLinks(node);
			return node;
		}`)
	} else {
		s = append(s, `
		s += '</div>';
		return s;
		}`)
	}
	f := strings.Join(s, "\n")
	ms = append(ms, f)
	return strings.Join(ms, "\n\n")
}

func Create(methodName, packageName, messageName string, g *generator.Generator) string {
	return CreateCustom(methodName, packageName, messageName, g, BuildField)
}

func CreateCustom(methodName, packageName, messageName string, g *generator.Generator, buildField FieldBuilder) string {
	msg := g.ObjectNamed("." + packageName + "." + messageName).(*generator.Descriptor)
	proto3 := g.AllFiles().IsProto3(packageName, messageName)
	text := `
	<form class="form-horizontal">
	<div id="form"><div class="children"></div></div>
    <a href="#" id="submit" class="btn btn-primary" role="button">Submit</a>
    </form>
    `
	text += `
	<script>`
	text += Header
	text += `var nodeFactory = {` + strings.Join(BuilderMap(make(map[string]struct{}),
		"RootKeyword", false, msg.DescriptorProto, g.AllFiles()), "\n") + `}
	`
	text += Builder(make(map[string]struct{}), true, "RootKeyword", "", false, msg, g, proto3, buildField)
	text += Init(methodName, "RootKeyword", false, msg.DescriptorProto)
	text += `
	init();

	</script>

	<style>

	.node{
		padding-left: 2em;
		min-height:20px;
	    padding:10px;
	    margin-top:10px;
	    margin-bottom:20px;
	    //border-left:0.5px solid #999;
	    -webkit-border-radius:4px;
	    -moz-border-radius:4px;
	    border-radius:4px;
	    -webkit-box-shadow:inset 0 1px 1px rgba(0, 0, 0, 0.05);
	    -moz-box-shadow:inset 0 1px 1px rgba(0, 0, 0, 0.05);
	    box-shadow:inset 0 1px 1px rgba(0, 0, 0, 0.05);
	    background-color:#eaeaea;
	}

	.node .node {
		background-color:#e2e2e2;
	}

	.node .node .node {
		background-color:#d9d9d9;
	}

	.node .node .node .node {
		background-color:#d1d1d1;
	}

	.node .node .node .node .node {
		background-color:#c7c7c7;
	}

	.node .node .node .node .node .node {
		background-color:#c0c0c0;
	}

	label{
	        font-weight: normal;
	}

	.heading {
		font-weight: bold;
	}

	</style>
	`
	return text
}
