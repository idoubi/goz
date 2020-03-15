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
	http.HandleFunc("/get-timeout", getTimeout)
	http.HandleFunc("/get-with-query", getWithQuery)
	http.HandleFunc("/post", post)
	http.HandleFunc("/post-with-headers", postWithHeaders)
	http.HandleFunc("/post-with-cookies", postWithCookies)
	http.HandleFunc("/post-with-form-params", postWithFormParams)
	http.HandleFunc("/post-with-json", postWithJSON)
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
