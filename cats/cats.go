package cats

import "errors"
import "io/ioutil"
import "log"
import "os"
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

// Define creates a category in the provided directory.
func Define(dir string, cat string) (*string, error) {
	// If the category already exists, do nothing.
	exists, err := os.Stat(dir + "/" + cat)
	if err == nil && exists.IsDir() {
		return nil, errors.New("category exists")
	}
	// If the category does not exist, create it.
	err = os.Mkdir(dir+"/"+cat, 0755)
	if err != nil {
		return nil, err
	}
	return &cat, nil
}

// Undefine deletes a category in the provided directory.
func Undefine(dir string, cat string) error {
	err := os.Remove(dir + "/" + cat + "/=A")
	if err != nil {
		return err
	}
	return nil
}

// Arrow connects two categories by creating a subdirectory in the source category.
func Arrow(from string, to string) error {
	err := os.MkdirAll(from+"/"+to+"/=A", 0755)
	if err != nil {
		return err
	}
	return nil
}

