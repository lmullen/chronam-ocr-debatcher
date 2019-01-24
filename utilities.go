package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Check that there is no problem with reading the files passed in
func checkPathsToBatches(paths []string) error {
	for _, path := range paths {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			return err
		} else if filepath.Ext(path) != ".bz2" {
			err := fmt.Errorf("Error: %s is not a .tar.bz2 batch from Chronicling America", path)
			return err
		} else if err != nil {
			return err
		}

	}
	return nil
}

// Convert the path within the .tar.bz2 file to a page ID used on the Chronicling
// America website.
func pathToID(path string) string {
	p := strings.Split(path, "/")
	return fmt.Sprintf("%s/%s-%s-%s/%s/%s/", p[0], p[1], p[2], p[3], p[4], p[5])
}
