// This utility converts Chronicling America OCR batches into CSVs of the OCR
// text. It takes as its arguments paths to Chronicling America OCR batches
// which are stored as .tar.bz2 files, which in turn contain directories of text
// files (which we care about) and XML files (which we don't). The path to the
// files comprise (with modification) an ID for that page on Chronicling
// America. This utility reads in each batch, extracts the page text, and writes
// each of them as a CSV file with a column for the batch ID, page ID, and text.
// It will process the batches in parallel.
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [--processes=8] <path/to/a/batch.tar.bz2 ...> \n", os.Args[0])
		flag.PrintDefaults()
	}
	processes := flag.Int("processes", 8, "Number of batches to process in parallel")
	flag.Parse()
	batches := flag.Args()

	// Check that we get at least one path
	if len(batches) == 0 {
		fmt.Println("Error: provide paths to Chronicling America .tar.bz2 batches.")
		flag.Usage()
		os.Exit(1)
	}

	// Check that the paths point to actual batches
	err := checkPathsToBatches(batches)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Process the batches in parallel
	_ = boundedParallelProcess(batches, *processes)

}
