BINARY_NAME=chronam-ocr-debatcher

build : clean
	go mod tidy
	go build

clean :
	rm -f $(BINARY_NAME)
	rm -f test-data/*.csv

test : 
	./$(BINARY_NAME) test-data/*.tar.bz2

release:
	goreleaser release --skip-publish --rm-dist
