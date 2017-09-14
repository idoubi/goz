package curl

import (
	"io/ioutil"
	"net/http"
)

type Response struct {
	Raw     *http.Response
	Headers map[string]string
	Body    string
}

func NewResponse() *Response {
	return &Response{}
}

func (this *Response) IsOk() bool {
	return this.Raw.StatusCode == 200
}

func (this *Response) parseHeaders() error {
	headers := map[string]string{}
	for k, v := range this.Raw.Header {
		headers[k] = v[0]
	}
	this.Headers = headers
	return nil
}

func (this *Response) parseBody() error {
	if body, err := ioutil.ReadAll(this.Raw.Body); err != nil {
		panic(err)
	} else {
		this.Body = string(body)
	}
	return nil
}
