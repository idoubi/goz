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
	http.HandleFunc("/post-with-form-params", postWithFormParams)
	http.HandleFunc("/post-with-json", postWithJSON)

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
