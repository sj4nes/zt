// Package main provides the entry point for ZT.
package main

import "fmt"
import "zt/cats"
import . "zt/quality"

func main() {
	Assert(true, "This should not panic in release mode")
	cats.Define(".", "foo")
	cs := cats.Find(".")
	for _, cat := range cs {
		fmt.Println(cats.Fmt(cat))
	}
	cats.Undefine(".", "foo")
}
