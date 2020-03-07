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

	switch req.Method {
	case "GET":
		getUsers(res, req)
	case "POST":
		addUser(res, req)
	case "PATCH":
		updateUser(res, req)
	case "DELETE":
		deleteUser(res, req)
	default:
		res.WriteHeader(405)
	}
}

func main() {
	fmt.Println("api started on port 8081")

	http.HandleFunc("/api", apiHandler)
	http.ListenAndServe(":8081", nil)
}
