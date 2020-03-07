package main

import (
	"fmt"
	"net/http"
)

func api(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/api" {
		http.Error(res, "Not found.", 404)
		return
	}

	switch req.Method {
	case "GET":
		fmt.Fprintf(res, "GET request")
	case "POST":
		fmt.Fprintf(res, "POST request")
	default:
		fmt.Fprintf(res, "Other method")
	}
}

func main() {
	fmt.Println("handler started on port 8081")

	http.HandleFunc("/api", api)
	http.ListenAndServe(":8081", nil)
}
