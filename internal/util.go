package internal

import (
	"fmt"
	"log"
	"os"
)

func GetOptions() map[string]string {
	options := make(map[string]string)
	lastArgErr := func(i int) {
		if i == len(os.Args)-1 {
			log.Fatal("Invalid Args")
		}
	}
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-csv":
			lastArgErr(i)
			options["-csv"] = os.Args[i+1]
			i++
		case "-json":
			lastArgErr(i)
			options["-json"] = os.Args[i+1]
			i++
		case "-txt":
			lastArgErr(i)
			options["-txt"] = os.Args[i+1]
			i++
		case "-f":
			lastArgErr(i)
			options["-f"] = os.Args[i+1]
			i++
		case "-c":
			lastArgErr(i)
			options["-c"] = os.Args[i+1]
		}

	}

	return options
}

func PrintResult(result []map[string]string) {
	for _, it := range result {
		fmt.Printf("%s %s %s\n", it["url"], it["status"], it["response_time"])
	}
}
