package goz

import "time"

// Options object
type Options struct {
	Headers    map[string]interface{}
	BaseURI    string
	Query      interface{}
	FormParams map[string]interface{}
	JSON       interface{}
	Timeout    float32
	timeout    time.Duration
	Cookies    interface{}
	Proxy      string
}
