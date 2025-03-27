// Package main provides the entry point for ZT.
package main

import "fmt"
import "io/ioutil"
import "log"

func main() {
	cats := GetCategories(".")
	for _, cat := range cats {
		fmt.Println(cat)
	}
}

// GetCategories returns a slice of strings representing the categories in the provided directory.
func GetCategories(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var cats []string
	for _, file := range files {
		if file.IsDir() {
			cats = append(cats, file.Name())
		}
	}

	return cats
}
