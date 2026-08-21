package main

import (
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

	options := internal.GetOptions()

	filepath, ok := options["-f"]
	if !ok {
		log.Fatal("No filepath provided")
	}

	urls := internal.GetUrlList(filepath)

	var wg sync.WaitGroup
	max, ok := options["-c"]
	if !ok {
		log.Fatal("No max go routines specified")
	}
	maxI, err := strconv.Atoi(max)
	check(err)
	if maxI < 1 { // no limit
		maxI = 999 // hard limit
	}
	maxRoutines := make(chan struct{}, maxI)
	get := internal.GetCallFunction(&wg, &maxRoutines)
	testStart := time.Now()

	responses := []map[string]string{}

	for _, url := range urls {
		wg.Add(1)

		if _, ok := options["-c"]; ok {
			maxRoutines <- struct{}{}
			go get(url, &responses)
		} else {
			get(url, &responses)
		}
	}
	wg.Wait()
	totalTestTime := time.Since(testStart)
	fmt.Printf("Total test time: %f\n", totalTestTime.Seconds())

	if outpath, ok := options["-csv"]; ok {
		internal.WriteCsv(responses, outpath)
	} else if outpath, ok := options["-txt"]; ok {
		internal.WriteTxt(responses, outpath)
	} else if outpath, ok := options["-json"]; ok {
		internal.WriteJson(responses, outpath)
	} else {
		internal.PrintResult(responses)
	}
}
