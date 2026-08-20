package internal

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func GetCallFunction(wg *sync.WaitGroup) func(url string) {

	client := &http.Client{}

	return func(url string) {

		defer wg.Done()

		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatalf("error %s", err.Error())
		}

		start := time.Now()
		res, err := client.Do(request)
		duration := time.Since(start).Seconds()

		if err != nil {
			log.Fatalf("error %s", err.Error())
		}

		fmt.Printf("%s %s %f\n", url, res.Status, duration)
	}
}
