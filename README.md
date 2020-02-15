# goz

A fantastic HTTP request library used in golang. Inspired by [guzzle](https://github.com/guzzle/guzzle)

## Installation

```
go get -u github.com/idoubi/goz
```


## Documentation

API documentation can be found here:
https://godoc.org/github.com/idoubi/goz


## Basic Usage

```go
package main

import (
    "github.com/idoubi/goz"
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
cli := goz.NewClient()

resp, err := cli.Get("http://127.0.0.1:8091/get-with-query", goz.Options{
    Query: map[string]interface{}{
        "key1": "value1",
        "key2": []string{"value21", "value22"},
        "key3": "333",
    },
})
if err != nil {
    log.Fatalln(err)
}

fmt.Printf("%s", resp.GetRequest().URL.RawQuery)
// Output: key1=value1&key2=value21&key2=value22&key3=333
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
cli := goz.NewClient()

resp, err := cli.Post("http://127.0.0.1:8091/post-with-form-params", goz.Options{
    Headers: map[string]interface{}{
        "Content-Type": "application/x-www-form-urlencoded",
    },
    FormParams: map[string]interface{}{
        "key1": "value1",
        "key2": []string{"value21", "value22"},
        "key3": "333",
    },
})
if err != nil {
    log.Fatalln(err)
}

body, _ := resp.GetBody()
fmt.Println(body)
// Output: form params:{"key1":["value1"],"key2":["value21","value22"],"key3":["333"]}
```

- post json 

```go
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

body, _ := resp.GetBody()
fmt.Println(body)
// Output: json:{"key1":"value1","key2":["value21","value22"],"key3":333}
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

# License

[MIT](https://opensource.org/licenses/MIT)

Copyright (c) 2017-present, [idoubi](http://idoubi.cc)
