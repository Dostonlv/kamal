package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloServer)
	http.ListenAndServe(":80", nil)
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body><h1>Hello, World!</h1></body></html>")
}
