package main

import (
	"archive/tar"
	"compress/bzip2"
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// This function reads in a bz2 compressed tar file and writes the contents to a CSV file of the same name.
func processOcrBatch(path string) (bool, error) {

	outPath := strings.Replace(path, ".tar.bz2", ".csv", 1)
	batch := strings.Replace(filepath.Base(path), ".tar.bz2", "", 1)

	log.Printf("Beginning to process batch %s\n", batch)

	// Skip processing if the exported csv file already exists
	if _, err := os.Stat(outPath); !os.IsNotExist(err) {
		log.Printf("Skipped already processed batch %s\n", batch)
		return false, err
	}

	// Set up a csv file to write output
	outFile, err := os.Create(outPath)
	if err != nil {
		log.Println("Cannot create file:", err)
		return false, err
	}
	defer outFile.Close()

	outWriter := csv.NewWriter(outFile)
	defer outWriter.Flush()

	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return false, err
	}
	defer f.Close()

	// Read through the tar file, writing each text file as a record to the csv
	bz2f := bzip2.NewReader(f)
	tarf := tar.NewReader(bz2f)

	for true {
		header, err := tarf.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println(err)
		}

		if header.Typeflag == tar.TypeReg && filepath.Ext(header.Name) == ".txt" {
			id := pathToID(header.Name)
			text, _ := ioutil.ReadAll(tarf)
			outWriter.Write([]string{batch, id, string(text)})
		}

		if err := outWriter.Error(); err != nil {
			log.Fatalln("Error writing csv:", err)
			return false, err
		}

	}

	log.Printf("Finished processing batch %s\n", batch)
	return true, nil

}
