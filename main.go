package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
	"url-health/internal"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	// options := internal.GetOptions()
	internal.DefineFlags() // declare flags
	flag.Parse()
	options := make(map[string]string) // map to store the flags
	options["c"] = "1"                 // default concurrency
	options["t"] = "10000"             // default timeout (10 seconds = 10000 ms)

	// check flags which are only set
	flag.Visit(func(f *flag.Flag) {
		options[f.Name] = f.Value.String()
	})

	// u = url flag
	// f = filepath flag
	// atleast one is required
	url, uok := options["u"]
	filepath, fok := options["f"]

	if !uok && !fok {
		log.Fatal("No url or filepath provided")
	}

	urls := []string{}

	switch {
	case uok:
		urls = append(urls, url)
	case fok:
		urls = internal.GetUrlList(filepath)
	}

	var wg sync.WaitGroup

	// c = concurrency flag; default is 1
	max, err := strconv.Atoi(options["c"])
	check(err)
	if max < 1 || max > 999 { // no limit
		max = 999 // hardlimit
	}

	timeout_ms, err := strconv.Atoi(options["t"])
	check(err)

	maxRoutines := make(chan struct{}, max) // semaphore

	// get the function to make the get calls
	get := internal.GetCallFunction(&wg, &maxRoutines)
	testStart := time.Now()

	responses := []map[string]string{}

	for _, url := range urls {
		wg.Add(1)
		maxRoutines <- struct{}{}
		go get(url, &responses, time.Duration(timeout_ms)*time.Millisecond)
	}

	wg.Wait()
	totalTestTime := time.Since(testStart)
	fmt.Printf("Total test time: %f\n", totalTestTime.Seconds())

	outtype := options["o"] // output file type
	outpath := options["n"] // output file path

	internal.Write(responses, outtype, outpath)

}
