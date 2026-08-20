package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"url-health/internal"
)

var wg sync.WaitGroup

func main() {

	if len(os.Args) < 2 {
		log.Fatal("No filepath provided")
	}

	filepath := os.Args[1]

	con := ""
	if len(os.Args) > 2 {
		con = os.Args[2]
	}

	urls := internal.GetUrlList(filepath)

	get := internal.GetCallFunction(&wg)
	testStart := time.Now()

	for _, url := range urls {
		wg.Add(1)

		if con == "-c" {
			go get(url)
		} else {
			get(url)
		}

	}
	wg.Wait()

	totalTestTime := time.Since(testStart)
	fmt.Printf("Total test time: %f\n", totalTestTime.Seconds())
}
