package goz

import (
	"github.com/idoubi/goz"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
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
	myRouter.HandleFunc("/server/msg1", temporaryHandler1).Methods("GET")
	myRouter.HandleFunc("/server/msg2", temporaryHandler2).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", myRouter))
}

func temporaryHandler1(w http.ResponseWriter, r *http.Request) {
	temp := "temporaryMsgInServer1\n"
	fmt.Fprintf(w, temp)
}

func temporaryHandler2(w http.ResponseWriter, r *http.Request) {
	temp := "temporaryMsgInServer2\n"
	fmt.Fprintf(w, temp)
}

func main() {
	temporaryServer()
}
*/

// 先建立單元測試隔離的測試環境，這裡測試要引入容器之後再處理
func Test_Example_Isolate(t *testing.T) {
	cli := goz.NewClient()
	// 測試第一隻 API 在未隔離狀態的的回應
	responseFormServer1, err := cli.Get("http://172.17.0.1:9000/server/msg1")
	if err != nil {
		log.Fatalln(err)
	}

	bodyFromServer1, _ := responseFormServer1.GetBody()
	require.Equal(t, string(bodyFromServer1), "temporaryMsgInServer1\n")

	// 測試第二隻 API 在未隔離狀態的的回應
	responseFormServer2, err := cli.Get("http://172.17.0.1:9000/server/msg2")
	if err != nil {
		log.Fatalln(err)
	}

	bodyFromServer2, _ := responseFormServer2.GetBody()
	require.Equal(t, string(bodyFromServer2), "temporaryMsgInServer2\n")

	// 啟動 Api 的截捷功能
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "http://172.17.0.1:9000/server/msg1",
		httpmock.NewStringResponder(200, "customMsg1\n"))


	// 測試第一隻 API 在隔離狀態的回應
	responseFormIsolated1, err := cli.Get("http://172.17.0.1:9000/server/msg1")
	if err != nil {
		log.Fatalln(err)
	}

	bodyFormIsolated1, _ := responseFormIsolated1.GetBody()
	require.Equal(t, string(bodyFormIsolated1), "customMsg1\n")

	// 測試第二隻 API 在隔離狀態的回應
	responseFormIsolated2, err := cli.Get("http://172.17.0.1:9000/server/msg2")
	if err != nil {
		log.Fatalln(err)
	}

	bodyFromIsolated2, _ := responseFormIsolated2.GetBody()
	require.Equal(t, string(bodyFromIsolated2), "temporaryMsgInServer2\n")
}