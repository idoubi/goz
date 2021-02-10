package goz

import "time"

// Options object
type Options struct {
	Debug      bool
	BaseURI    string
	Timeout    float32
	timeout    time.Duration
	Query      interface{}
	Headers    map[string]interface{}
	Cookies    interface{}
	FormParams map[string]interface{}
	JSON       interface{}
	XML        interface{}
	Proxy      string
}
