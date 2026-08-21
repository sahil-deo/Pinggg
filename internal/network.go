package internal

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func GetCallFunction(wg *sync.WaitGroup, maxRoutines *chan struct{}) func(string, *[]map[string]string) {

	client := &http.Client{}

	return func(url string, responses *[]map[string]string) {

		defer wg.Done()

		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatalf("error %s", err.Error())
		}

		start := time.Now()
		res, err := client.Do(request)
		duration := time.Since(start).String()

		if err != nil {
			log.Fatalf("error %s", err.Error())
		}
		response := map[string]string{
			"url":           url,
			"status":        res.Status,
			"response_time": string(duration),
		}

		*responses = append(*responses, response)
		<-*maxRoutines
	}
}
