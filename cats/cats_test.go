package cats

import "testing"

func TestFmt(t *testing.T) {
	if Fmt("foo") != "\u23fa foo" {
		t.Fail()
	}
}

func TestCategoryFmt(t *testing.T) {
	c := Category{"foo"}
	if c.Fmt() != "\u23fa foo" {
		t.Fail()
	}
}
