package main

type result struct {
	index int
	res   bool
	err   error
}

// Run a processing function on each of the paths. Up to the `concurrencyLimit`
// number of paths will be processed concurrently.
//
// Adapted from: https://gist.github.com/montanaflynn/ea4b92ed640f790c4b9cee36046a5383
func boundedParallelProcess(batches []string, concurrencyLimit int) []bool {

	// this buffered channel will block at the concurrency limit
	semaphoreChan := make(chan struct{}, concurrencyLimit)

	// this channel will not block and collect the results
	resultsChan := make(chan bool)

	// make sure we close these channels when we're done with them
	defer func() {
		close(semaphoreChan)
		close(resultsChan)
	}()

	// keep an index and loop through every batch we will process
	for i, batch := range batches {

		// start a go routine with the index and url in a closure
		go func(i int, batch string) {

			// this sends an empty struct into the semaphoreChan which is basically
			// saying add one to the limit, but when the limit has been reached block
			// until there is room
			semaphoreChan <- struct{}{}

			// process the batch
			res, _ := processOcrBatch(batch)

			// now we can send the result through the resultsChan
			resultsChan <- res

			// once we're done it's we read from the semaphoreChan which has the
			// effect of removing one from the limit and starting another goroutine
			<-semaphoreChan

		}(i, batch)
	}

	// make a slice to hold the results we're expecting
	var results []bool

	// start listening for any results over the resultsChan
	// once we get a result append it to the result slice
	for {
		result := <-resultsChan
		results = append(results, result)

		// if we've reached the expected amount of batches then stop
		if len(results) == len(batches) {
			break
		}
	}

	// now we're done we return the results
	return results

}
