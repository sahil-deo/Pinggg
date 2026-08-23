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
	ResponseBody  string
	Response_time int64
}

type Request struct {
	Url         string
	Method      Method
	RequestBody string
	Response
}

type RequestBuilder struct {
	request Request
}

func GetRequestBuilder() *RequestBuilder {
	return &RequestBuilder{request: Request{}}
}

func (r *RequestBuilder) SetUrl(url string) *RequestBuilder {
	r.request.Url = url
	return r
}

func (r *RequestBuilder) SetMethod(method Method) *RequestBuilder {
	r.request.Method = method
	return r
}

func (r *RequestBuilder) SetBody(body string) *RequestBuilder {
	r.request.RequestBody = body
	return r
}

func (r *RequestBuilder) Build() Request {
	return r.request
}
