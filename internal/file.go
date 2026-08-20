package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
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

func WriteCsv(result []map[string]string, outpath string) {
	file, err := os.Create(outpath)
	if err != nil {
		log.Fatalf("err %s", err.Error())
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("url,status,response_time\n")
	for _, it := range result {
		line := fmt.Sprintf("%s,%s,%s\n", it["url"], it["status"], it["response_time"])
		writer.WriteString(line)
	}
	writer.Flush()
}

func WriteJson(result []map[string]string, outpath string) {
	json, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("Err %s", err.Error())
	}
	file, err := os.Create(outpath)
	if err != nil {
		log.Fatalf("err %s", err.Error())
	}
	defer file.Close()
	file.WriteString(string(json))
}

func WriteTxt(result []map[string]string, outpath string) {
	file, err := os.Create(outpath)
	if err != nil {
		log.Fatalf("err %s", err.Error())
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, it := range result {
		line := fmt.Sprintf("%s %s %s\n", it["url"], it["status"], it["response_time"])
		writer.WriteString(line)
	}
	writer.Flush()
}
