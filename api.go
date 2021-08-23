package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

var (
	empty = []byte("")
)

func main() {
	http.HandleFunc("/", test)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func test(w http.ResponseWriter, r *http.Request) {
	debugRequest(r)
	addCorsHeaders(w)
	switch r.Method {
		case "OPTIONS":
			w.Write(empty)
		case "PUT":
			w.Write(empty)
		case "GET":
			w.Write(empty)
	}
}

func debugRequest(r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
}

func addCorsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, PUT, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
