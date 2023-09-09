package main

import (
	"fmt"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "The server is running")
	})

	http.ListenAndServe(":8080", nil)
}