# goz

- 基于goz改造，感谢goz（https://github.com/idoubi/goz.git）原作者提供了大量的基础代码
- 相比原版变化：
>   1.增加了文件下载功能  
>   2.GetBody() 返回io.ReaderCloser ,而不是原版本中的文本格式数据。
>   3.GetBody() 将专门负责处理流式数据，因此代码逻辑处理完毕，必须使用io.ReaderCloser 接口提供的函数关闭。  
>   4.原版本的GetBody()被现有版本GetContents()代替，由于文本数据,一次性返回，因此不需要手动关闭，程序会自动释放相关io资源。  
>   5.删除、简化了原版本中为处理数据类型转换而定义的ResponseBody,本版本中使用系统系统默认的数据类型转换即可，简单快捷。  
>   6.增强原版本中表单参数只能传递string、[]string的问题，该版本支持数字、文本、[]string等。

## Installation

```
go get -u github.com/qifengzhang007/goz.git
```


## Documentation

API documentation can be found here:
https://github.com/qifengzhang007/goz.git


## Basic Usage

```go
package main

import (
    "github.com/qifengzhang007/goz.git"
)

func main() {
    cli := goz.NewClient()

	resp, err := cli.Get("http://127.0.0.1:8091/get")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T", resp)
	// Output: *goz.Response
}
```

## Query Params

- query map

```go
func ExampleRequest_Get_withQuery_arr() {
	cli := goz.NewClient()

	resp, err := cli.Get("http://127.0.0.1:8091/get-with-query", goz.Options{
		Query: map[string]interface{}{
			"key1": 123,
			"key2": []string{"value21", "value22"},
			"key3": "abc456",
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", resp.GetRequest().URL.RawQuery)
	// Output: key1=123&key2=value21&key2=value22&key3=abc456
}
```

- query string

```go
cli := goz.NewClient()

resp, err := cli.Get("http://127.0.0.1:8091/get-with-query?key0=value0", goz.Options{
    Query: "key1=value1&key2=value21&key2=value22&key3=333",
})
if err != nil {
    log.Fatalln(err)
}

fmt.Printf("%s", resp.GetRequest().URL.RawQuery)
// Output: key1=value1&key2=value21&key2=value22&key3=333
```

## Post Data

- post form 

```go
func ExampleRequest_Post_withFormParams() {
	cli := goz.NewClient()

	resp, err := cli.Post("http://127.0.0.1:8091/post-with-form-params", goz.Options{
		Headers: map[string]interface{}{
			"Content-Type": "application/x-www-form-urlencoded",
		},
		FormParams: map[string]interface{}{
			"key1": 2020,
			"key2": []string{"value21", "value22"},
			"key3": "abcd张",
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	body,err := resp.GetContents()

	fmt.Printf("%v", body)
	// Output:  form params:{"key1":["2020"],"key2":["value21","value22"],"key3":["abcd张"]}
}
```

- post json 

```go
func ExampleRequest_Post_withJSON() {
	cli := goz.NewClient()

	resp, err := cli.Post("http://127.0.0.1:8091/post-with-json", goz.Options{
		Headers: map[string]interface{}{
			"Content-Type": "application/json",
		},
		JSON: struct {
			Key1 string   `json:"key1"`
			Key2 []string `json:"key2"`
			Key3 int      `json:"key3"`
		}{"value1", []string{"value21", "value22"}, 333},
	})
	if err != nil {
		log.Fatalln(err)
	}

	body := resp.GetBody()
	defer body.Close()
	fmt.Printf("%T", body)
	// Output:  *http.cancelTimerBody
}
```

## Request Headers 

```go
cli := goz.NewClient()

resp, err := cli.Post("http://127.0.0.1:8091/post-with-headers", goz.Options{
    Headers: map[string]interface{}{
        "User-Agent": "testing/1.0",
        "Accept":     "application/json",
        "X-Foo":      []string{"Bar", "Baz"},
    },
})
if err != nil {
    log.Fatalln(err)
}

headers := resp.GetRequest().Header["X-Foo"]
fmt.Println(headers)
// Output: [Bar Baz]
```

## Response 

```go
cli := goz.NewClient()
resp, err := cli.Get("http://127.0.0.1:8091/get")
if err != nil {
    log.Fatalln(err)
}

body, err := resp.GetBody()
if err != nil {
    log.Fatalln(err)
}
fmt.Printf("%T", body)
// Output: goz.ResponseBody

part := body.Read(30)
fmt.Printf("%T", part)
// Output: []uint8

contents := body.GetContents()
fmt.Printf("%T", contents)
// Output: string

fmt.Println(resp.GetStatusCode())
// Output: 200

fmt.Println(resp.GetReasonPhrase())
// Output: OK

headers := resp.GetHeaders()
fmt.Printf("%T", headers)
// Output: map[string][]string

flag := resp.HasHeader("Content-Type")
fmt.Printf("%T", flag)
// Output: bool

header := resp.GetHeader("content-type")
fmt.Printf("%T", header)
// Output: []string
    
headerLine := resp.GetHeaderLine("content-type")
fmt.Printf("%T", headerLine)
// Output: string
```

## Proxy

```go
cli := goz.NewClient()

resp, err := cli.Get("https://www.fbisb.com/ip.php", goz.Options{
    Timeout: 5.0,
    Proxy:   "http://127.0.0.1:1087",
})
if err != nil {
    log.Fatalln(err)
}

fmt.Println(resp.GetStatusCode())
// Output: 200
```

## Timeout 

```go
cli := goz.NewClient(goz.Options{
    Timeout: 0.9,
})
resp, err := cli.Get("http://127.0.0.1:8091/get-timeout")
if err != nil {
    if resp.IsTimeout() {
        fmt.Println("timeout")
        // Output: timeout
        return
    }
}

fmt.Println("not timeout")
```

## Download File 

```go
func ExampleRequest_Down() {
	cli := goz.NewClient()

	res := cli.Down("http://139.196.101.31:2080/GinSkeleton.jpg", "F:/2020_project/go/goz/examples/", goz.Options{
		Timeout: 5.0,
	})
	fmt.Printf("%t", res)
	// Output: true
}
```

# License

[MIT](https://opensource.org/licenses/MIT)

Copyright (c) 2017-present, [idoubi](http://idoubi.cc)
