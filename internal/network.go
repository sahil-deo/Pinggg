package internal

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"
)

func GetCallFunction(wg *sync.WaitGroup, maxRoutines *chan struct{}) func(*Request, time.Duration) {

	client := &http.Client{}

	return func(request *Request, timeout time.Duration) {

		defer wg.Done()
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		url := request.Url
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			log.Fatalf("error %s", err.Error())
		}

		start := time.Now()
		res, err := client.Do(req)
		duration := time.Since(start)

		if err != nil {
			log.Printf("request error %s", err.Error())
			request.Response_time = timeout.Milliseconds()
			request.Status = "time limit exceed"

		} else {
			request.Status = res.Status
			request.Response_time = duration.Milliseconds()
		}
		<-*maxRoutines
	}
}
