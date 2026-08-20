package internal

import (
	"bufio"
	"log"
	"os"
)

func GetUrlList(filepath string) []string {
	file, _ := os.Open(filepath)
	defer file.Close()

	r := bufio.NewScanner(file)

	urls := []string{}

	for r.Scan() {
		line := r.Text()
		urls = append(urls, line)
	}

	err := r.Err()
	if err != nil {
		log.Fatalf("scan error %s\n", err.Error())
	}

	return urls
}
