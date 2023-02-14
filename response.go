package goz

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/launchdarkly/eventsource"
	"github.com/tidwall/gjson"
)

// Response response object
type Response struct {
	resp   *http.Response
	req    *http.Request
	body   []byte
	stream chan []byte
	err    error
}

// ResponseBody response body
type ResponseBody []byte

// String fmt outout
func (r ResponseBody) String() string {
	return string(r)
}

// Read get slice of response body
func (r ResponseBody) Read(length int) []byte {
	if length > len(r) {
		length = len(r)
	}

	return r[:length]
}

// GetContents format response body as string
func (r ResponseBody) GetContents() string {
	return string(r)
}

// GetRequest get request object
func (r *Response) GetRequest() *http.Request {
	return r.req
}

// GetBody parse response body
func (r *Response) GetBody() (ResponseBody, error) {
	return ResponseBody(r.body), r.err
}

// GetParsedBody parse response body with gjson
func (r *Response) GetParsedBody() (*gjson.Result, error) {
	pb := gjson.ParseBytes(r.body)

	return &pb, nil
}

// GetStatusCode get response status code
func (r *Response) GetStatusCode() int {
	return r.resp.StatusCode
}

// GetReasonPhrase get response reason phrase
func (r *Response) GetReasonPhrase() string {
	status := r.resp.Status
	arr := strings.Split(status, " ")

	return arr[1]
}

// IsTimeout get if request is timeout
func (r *Response) IsTimeout() bool {
	if r.err == nil {
		return false
	}
	netErr, ok := r.err.(net.Error)
	if !ok {
		return false
	}
	if netErr.Timeout() {
		return true
	}

	return false
}

// GetHeaders get response headers
func (r *Response) GetHeaders() map[string][]string {
	return r.resp.Header
}

// GetHeader get response header
func (r *Response) GetHeader(name string) []string {
	headers := r.GetHeaders()
	for k, v := range headers {
		if strings.ToLower(name) == strings.ToLower(k) {
			return v
		}
	}

	return nil
}

// GetHeaderLine get a single response header
func (r *Response) GetHeaderLine(name string) string {
	header := r.GetHeader(name)
	if len(header) > 0 {
		return header[0]
	}

	return ""
}

// HasHeader get if header exsits in response headers
func (r *Response) HasHeader(name string) bool {
	headers := r.GetHeaders()
	for k := range headers {
		if strings.ToLower(name) == strings.ToLower(k) {
			return true
		}
	}

	return false
}

// Err: return response error
func (r *Response) Err() error {
	return r.err
}

// Stream: return response stream
func (r *Response) Stream() chan []byte {
	return r.stream
}

// parse response stream
func (r *Response) parseSteam() {
	r.stream = make(chan []byte)
	decoder := eventsource.NewDecoder(r.resp.Body)

	go func() {
		defer r.resp.Body.Close()
		defer close(r.stream)

		for {
			event, err := decoder.Decode()
			if err != nil {
				r.err = fmt.Errorf("decode data failed: %v", err)
				return
			}

			data := event.Data()
			if data == "" || data == "[DONE]" {
				// read data finished, success return
				return
			}

			r.stream <- []byte(data)
		}
	}()
}
