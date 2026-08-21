package internal

import (
	"flag"
	"fmt"
)

func DefineFlags() {
	flag.String("f", "", "filepath: input filepath, simple text files only")
	flag.String("o", "", "outtype: output file type [json | csv | txt]; prints to cli if not specified")
	flag.String("n", "", "name: output filepath")
	flag.String("u", "", "url: single url to ping")
	flag.String("c", "1", "concurrency: number of max go routinues allowed; use < 1 for no limits (hard internal limit at 999)")
}

func PrintResult(result []map[string]string) {
	for _, it := range result {
		fmt.Printf("%s %s %s\n", it["url"], it["status"], it["response_time"])
	}
}
