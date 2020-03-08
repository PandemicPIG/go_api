package main

import (
	"fmt"
	"net/http"
)

func apiHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/api" {
		http.Error(res, "Forbidden", 403)
		return
	}

	res.Header().Set("Access-Control-Allow-Origin", "*")

	switch req.Method {
	case "GET":
		getUsers(res, req)
	case "POST":
		addUser(res, req)
	case "PATCH":
		updateUser(res, req)
	case "DELETE":
		deleteUser(res, req)
	case "OPTIONS":
		res.Header().Add("Connection", "keep-alive")
		res.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PATCH")
		res.Header().Add("Access-Control-Allow-Headers", "content-type")
		res.Header().Add("Access-Control-Max-Age", "86400")
		res.WriteHeader(200)
	default:
		res.WriteHeader(405)
	}
}

func main() {
	fmt.Println("api started on port 8081")

	http.HandleFunc("/api", apiHandler)
	http.ListenAndServe(":8081", nil)
}
