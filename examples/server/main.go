package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/get", get)
	http.HandleFunc("/get-response-json", getResponseJSON)
	http.HandleFunc("/get-timeout", getTimeout)
	http.HandleFunc("/get-with-query", getWithQuery)
	http.HandleFunc("/post", post)
	http.HandleFunc("/post-with-headers", postWithHeaders)
	http.HandleFunc("/post-with-cookies", postWithCookies)
	http.HandleFunc("/post-with-form-params", postWithFormParams)
	http.HandleFunc("/post-with-json", postWithJSON)
	http.HandleFunc("/post-with-xml", postWithXML)
	http.HandleFunc("/post-with-multipart", postWithMultipart)
	http.HandleFunc("/post-with-stream-response", postWithStreamResponse)
	http.HandleFunc("/put", put)
	http.HandleFunc("/patch", patch)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/options", options)

	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		log.Fatal("Listen And Server:", err)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "http get")
}

func getResponseJSON(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{
		"code":    10001,
		"message": "参数错误",
	}
	b, _ := json.Marshal(m)

	fmt.Fprintf(w, string(b))
}

func getTimeout(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Fprintf(w, "http get timeout")
}

func getWithQuery(w http.ResponseWriter, r *http.Request) {
	q := r.URL.RawQuery
	fmt.Fprintf(w, "query:%s", q)
}

func post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "need post")
		return
	}

	fmt.Fprintf(w, "http post")
}

func postWithHeaders(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "need post")
		return
	}

	headers, _ := json.Marshal(&r.Header)

	fmt.Fprintf(w, "headers:%s", headers)
}

func postWithCookies(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "need post")
		return
	}

	cookies, _ := json.Marshal(r.Cookies())
	fmt.Fprintf(w, "cookies:%s", cookies)
}

func postWithFormParams(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "need post")
		return
	}

	r.ParseForm()

	params, _ := json.Marshal(r.Form)

	fmt.Fprintf(w, "form params:%s", params)
}

func postWithJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "need post")
		return
	}

	json, _ := ioutil.ReadAll(r.Body)

	fmt.Fprintf(w, "json:%s", json)
}

func postWithXML(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "need post")
		return
	}

	xml, _ := ioutil.ReadAll(r.Body)

	fmt.Fprintf(w, "xml:%s", xml)
}

func postWithMultipart(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "need post")
		return
	}

	r.ParseMultipartForm(100)

	for k, v := range r.Form {
		fmt.Printf("form fields: %s, %v\n", k, v)
	}
	for k := range r.MultipartForm.File {
		file, fileHeader, _ := r.FormFile(k)
		defer file.Close()
		fmt.Printf("form file: %s, %d, %v\n", fileHeader.Filename, fileHeader.Size, file)
	}

	fmt.Fprintf(w, "body:%s", "")
}

func postWithStreamResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "need post")
		return
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		fmt.Fprintf(w, "not support stream")
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	timeout := time.After(20 * time.Second)

	msgch := make(chan rune)

	go func() {
		message := "this message will response with stream\n"
		for _, c := range message {
			msgch <- c
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for {
		select {
		case msg := <-msgch:
			fmt.Fprintf(w, "data: %c\n\n", msg)
			flusher.Flush()
			log.Printf("flush data: %c\n", msg)

			if msg == '\n' {
				fmt.Fprintf(w, "data: %s\n\n", "[DONE]")
				flusher.Flush()
				log.Printf("end flush: %c\n", msg)
				return
			}
		case <-timeout:
			fmt.Fprintf(w, "data: %s\n\n", "[DONE]")
			flusher.Flush()
			log.Printf("end flush cause timeout: %c\n", '\n')
			return
		}
	}
}

func put(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		fmt.Fprintf(w, "need put")
		return
	}

	fmt.Fprintf(w, "http put")
}

func patch(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		fmt.Fprintf(w, "need patch")
		return
	}

	fmt.Fprintf(w, "http patch")
}

func delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		fmt.Fprintf(w, "need delete")
		return
	}

	fmt.Fprintf(w, "http delete")
}

func options(w http.ResponseWriter, r *http.Request) {
	if r.Method != "OPTIONS" {
		fmt.Fprintf(w, "need options")
		return
	}

	fmt.Fprintf(w, "http options")
}
