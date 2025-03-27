package cats

import "io/ioutil"
import "log"
import "strings"

// Fmt returns a formatted string for the provided category.
func Fmt(cat string) string {
	return "\u23fa " + cat
}

// Find returns a slice of strings representing the categories in the provided directory.
func Find(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var cats []string
	for _, file := range files {
		if file.IsDir() {
			// if hidden skip
			if strings.HasPrefix(file.Name(), ".") {
				continue
			}
			cats = append(cats, file.Name())
		}
	}

	return cats
}
