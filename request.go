/**
 * golang版本的curl请求库
 * Request构造类，用于设置请求参数，发起http请求
 * @author mike <mikemintang@126.com>
 * @blog http://idoubi.cc
 */

package curl

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"time"
)

// Request构造类
type Request struct {
	cli             *http.Client
	req             *http.Request
	Raw             *http.Request
	Method          string
	Url             string
	dialTimeout     time.Duration
	responseTimeOut time.Duration
	Headers         map[string]string
	Cookies         map[string]string
	Queries         map[string]string
	PostData        map[string]interface{}
}

// 创建一个Request实例
func NewRequest() *Request {
	r := &Request{}
	r.dialTimeout = 5
	r.responseTimeOut = 5
	return r
}

// 设置请求方法
func (this *Request) SetMethod(method string) *Request {
	this.Method = method
	return this
}

// 设置请求地址
func (this *Request) SetUrl(url string) *Request {
	this.Url = url
	return this
}

// 设置请求头
func (this *Request) SetHeaders(headers map[string]string) *Request {
	this.Headers = headers
	return this
}

// 将用户自定义请求头添加到http.Request实例上
func (this *Request) setHeaders() error {
	for k, v := range this.Headers {
		this.req.Header.Set(k, v)
	}
	return nil
}

// 设置请求cookies
func (this *Request) SetCookies(cookies map[string]string) *Request {
	this.Cookies = cookies
	return this
}

// 将用户自定义cookies添加到http.Request实例上
func (this *Request) setCookies() error {
	for k, v := range this.Cookies {
		this.req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}
	return nil
}

// 设置url查询参数
func (this *Request) SetQueries(queries map[string]string) *Request {
	this.Queries = queries
	return this
}

// 将用户自定义url查询参数添加到http.Request上
func (this *Request) setQueries() error {
	q := this.req.URL.Query()
	for k, v := range this.Queries {
		q.Add(k, v)
	}
	this.req.URL.RawQuery = q.Encode()
	return nil
}

// 设置post请求的提交数据
func (this *Request) SetPostData(postData map[string]interface{}) *Request {
	this.PostData = postData
	return this
}

// 发起get请求
func (this *Request) Get() (*Response, error) {
	return this.Send(this.Url, http.MethodGet)
}

// 发起Delete请求
func (this *Request) Delete() (*Response, error) {
	return this.Send(this.Url, http.MethodDelete)
}

// 发起Delete请求
func (this *Request) Put() (*Response, error) {
	return this.Send(this.Url, http.MethodPut)
}

// 发起post请求
func (this *Request) Post() (*Response, error) {
	return this.Send(this.Url, http.MethodPost)
}

// 发起put请求
func (this *Request) PUT() (*Response, error) {
	return this.Send(this.Url, http.MethodPut)
}

// 发起put请求
func (this *Request) PATCH() (*Response, error) {
	return this.Send(this.Url, http.MethodPatch)
}

//SetDialTimeOut
func (this *Request) SetDialTimeOut(TimeOutSecond int) {
	this.dialTimeout = time.Duration(TimeOutSecond)
}

//SetResponseTimeOut
func (this *Request) SetResponseTimeOut(TimeOutSecond int) {
	this.responseTimeOut = time.Duration(TimeOutSecond)
}

// 发起请求
func (this *Request) Send(url string, method string) (*Response, error) {
	// 检测请求url是否填了
	if url == "" {
		return nil, errors.New("Lack of request url")
	}
	// 检测请求方式是否填了
	if method == "" {
		return nil, errors.New("Lack of request method")
	}
	// 初始化Response对象
	response := NewResponse()
	// 初始化http.Client对象
	this.cli = &http.Client{
		////////
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*this.dialTimeout)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * this.dialTimeout))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * this.responseTimeOut,
		},
		////////////
	}
	// 加载用户自定义的post数据到http.Request
	var payload io.Reader
	if method == "POST" && this.PostData != nil {
		if jData, err := json.Marshal(this.PostData); err != nil {
			return nil, err
		} else {
			payload = bytes.NewReader(jData)
		}
	} else {
		payload = nil
	}

	if req, err := http.NewRequest(method, url, payload); err != nil {
		return nil, err
	} else {
		this.req = req
	}

	this.setHeaders()
	this.setCookies()
	this.setQueries()

	this.Raw = this.req

	if resp, err := this.cli.Do(this.req); err != nil {
		return nil, err
	} else {
		response.Raw = resp
	}

	defer response.Raw.Body.Close()

	response.parseHeaders()
	response.parseBody()

	return response, nil
}
