package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Given a directory and an extension to look for, this function finds all of
// the files matching an extension recursively within a directory, and returns a
// slice of paths.
func listFiles(dir string, ext string) ([]string, error) {
	files := make([]string, 0)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		log.Printf("Error walking the path %q: %v\n", dir, err)
		return nil, err
	}
	return files, nil
}

func pathToID(path string) string {
	p := strings.Split(path, "/")
	return fmt.Sprintf("%s/%s-%s-%s/%s/%s/", p[0], p[1], p[2], p[3], p[4], p[5])
}
