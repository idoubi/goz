package goz

import (
	"fmt"
	"github.com/axgle/mahonia"
)

// NewClient new request object
func NewClient(opts ...Options) *Request {
	req := &Request{}

	if len(opts) > 0 {
		req.opts = mergeHeaders(defaultHeader(), opts[0])
	}

	return req
}

// 合并用户提供的header头字段信息，用户提供的header头优先于默认头字段信息
func mergeHeaders(default_headers Options, options ...Options) Options {
	if len(options) == 0 {
		return defaultHeader()
	} else {
		for key, value := range default_headers.Headers {
			if options[0].Headers != nil {
				if _, exists := options[0].Headers[key]; !exists {
					options[0].Headers[key] = fmt.Sprintf("%v", value)
				}
			} else {
				options[0].Headers = make(map[string]interface{}, 1)
				options[0].Headers[key] = fmt.Sprintf("%v", value)
			}
		}
		return options[0]
	}
}

// 默认设置headers头信息，尽可能伪装成为真实的浏览器
func defaultHeader() Options {
	headers := Options{
		Headers: map[string]interface{}{
			"User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.81 Safari/537.36 SE 2.X MetaSr 1.0",
			"Accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
			//	特别提醒：真实的浏览器该值为 Accept-Encoding: gzip, deflate，表示浏览器接受压缩后的二进制，浏览器端再解析为html展示，
			//	但是HttpClient解析就麻烦了，所以必须为空或者不设置该值，接受原始数据。否则很容易出现乱码
			"Accept-Encoding":           "",
			"Accept-Language":           "zh-CN,zh;q=0.9",
			"Upgrade-Insecure-Requests": "1",
			"Connection":                "keep-alive",
			"Cache-Control":             "max-age=0",
		},
	}
	return headers
}

// 编码转换，中文编码转码
func simpleChinese2Utf8(v_bytes []byte) string {
	return mahonia.NewDecoder("GB18030").ConvertString(string(v_bytes))
}
