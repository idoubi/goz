package goz

import (
	"fmt"
	"github.com/qifengzhang007/goz"
	"log"
)

func ExampleResponse_GetBody() {
	cli := goz.NewClient()
	resp, err := cli.Get("http://127.0.0.1:8091/get")
	if err != nil {
		log.Fatalln(err)
	}

	body := resp.GetBody()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T", body) //   *http.cancelTimerBody 就是对 body 数据类型 io.ReadCloser 的二次封装
	// Output:  *http.cancelTimerBody
}

func ExampleResponseBody_GetContents() {
	cli := goz.NewClient()
	resp, err := cli.Get("http://127.0.0.1:8091/get")
	if err != nil {
		log.Fatalln(err)
	}

	contents, err := resp.GetContents()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", contents)
	// Output: http get
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
	// Output: not timeout

}
