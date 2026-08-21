## Pinggg

Pinggg is a concurrent CLI tool written in Golang for checking url health. It returns url status code and response time. 

It takes in a file consisting a list of urls as inputs and outputs the responses either in the terminal or a specified output file. 

Output file types supported: 
- .json
- .csv
- .txt

#### Usage

``` bash
go run . -f test/test -c 5
```

``` bash
Total test time: 0.390489
https://www.reddit.com 200 OK 88.814625ms
https://www.scalant.in 200 OK 139.250042ms
https://www.google.com 200 OK 217.849333ms
https://www.youtube.com 200 OK 248.239625ms
https://www.linkedin.com 200 OK 390.19525ms
```

#### Options 
|Option               |Usage|Description 
| :---                | :----                 | :----
| `-f [filepath]`     | `-f test.txt`         | Input file path   
| `-u [url]`          | `-u htts://google.com`| Input single url to ping 
| `-c [n]`            | `-c 5`                | Use go routines for simulteneously make http calls, n is max routines allowed, use < 0 for no limits
| `-o [outtype]`      | `-o json`             | Output file type [ json \| csv \| txt ] 
| `-n [filepath]`     | `-n out.json`         | Output file path
| `-h`                | `-h`                  | Print Help

#### Future Scope

- <s>Max go routines allowed along with `-c` flag (right now, no limit).</s>
- <s>Directly take url input from CLI command (right now, input only supported via input file).</s>
- Add request timeout for http requests, also configurable via CLI options `-t 5 (in seconds)`.
- <s>Use os/flag for cli options.</s>


<p align="center">
  <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" alt="Go">
</p>