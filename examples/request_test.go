package goz

import (
	"fmt"
	"github.com/qifengzhang007/goz"
	"io/ioutil"
	"log"
	"net/http"
)

func ExampleRequest_Get() {
	cli := goz.NewClient()

	resp, err := cli.Get("http://127.0.0.1:8091/get")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T", resp)
	// Output: *goz.Response
}

func ExampleRequest_Down() {
	cli := goz.NewClient()

	res := cli.Down("http://139.196.101.31:2080/GinSkeleton.jpg", "F:/2020_project/go/goz/examples/", goz.Options{
		Timeout: 5.0,
	})
	fmt.Printf("%t", res)
	// Output: true
}

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

func ExampleRequest_Get_withQuery_str() {
	cli := goz.NewClient()

	resp, err := cli.Get("http://127.0.0.1:8091/get-with-query?key0=value0", goz.Options{
		Query: "key1=value1&key2=value21&key2=value22&key3=333",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", resp.GetRequest().URL.RawQuery)
	// Output: key1=value1&key2=value21&key2=value22&key3=333
}

func ExampleRequest_Get_withProxy() {
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
	fmt.Println(resp.GetContents())
	// Output: 116.153.43.128
}

func ExampleRequest_Post() {
	cli := goz.NewClient()

	resp, err := cli.Post("http://127.0.0.1:8091/post")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T", resp)
	// Output: *goz.Response
}

func ExampleRequest_Post_withHeaders() {
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
}

func ExampleRequest_Post_withCookies_str() {
	cli := goz.NewClient()

	resp, err := cli.Post("http://127.0.0.1:8091/post-with-cookies", goz.Options{
		Cookies: "cookie1=value1;cookie2=value2",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%d", resp.GetContentLength())
	//Output: 385
}

func ExampleRequest_Post_withCookies_map() {
	cli := goz.NewClient()

	//resp, err := cli.Post("http://127.0.0.1:8091/post-with-cookies", goz.Options{
	resp, err := cli.Post("http://101.132.69.236/api/v2/test_network", goz.Options{
		Cookies: map[string]string{
			"cookie1": "value1",
			"cookie2": "value2",
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	body := resp.GetBody()
	defer body.Close()
	bytes, _ := ioutil.ReadAll(body)
	fmt.Printf("%s", bytes)
	// Output: {"code":200,"msg":"OK","data":""}
}

func ExampleRequest_Post_withCookies_obj() {
	cli := goz.NewClient()

	cookies := make([]*http.Cookie, 0, 2)
	cookies = append(cookies, &http.Cookie{
		Name:     "cookie133",
		Value:    "value1",
		Domain:   "httpbin.org",
		Path:     "/cookies",
		HttpOnly: true,
	})
	cookies = append(cookies, &http.Cookie{
		Name:   "cookie2",
		Value:  "value2",
		Domain: "httpbin.org",
		Path:   "/cookies",
	})

	resp, err := cli.Post("http://127.0.0.1:8091/post-with-cookies", goz.Options{
		Cookies: cookies,
	})
	if err != nil {
		log.Fatalln(err)
	}

	body := resp.GetBody()
	fmt.Printf("%T", body)
	//Output: *http.cancelTimerBody
}
func ExampleRequest_SimplePost() {
	cli := goz.NewClient()

	resp, err := cli.Post("http://101.132.69.236/api/v2/test_network", goz.Options{
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

	contents, _ := resp.GetContents()
	fmt.Printf("%s", contents)
	// Output:  {"code":200,"msg":"OK","data":""}
}

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

	body, err := resp.GetContents()

	fmt.Printf("%v", body)
	// Output:  form params:{"key1":["2020"],"key2":["value21","value22"],"key3":["abcd张"]}
}

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

func ExampleRequest_Put() {
	cli := goz.NewClient()

	resp, err := cli.Put("http://127.0.0.1:8091/put")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T", resp)
	// Output: *goz.Response
}

func ExampleRequest_Patch() {
	cli := goz.NewClient()

	resp, err := cli.Patch("http://127.0.0.1:8091/patch")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T", resp)
	// Output: *goz.Response
}

func ExampleRequest_Delete() {
	cli := goz.NewClient()

	resp, err := cli.Delete("http://127.0.0.1:8091/delete")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T", resp)
	// Output: *goz.Response
}

func ExampleRequest_Options() {
	cli := goz.NewClient()

	resp, err := cli.Options("http://127.0.0.1:8091/options")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T", resp)
	// Output: *goz.Response
}
