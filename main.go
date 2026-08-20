package main

import (
	"bufio"
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

	options := make(map[string]bool)

	for _, arg := range os.Args[2:] {
		options[arg] = true
	}

	filepath := os.Args[1]

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

	if _, ok := options["-j"]; ok {
		file, err := os.Create("out.csv")
		if err != nil {

		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		writer.WriteString("url,status,response_time\n")
		for _, it := range result {
			line := fmt.Sprintf("%s,%s,%s\n", it["url"], it["status"], it["response_time"])
			writer.WriteString(line)
		}
		writer.Flush()

	} else {
		for _, it := range result {
			fmt.Printf("%s %s %s\n", it["url"], it["status"], it["response_time"])
		}
	}
}
