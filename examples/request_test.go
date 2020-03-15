package goz

import (
	"fmt"
	"log"
	"net/http"

	"github.com/idoubi/goz"
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

func ExampleRequest_Get_withQuery_arr() {
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

	body, _ := resp.GetBody()
	fmt.Printf("%T", body)
	// Output: goz.ResponseBody
}

func ExampleRequest_Post_withCookies_map() {
	cli := goz.NewClient()

	resp, err := cli.Post("http://127.0.0.1:8091/post-with-cookies", goz.Options{
		Cookies: map[string]string{
			"cookie1": "value1",
			"cookie2": "value2",
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	body, _ := resp.GetBody()
	fmt.Printf("%T", body)
	// Output: goz.ResponseBody
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

	body, _ := resp.GetBody()
	fmt.Printf("%T", body)
	// Output: goz.ResponseBody
}

func ExampleRequest_Post_withFormParams() {
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

	body, _ := resp.GetBody()
	fmt.Println(body)
	// Output: json:{"key1":"value1","key2":["value21","value22"],"key3":333}
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
