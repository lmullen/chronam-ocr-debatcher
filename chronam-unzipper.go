// This utility converts Chronicling America OCR batches into CSVs of the OCR
// text. It reads in a directory of Chronicling America OCR batches which are
// stored as .tar.bz2 files, which in turn contain directories of text files
// (which we care about) and XML files (which we don't). The path to the files
// comprise (with modification) an ID for that page on Chronicling America. This
// utility reads in each batch, extracts the page text, and writes each of them
// as a CSV file with a column for the batch ID, page ID, and text.
package main

import (
	"log"
	"os"
	"sync"
)

func main() {
	// Get the path to the directory with the batches
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalln("Provide one argument: the path to the directory with the batches.")
	}
	dir := args[0]

	batches, err := listFiles(dir, ".bz2")
	if err != nil {
		log.Fatalln(err)
	}

	// Process each of the batches concurrently
	var wg sync.WaitGroup
	wg.Add(len(batches))
	for _, batch := range batches {
		// go processOcrBatch(batch, &wg)
		processOcrBatch(batch, &wg)
	}
	wg.Wait()
}
