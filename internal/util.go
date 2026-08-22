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
	flag.String("t", "10000", "timeout(ms): max time to wait for a request")
}

func PrintResult(result *[]Request) {
	for _, it := range *result {
		fmt.Printf("%s %s %d\n", it.Url, it.Status, it.Response_time)
	}
}
