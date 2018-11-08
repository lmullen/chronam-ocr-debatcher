[![Build Status](https://travis-ci.org/public-bible/chronam-ocr-debatcher.svg?branch=master)](https://travis-ci.org/public-bible/chronam-ocr-debatcher)

# Chronicling America OCR debatcher

This program looks for a directory containing `.tar.bz2` batches of OCR files from the *Chronicling America* [bulk data downloads](https://chroniclingamerica.loc.gov/about/api/#bulk-data). It converts each batch into a CSV file which you can load into a database or do whatever you like with.

Usage:

```
./chronam-ocr-debatcher path/to/batches
```

