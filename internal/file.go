package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func GetRequestList(filepath string) []Request {
	file, _ := os.Open(filepath)
	defer file.Close()

	r := bufio.NewScanner(file)

	requests := []Request{}

	for r.Scan() {
		line := r.Text()
		requests = append(requests, Request{Url: line})
	}

	err := r.Err()
	if err != nil {
		log.Fatalf("scan error %s\n", err.Error())
	}

	return requests
}

func Write(requests *[]Request, outtype string, outpath string) {
	outpath = getOutFilepath(outtype, outpath)
	switch outtype {
	case "json":
		writeJson(requests, outpath)
	case "csv":
		writeCsv(requests, outpath)
	case "txt":
		writeTxt(requests, outpath)
	default:
		PrintResult(requests)
	}
}

func getOutFilepath(outtype string, outpath string) string {
	if outpath == "" {
		outpath = "out." + outtype
	}
	return outpath
}

func writeCsv(result *[]Request, outpath string) {
	file, err := os.Create(outpath)
	if err != nil {
		log.Fatalf("err %s", err.Error())
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("url,status,response_time\n")
	for _, it := range *result {
		line := fmt.Sprintf("%s,%s,%d\n", it.Url, it.Status, it.Response_time)
		writer.WriteString(line)
	}
	writer.Flush()
}

func writeJson(result *[]Request, outpath string) {
	json, err := json.Marshal(*result)
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

func writeTxt(result *[]Request, outpath string) {
	file, err := os.Create(outpath)
	if err != nil {
		log.Fatalf("err %s", err.Error())
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, it := range *result {
		line := fmt.Sprintf("%s,%s,%d\n", it.Url, it.Status, it.Response_time)
		writer.WriteString(line)
	}
	writer.Flush()
}
