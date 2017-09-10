/**
 * golang版本curl请求库
 * @author mike <mikemintang@126.com>
 * @blog http://idoubi.cc
 */

package curl

import (
	// "fmt"
	"io/ioutil"
	"net/http"
	"strings"
	// "encoding/json"
	// "strconv"
)

const Version = "0.1.0"

// 请求对象
type Request struct {
	Url      string            // 请求url
	Method   string            // 请求方式
	Queries  map[string]string // url请求参数
	PostData map[string]string // post的数据
	Headers  map[string]string // 请求头
	Cookies  map[string]string // 请求cookies
}

// 响应对象
type Response struct {
	HttpStatusCode int               // http响应状态码
	Headers        map[string]string // 响应头
	Body           string            // 响应内容
}

// curl实例
type Client struct {
	hCli  *http.Client   // http实例
	hReq  *http.Request  // http请求
	hResp *http.Response // http响应
}

var (
	Cli *Client
	Req *Request
	Res *Response
)

// 包初始化
func init() {
	Cli = &Client{}
	Req = &Request{}
	Res = &Response{}
}

// 设置请求url
func (c *Client) SetUrl(url string) *Client {
	Req.Url = url
	return c
}

// 设置查询参数
func (c *Client) SetQueries(queries map[string]string) *Client {
	Req.Queries = queries
	return c
}

// 设置post数据
func (c *Client) SetPostData(postData map[string]string) *Client {
	Req.PostData = postData
	return c
}

// 设置请求头
func (c *Client) SetHeaders(headers map[string]string) *Client {
	Req.Headers = headers
	return c
}

// 设置请求cookies
func (c *Client) SetCookies(cookies map[string]string) *Client {
	Req.Cookies = cookies
	return c
}

// 发起get请求
func (c *Client) Get() (*Response, error) {
	Req.Method = http.MethodGet
	return c.Send()
}

// 发起post请求
func (c *Client) Post() (*Response, error) {
	Req.Method = http.MethodPost
	return c.Send()
}

// 通用发起请求逻辑
func (c *Client) Send() (*Response, error) {
	payload := ""

	// 设置post数据
	if postData := Req.PostData; true {
		var req http.Request
		req.ParseForm()
		for k, v := range postData {
			req.Form.Add(k, v)
		}
		payload = strings.TrimSpace(req.Form.Encode())
	}

	// 创建http请求对象
	if req, err := http.NewRequest(Req.Method, Req.Url, strings.NewReader(payload)); err != nil {
		return Res, err
	} else {
		c.hReq = req
	}

	c.hCli = &http.Client{}

	// 设置请求头
	Req.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	if headers := Req.Headers; true {
		for k, v := range headers {
			c.hReq.Header.Set(k, v)
		}
	}

	// 设置请求cookies
	if cookies := Req.Cookies; true {
		for k, v := range cookies {
			cookie := &http.Cookie{
				Name:  k,
				Value: v,
			}
			c.hReq.AddCookie(cookie)
		}
	}

	// 设置请求参数
	if queries := Req.Queries; true {
		q := c.hReq.URL.Query()
		for k, v := range queries {
			q.Add(k, v)
		}
		c.hReq.URL.RawQuery = q.Encode()
	}

	// 发起请求
	if resp, err := c.hCli.Do(c.hReq); err != nil {
		return Res, err
	} else {
		c.hResp = resp
	}

	// 设置响应对象
	Res.HttpStatusCode = c.hResp.StatusCode
	Res.Headers = parseHeaders(c.hResp.Header)

	defer c.hResp.Body.Close()

	// 解析响应内容
	if body, err := ioutil.ReadAll(c.hResp.Body); err != nil {
		return Res, err
	} else {
		Res.Body = string(body)
	}

	return Res, nil
}

// 闭包形式发起get请求
func Get(url string) (*Response, error) {
	Req.Url = url
	Req.Method = http.MethodGet
	return Cli.Send()
}

// 解析请求/响应headers
func parseHeaders(hHeaders http.Header) map[string]string {
	headers := map[string]string{}
	for k, v := range hHeaders {
		headers[k] = v[0]
	}
	return headers
}
