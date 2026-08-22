package internal

type Method int

const (
	GET    Method = iota
	POST          = iota
	PUT           = iota
	DELETE        = iota
)

type Response struct {
	Status        string
	Body          string
	Response_time int64
}

type Request struct {
	Url    string
	Method Method
	Response
}
