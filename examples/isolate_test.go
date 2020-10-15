package goz

import (
	"fmt"
	"github.com/idoubi/goz"
	"log"
	"testing"
)

/* 末來要進入容器的程式碼
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func temporaryServer() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/temporary", temporaryHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", myRouter))
}

func temporaryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/temporary":
		temp := "temporary\n"
		fmt.Fprintf(w, temp)
	}
}

func main() {
	temporaryServer()
}
*/

// 先建立單元測試隔離的測試環境，這裡測試要引入容器之後再處理
func Test_Example_Isolate(t *testing.T) {
	cli := goz.NewClient()
	resp, err := cli.Get("http://172.17.0.1:9000/temporary")
	if err != nil {
		log.Fatalln(err)
	}

	body, _ := resp.GetBody()
	fmt.Println(resp.GetStatusCode(), body)
	// Output: 200
}