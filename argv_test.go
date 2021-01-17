package main

import (
	"strings"
	"testing"
)

func TestFmtArg(t *testing.T) {
	for params, want := range map[struct {
		i int
		s string
	}]string{
		{7, ""}:             "7 (0 codepoints, 0 bytes): \"\"",
		{13, "\r\n\" \t"}:   "13 (5 codepoints, 5 bytes): \"\\r\\n\\\" \\t\"",
		{0, "P⇒Q"}:          "0 (3 codepoints, 5 bytes): \"P⇒Q\"",
		{1, "≡\xcc\xffØ\n"}: "1 (5 codepoints, 8 bytes): \"≡\\xcc\\xffØ\\n\"",
	} {
		got := fmtArg(params.i, params.s)
		if got != want {
			t.Errorf("%v: want %#v got %#v", params, want, got)
		}
	}
}

func TestWriteArgvZero(t *testing.T) {
	for _, argv := range [][]string{nil, {}} {
		b := &strings.Builder{}
		writeArgv(argv, b)
		want := "argv length: 0\n"
		got := b.String()

		if got != want {
			t.Errorf("want %#v got %#v", want, got)
		}
	}
}

func TestWriteArgvNonZero(t *testing.T) {
	argv := []string{"prg", "a", "b\n\tc"}
	want := `argv length: 3
0 (3 codepoints, 3 bytes): "prg"
1 (1 codepoints, 1 bytes): "a"
2 (4 codepoints, 4 bytes): "b\n\tc"
`
	b := &strings.Builder{}
	writeArgv(argv, b)
	got := b.String()
	if got != want {
		t.Errorf("want %#v got %#v", want, got)
	}
}
