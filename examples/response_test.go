package goz

import (
	"fmt"
	"log"

	"github.com/idoubi/goz"
)

func ExampleResponse_GetBody() {
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
}

func ExampleResponse_GetParsedBody() {
	cli := goz.NewClient()
	resp, err := cli.Get("http://127.0.0.1:8091/get-response-json")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := resp.GetParsedBody()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T,%v,%v", body, body.Get("code").Int(), body.Get("message").String())
	// Output: *gjson.Result,10001,参数错误
}

func ExampleResponseBody_Read() {
	cli := goz.NewClient()
	resp, err := cli.Get("http://127.0.0.1:8091/get")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := resp.GetBody()
	if err != nil {
		log.Fatalln(err)
	}

	contents := body.Read(30)

	fmt.Printf("%T", contents)
	// Output: []uint8
}

func ExampleResponseBody_GetContents() {
	cli := goz.NewClient()
	resp, err := cli.Get("http://127.0.0.1:8091/get")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := resp.GetBody()
	if err != nil {
		log.Fatalln(err)
	}

	contents := body.GetContents()

	fmt.Printf("%T", contents)
	// Output: string
}

func ExampleResponse_GetStatusCode() {
	cli := goz.NewClient()
	resp, err := cli.Get("http://127.0.0.1:8091/get")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.GetStatusCode())
	// Output: 200
}

func ExampleResponse_GetReasonPhrase() {
	cli := goz.NewClient()
	resp, err := cli.Get("http://127.0.0.1:8091/get")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.GetReasonPhrase())
	// Output: OK
}

func ExampleResponse_GetHeaders() {
	cli := goz.NewClient()
	resp, err := cli.Get("http://127.0.0.1:8091/get")
	if err != nil {
		log.Fatalln(err)
	}

	headers := resp.GetHeaders()
	fmt.Printf("%T", headers)
	// Output: map[string][]string
}

func ExampleResponse_HasHeader() {
	cli := goz.NewClient()
	resp, err := cli.Get("http://127.0.0.1:8091/get")
	if err != nil {
		log.Fatalln(err)
	}

	flag := resp.HasHeader("Content-Type")
	fmt.Printf("%T", flag)
	// Output: bool
}

func ExampleResponse_GetHeader() {
	cli := goz.NewClient()
	resp, err := cli.Get("http://127.0.0.1:8091/get")
	if err != nil {
		log.Fatalln(err)
	}

	header := resp.GetHeader("content-type")
	fmt.Printf("%T", header)
	// Output: []string
}

func ExampleResponse_GetHeaderLine() {
	cli := goz.NewClient()
	resp, err := cli.Get("http://127.0.0.1:8091/get")
	if err != nil {
		log.Fatalln(err)
	}

	header := resp.GetHeaderLine("content-type")
	fmt.Printf("%T", header)
	// Output: string
}

func ExampleResponse_IsTimeout() {
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
}
