package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func cleanDir(dir string, exclude []string) {
	files, err := filepath.Glob(path.Join(dir, "*"))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if contains(exclude, path.Base(file)) {
			continue
		}
		os.Remove(file)
	}
}

func main() {
	// `ripplet_lexer_base.go` is our custom base definition, keep it
	cleanDir("./internal/grammar", []string{"ripplet_lexer_base.go"})
}
