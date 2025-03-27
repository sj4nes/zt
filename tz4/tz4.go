// Package tz4 is a utility for .tar.lz4 file management.

package tz4

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)
import "errors"

// FILE_EXT is the default file extension for .tar.lz4 files that is a three character.
const FILE_EXT = ".tz4"

// Fmt returns a formatted string for the provided tz4 name.
func Fmt(name string) string {
	return "\u2894 " + name
}

// Find returns a slice of strings representing the .tz4 files in the provided directory.
func Find(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var arcs []string
	for _, file := range files {
		if file.IsDir() {
			// if hidden skip
			if strings.HasSuffix(file.Name(), FILE_EXT) {
				continue
			}
			arcs = append(arcs, file.Name())
		}
	}

	return arcs
}

// Create a .tz4 file in the provided directory from the given items.
func Create(dir string, name string, items []string) error {
	// If the .tz4 file already exists, do nothing.
	exists, err := os.Stat(dir + "/" + name + FILE_EXT)
	if err == nil && exists != nil {
		return errors.New("archive exists")
	}
	// If the .tz4 file does not exist, create it.
	err = os.Mkdir(dir+"/"+name+FILE_EXT, 0755)
	if err != nil {
		return err
	}
	return nil
}
