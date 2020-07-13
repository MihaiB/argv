package main

import (
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
