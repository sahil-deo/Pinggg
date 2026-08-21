package internal

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"
)

func GetCallFunction(wg *sync.WaitGroup, maxRoutines *chan struct{}) func(string, *[]map[string]string, time.Duration) {

	client := &http.Client{}

	return func(url string, responses *[]map[string]string, timeout time.Duration) {

		defer wg.Done()
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			log.Fatalf("error %s", err.Error())
		}

		start := time.Now()
		res, err := client.Do(request)
		duration := time.Since(start).String()

		if err != nil {
			log.Printf("request error %s", err.Error())
			response := map[string]string{
				"url":           url,
				"status":        "time limit execeed",
				"response_time": timeout.String(),
			}
			*responses = append(*responses, response)
		} else {
			response := map[string]string{
				"url":           url,
				"status":        res.Status,
				"response_time": string(duration),
			}
			*responses = append(*responses, response)
		}

		<-*maxRoutines
	}
}
