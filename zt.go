// Package main provides the entry point for ZT.
package main

import "fmt"
import "zt/cats"

func main() {
	cats.Define(".", "foo")
	cs := cats.Find(".")
	for _, cat := range cs {
		fmt.Println(cats.Fmt(cat))
	}
	cats.Undefine(".", "foo")
}
