package tz4

import "testing"

func TestLabelFmt(t *testing.T) {
	if l.Fmt("name") != "\U0001F3F7 name" {
		t.Fail()
	}
}
