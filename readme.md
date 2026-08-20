## Pinggg

Pinggg is a concurrent CLI tool written in Golang for checking url health. It returns url status code and response time. 

It takes in a file consisting a list of urls as inputs and outputs the responses either in the terminal or a specified output file. 

Output file types supported: 
- .json
- .csv
- .txt

#### Usage

``` bash
go run . -f test/test -c
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
|Option             |Description |
| :---              | :----
| `-f filepath`     | Input file path   
| `-c`              | Use go routines for simulteneously make http calls
| `-json filepath`  | Output to .json file 
| `-csv filepath`   | Output to .csv file 
| `-txt filepath`   | Output to .txt file 

#### Future Scope

- Max go routines allowed along with `-c` flag (right now, no limit).
- Directly take url input from CLI command (right now, input only supported via input file).
- Add request timeout for http requests, also configurable via CLI options `-t 5 (in seconds)`


<p align="center">
  <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" alt="Go">
</p>