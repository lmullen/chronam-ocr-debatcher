[![Build Status](https://travis-ci.org/public-bible/chronam-ocr-debatcher.svg?branch=master)](https://travis-ci.org/public-bible/chronam-ocr-debatcher)

# Chronicling America OCR debatcher

This program takes paths to `.tar.bz2` batches of OCR files from the
*Chronicling America* [bulk data
downloads](https://chroniclingamerica.loc.gov/about/api/#bulk-data). It converts
each batch into a CSV file, which you can load into a database or do whatever
you like with. It will process the batches concurrently.

Usage:

```
./chronam-ocr-debatcher [--processes=8] <path/to/a/batch.tar.bz2 ...>
```
