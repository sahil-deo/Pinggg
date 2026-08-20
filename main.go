package main

import (
	"fmt"
	"log"
	"sync"
	"time"
	"url-health/internal"
)

var wg sync.WaitGroup

func main() {

	options := internal.GetOptions()

	filepath, ok := options["-f"]
	if !ok {
		log.Fatal("No filepath provided")
	}

	urls := internal.GetUrlList(filepath)

	get := internal.GetCallFunction(&wg)
	testStart := time.Now()

	result := []map[string]string{}

	for _, url := range urls {
		wg.Add(1)

		if _, ok := options["-c"]; ok {
			go get(url, &result)
		} else {
			get(url, &result)
		}
	}
	wg.Wait()
	totalTestTime := time.Since(testStart)
	fmt.Printf("Total test time: %f\n", totalTestTime.Seconds())

	if outpath, ok := options["-csv"]; ok {
		internal.WriteCsv(result, outpath)
	} else if outpath, ok := options["-txt"]; ok {
		internal.WriteTxt(result, outpath)
	} else if outpath, ok := options["-json"]; ok {
		internal.WriteJson(result, outpath)
	} else {
		internal.PrintResult(result)
	}
}
