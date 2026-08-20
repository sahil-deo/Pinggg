package main

import (
	"fmt"
	"log"
	"sync"
	"time"
	"url-health/internal"
)

func main() {

	options := internal.GetOptions()

	filepath, ok := options["-f"]
	if !ok {
		log.Fatal("No filepath provided")
	}

	urls := internal.GetUrlList(filepath)

	var wg sync.WaitGroup
	get := internal.GetCallFunction(&wg)
	testStart := time.Now()

	responses := []map[string]string{}

	for _, url := range urls {
		wg.Add(1)

		if _, ok := options["-c"]; ok {
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
