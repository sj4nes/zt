// Package main provides the entry point for ZT.
package main

import "fmt"
import "zt/cats"

func main() {
	cs := cats.Find(".")
	for _, cat := range cs {
		fmt.Println(cats.Fmt(cat))
	}
}
