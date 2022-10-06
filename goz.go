package goz

// NewClient new request object
func NewClient(opts ...Options) *Request {
	req := &Request{}

	opts0 := Options{}
	if len(opts) > 0 {
		opts0 = opts[0]
	}

	req.SetOptions(opts0)

	return req
}

// Get send get request
func Get(uri string, opts ...Options) (*Response, error) {
	r := NewClient()
	return r.Request("GET", uri, opts...)
}

// Post send post request
func Post(uri string, opts ...Options) (*Response, error) {
	r := NewClient()
	return r.Request("POST", uri, opts...)
}

// Put send put request
func Put(uri string, opts ...Options) (*Response, error) {
	r := NewClient()
	return r.Request("PUT", uri, opts...)
}

// Patch send patch request
func Patch(uri string, opts ...Options) (*Response, error) {
	r := NewClient()
	return r.Request("PATCH", uri, opts...)
}

// Delete send delete request
func Delete(uri string, opts ...Options) (*Response, error) {
	r := NewClient()
	return r.Request("DELETE", uri, opts...)
}
