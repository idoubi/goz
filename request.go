package goz

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

// Request object
type Request struct {
	opts Options
	cli  *http.Client
	req  *http.Request
	body io.Reader
}

// Get send get request
func (r *Request) Get(uri string, opts ...Options) (*Response, error) {
	return r.Request("GET", uri, opts...)
}

// Get method  download files
func (r *Request) Down(resource_url string, sava_path string, opts ...Options) bool {
	uri, err := url.ParseRequestURI(resource_url)
	if err != nil {
		log.Panic("网址无法访问")
	}

	if resp, err := r.Request("GET", resource_url, opts...); err == nil {
		filename := path.Base(uri.Path)
		if resp.GetContentLength() > 0 {
			body := resp.GetBody()
			return r.saveFile(body, sava_path+filename)
		} else {
			log.Panic("被下载的文件内容为空")
		}
	} else {
		log.Panic(err.Error())
	}
	return false
}

func (r *Request) saveFile(body io.ReadCloser, file_name string) bool {
	var is_occur_error bool
	defer body.Close()
	reader := bufio.NewReaderSize(body, 1024*50) //相当于一个临时缓冲区(设置为可以单次存储5M的文件)，每次读取以后就把原始数据重新加载一份，等待下一次读取
	file, err := os.OpenFile(file_name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Panic("创建镜像文件失败，无法进行后续的写入操作" + err.Error())
		is_occur_error = true
	}
	writer := bufio.NewWriter(file)
	buff := make([]byte, 50*1024)

	for {
		curr_read_size, reader_err := reader.Read(buff)
		if curr_read_size > 0 {
			write_size, write_err := writer.Write(buff[0:curr_read_size])
			if write_err != nil {
				log.Panic("写入发生错误"+write_err.Error(), "写入长度：", write_size)
				is_occur_error = true
				break
			}
		}
		// 读取结束
		if reader_err == io.EOF {
			writer.Flush()
			break
		}
	}
	// 如果没有发生错误，就返回 true
	if is_occur_error == false {
		return true
	} else {
		return false
	}

}

// Post send post request
func (r *Request) Post(uri string, opts ...Options) (*Response, error) {
	return r.Request("POST", uri, opts...)
}

// Put send put request
func (r *Request) Put(uri string, opts ...Options) (*Response, error) {
	return r.Request("PUT", uri, opts...)
}

// Patch send patch request
func (r *Request) Patch(uri string, opts ...Options) (*Response, error) {
	return r.Request("PATCH", uri, opts...)
}

// Delete send delete request
func (r *Request) Delete(uri string, opts ...Options) (*Response, error) {
	return r.Request("DELETE", uri, opts...)
}

// Options send options request
func (r *Request) Options(uri string, opts ...Options) (*Response, error) {
	return r.Request("OPTIONS", uri, opts...)
}

// Request send request
func (r *Request) Request(method, uri string, opts ...Options) (*Response, error) {
	if len(opts) > 0 {
		r.opts = mergeHeaders(defaultHeader(), opts...)
	}
	switch method {
	case http.MethodGet, http.MethodDelete:
		req, err := http.NewRequest(method, uri, nil)
		if err != nil {
			return nil, err
		}

		r.req = req
	case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodOptions:
		// parse body
		r.parseBody()

		req, err := http.NewRequest(method, uri, r.body)
		if err != nil {
			return nil, err
		}

		r.req = req
	default:
		return nil, errors.New("invalid request method")
	}
	//r.opts.Headers["Host"]=make()
	r.opts.Headers["Host"] = fmt.Sprintf("%v", r.req.Host)
	// parseOptions
	r.parseOptions()

	// parseClient
	r.parseClient()

	// parse query
	r.parseQuery()

	// parse headers
	r.parseHeaders()

	// parse cookies
	r.parseCookies()

	_resp, err := r.cli.Do(r.req)
	resp := &Response{
		resp: _resp,
		req:  r.req,
		err:  err,
	}

	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (r *Request) parseOptions() {
	// default timeout 30s
	if r.opts.Timeout == 0 {
		r.opts.Timeout = 30
	}
	r.opts.timeout = time.Duration(r.opts.Timeout*1000) * time.Millisecond
}

func (r *Request) parseClient() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if r.opts.Proxy != "" {
		proxy, err := url.Parse(r.opts.Proxy)
		if err == nil {
			tr.Proxy = http.ProxyURL(proxy)
		}
	}

	r.cli = &http.Client{
		Timeout:   r.opts.timeout,
		Transport: tr,
	}
}

func (r *Request) parseQuery() {
	switch r.opts.Query.(type) {
	case string:
		str := r.opts.Query.(string)
		r.req.URL.RawQuery = str
	case map[string]interface{}:
		q := r.req.URL.Query()
		for k, v := range r.opts.Query.(map[string]interface{}) {
			if vv, ok := v.([]string); ok {
				for _, vvv := range vv {
					q.Add(k, vvv)
				}
				continue
			}
			vv := fmt.Sprintf("%v", v)
			q.Set(k, vv)
			//if vv, ok := v.(string); ok {
			//	q.Set(k, vv)
			//	continue
			//}

		}
		r.req.URL.RawQuery = q.Encode()
	}
}

func (r *Request) parseCookies() {
	switch r.opts.Cookies.(type) {
	case string:
		cookies := r.opts.Cookies.(string)
		r.req.Header.Add("Cookie", cookies)
	case map[string]string:
		cookies := r.opts.Cookies.(map[string]string)
		for k, v := range cookies {
			r.req.AddCookie(&http.Cookie{
				Name:  k,
				Value: v,
			})
		}
	case []*http.Cookie:
		cookies := r.opts.Cookies.([]*http.Cookie)
		for _, cookie := range cookies {
			r.req.AddCookie(cookie)
		}
	}
}

func (r *Request) parseHeaders() {
	if r.opts.Headers != nil {
		for k, v := range r.opts.Headers {
			if vv, ok := v.([]string); ok {
				for _, vvv := range vv {
					r.req.Header.Add(k, vvv)
				}
				continue
			}
			vv := fmt.Sprintf("%v", v)
			r.req.Header.Set(k, vv)
		}
	}
}

func (r *Request) parseBody() {
	// application/x-www-form-urlencoded
	if r.opts.FormParams != nil {
		values := url.Values{}
		for k, v := range r.opts.FormParams {
			if vv, ok := v.([]string); ok {
				for _, vvv := range vv {
					values.Add(k, vvv)
				}
				continue
			}
			vv := fmt.Sprintf("%v", v)
			values.Set(k, vv)
			//if vv, ok := v.(string); ok {
			//	values.Set(k, vv)
			//}

		}
		r.body = strings.NewReader(values.Encode())

		return
	}

	// application/json
	if r.opts.JSON != nil {
		b, err := json.Marshal(r.opts.JSON)
		if err == nil {
			r.body = bytes.NewReader(b)

			return
		}
	}

	return
}
