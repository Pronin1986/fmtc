// Copyright (c) 2017 Pronin S.V.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fmtc

import (
	"bytes"
	"testing"
)

func TestStackSuccessDelete(t *testing.T) {
	oneTag := tag{name: "b"}
	tagStack := stack{}
	tagStack.push(oneTag)

	if tagStack.pop(oneTag) == false {
		t.Error("tagStack.pop(oneTag) == false - сould not delete existing text")
	}
}

func TestStackFalseDelete(t *testing.T) {
	oneTag := tag{name: "b"}
	tagStack := stack{}
	tagStack.push(oneTag)

	oneTag2 := tag{name: "u"}
	if tagStack.pop(oneTag2) == true {
		t.Error("tagStack.pop(oneTag2) == true - error deleting")
	}
}

func TestStackFalseDeleteEmptyStack(t *testing.T) {
	tagStack := stack{}
	oneTag2 := tag{name: "u"}
	if tagStack.pop(oneTag2) == true {
		t.Error("tagStack.pop(oneTag2) == true - error deleting")
	}
}

func TestGetASCICode(t *testing.T) {
	var tests = []struct {
		input tag
		want  string
	}{
		{tag{name: "b"}, "\033[1m"},
		{tag{name: "strong"}, "\033[1m"},
		{tag{name: "u"}, "\033[4m"},
		{tag{name: "dim"}, "\033[2m"},
		{tag{name: "reverse"}, "\033[7m"},
		{tag{name: "blink"}, "\033[5m"},
		{tag{name: "black"}, "\033[30m"},
		{tag{name: "red"}, "\033[31m"},
		{tag{name: "green"}, "\033[32m"},
		{tag{name: "yellow"}, "\033[33m"},
		{tag{name: "blue"}, "\033[34m"},
		{tag{name: "magenta"}, "\033[35m"},
		{tag{name: "cyan"}, "\033[36m"},
		{tag{name: "grey"}, "\033[37m"},
		{tag{name: "white"}, "\033[97m"},
		{tag{name: "wrong"}, ""},
		{tag{name: "wrong wrong"}, ""},
		{tag{name: "hui"}, ""},
		{tag{name: "   "}, ""},
		{tag{name: "****"}, ""},
		{tag{name: "\033[1m"}, ""},
		{tag{name: "\033[1m;;;;;"}, ""},
		{tag{name: "\\\\\\////\\"}, ""},
	}

	for _, test := range tests {
		if got := getASCICode(test.input); got != test.want {
			t.Errorf("getASCICode(%v) = '%v' (SUCCEESS: %v)", test.input.name, got, test.want)
		}
	}
}

func TestDecorate(t *testing.T) {
	someVar := "SOME TEXT"
	var tests = []struct {
		want string
		got  string
	}{
		{"<b>BOLD</b>", "\033[1mBOLD\033[0m\033[0m"},
		{"<b>" + someVar + "</b>", "\033[1m" + someVar + "\033[0m\033[0m"},
		{"<b>BOLD <blue>BLUE</blue></b>", "\033[1mBOLD \033[34mBLUE\033[0m\033[1m\033[0m\033[0m"},
		{"<b>HELLO <blue>BLUE</blue> <bg color=\"green\">TEXT</bg></b>", "\033[1mHELLO \033[34mBLUE\033[0m\033[1m \033[42mTEXT\033[0m\033[1m\033[0m\033[0m"},
	}

	for _, test := range tests {
		if decorate(test.want) != test.got {
			t.Errorf("decorate(\"%v\") != ' %v '", test.want, test.got)
		}
	}
}

func TestPrint(t *testing.T) {
	_, err := Print("<b>BOLD <blue>BLUE</blue> TEXT</b>")
	if err != nil {
		t.Error("Print failed")
	}
}

func TestPrintln(t *testing.T) {
	_, err := Println("<b>BOLD <blue>BLUE</blue> <bg color=\"green\">TEXT</bg></b>")
	if err != nil {
		t.Error("Println failed")
	}
}

func TestPrintf(t *testing.T) {
	bold := "BOLD"
	blue := "BLUE"
	txt := "TEXT"

	_, err := Printf("<b>%v <blue>%v</blue> %v</b>", bold, blue, txt)
	if err != nil {
		t.Error("Printf failed")
	}
}

func TestSprint(t *testing.T) {
	someVar := "SOME TEXT"
	var tests = []struct {
		want string
		got  string
	}{
		{"<b>BOLD</b>", "\033[1mBOLD\033[0m\033[0m"},
		{"<b>" + someVar + "</b>", "\033[1m" + someVar + "\033[0m\033[0m"},
		{"<b>BOLD <blue>BLUE</blue></b>", "\033[1mBOLD \033[34mBLUE\033[0m\033[1m\033[0m\033[0m"},
		{"<b>HELLO <blue>BLUE</blue> <bg color=\"green\">TEXT</bg></b>", "\033[1mHELLO \033[34mBLUE\033[0m\033[1m \033[42mTEXT\033[0m\033[1m\033[0m\033[0m"},
	}

	for _, test := range tests {
		if Sprint(test.want) != test.got {
			t.Errorf("Sprint(\"%v\") != '%v'", test.want, test.got)
		}
	}
}

func TestSprintln(t *testing.T) {
	someVar := "SOME TEXT"
	var tests = []struct {
		want string
		got  string
	}{
		{"<b>BOLD</b>", "\033[1mBOLD\033[0m\033[0m\n"},
		{"<b>" + someVar + "</b>", "\033[1m" + someVar + "\033[0m\033[0m\n"},
		{"<b>BOLD <blue>BLUE</blue></b>", "\033[1mBOLD \033[34mBLUE\033[0m\033[1m\033[0m\033[0m\n"},
		{"<b>HELLO <blue>BLUE</blue> <bg color=\"green\">TEXT</bg></b>", "\033[1mHELLO \033[34mBLUE\033[0m\033[1m \033[42mTEXT\033[0m\033[1m\033[0m\033[0m\n"},
	}

	for _, test := range tests {
		if Sprintln(test.want) != test.got {
			t.Errorf("Sprintln(\"%v\") != '%v'", test.want, test.got)
		}
	}
}

func TestSprintf(t *testing.T) {
	if Sprintf("<b>%v</b>", "BOLD") != "\033[1mBOLD\033[0m\033[0m" {
		t.Errorf("Sprintf(\"<b>%v</b>\") != '%v'", "BOLD", "\033[1mBOLD\033[0m\033[0m")
	}

	if Sprintf("<b>%v <blue>%v</blue></b>", "BOLD", "BLUE") != "\033[1mBOLD \033[34mBLUE\033[0m\033[1m\033[0m\033[0m" {
		t.Errorf("Sprintf(\"<b>%v <blue>%v</blue></b>\") != '%v'", "BOLD", "BLUE", "\033[1mBOLD\033[0m\033[0m")
	}
}

func TestFprint(t *testing.T) {
	someVar := "SOME TEXT"
	var tests = []struct {
		want string
		got  string
	}{
		{"HELLO", "HELLO\033[0m"},
		{"<b>BOLD</b>", "\033[1mBOLD\033[0m\033[0m"},
		{"<b>" + someVar + "</b>", "\033[1m" + someVar + "\033[0m\033[0m"},
		{"<b>BOLD <blue>BLUE</blue></b>", "\033[1mBOLD \033[34mBLUE\033[0m\033[1m\033[0m\033[0m"},
		{"<b>HELLO <blue>BLUE</blue> <bg color=\"green\">TEXT</bg></b>", "\033[1mHELLO \033[34mBLUE\033[0m\033[1m \033[42mTEXT\033[0m\033[1m\033[0m\033[0m"},
	}

	for _, test := range tests {
		buf := new(bytes.Buffer)
		Fprint(buf, test.want)
		if buf.String() != test.got {
			t.Errorf("Fprint(buf, \"%v\") != '%v'", test.want, test.got)
		}
	}
}

func TestFprintln(t *testing.T) {
	someVar := "SOME TEXT"
	var tests = []struct {
		want string
		got  string
	}{
		{"HELLO", "HELLO\033[0m\n"},
		{"<b>BOLD</b>", "\033[1mBOLD\033[0m\033[0m\n"},
		{"<b>" + someVar + "</b>", "\033[1m" + someVar + "\033[0m\033[0m\n"},
		{"<b>BOLD <blue>BLUE</blue></b>", "\033[1mBOLD \033[34mBLUE\033[0m\033[1m\033[0m\033[0m\n"},
		{"<b>HELLO <blue>BLUE</blue> <bg color=\"green\">TEXT</bg></b>", "\033[1mHELLO \033[34mBLUE\033[0m\033[1m \033[42mTEXT\033[0m\033[1m\033[0m\033[0m\n"},
	}

	for _, test := range tests {
		buf := new(bytes.Buffer)
		Fprintln(buf, test.want)
		if buf.String() != test.got {
			t.Errorf("Fprintln(buf, \"%v\") != '%v'", test.want, test.got)
		}
	}
}

func TestFprintf(t *testing.T) {
	buf := new(bytes.Buffer)
	Fprintf(buf, "<b>%v</b>", "BOLD")
	if buf.String() != "\033[1mBOLD\033[0m\033[0m" {
		t.Errorf("Fprintf(buf, \"<b>%v</b>\") != '%v'", "BOLD", "\033[1mBOLD\033[0m\033[0m")
	}

	buf2 := new(bytes.Buffer)
	Fprintf(buf2, "<b>%v <blue>%v</blue></b>", "BOLD", "BLUE")
	if buf2.String() != "\033[1mBOLD \033[34mBLUE\033[0m\033[1m\033[0m\033[0m" {
		t.Errorf("Fprintf(buf2, \"<b>%v <blue>%v</blue></b>\") != '%v'", "BOLD", "BLUE", "\033[1mBOLD\033[0m\033[0m")
	}
}
