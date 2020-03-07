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
		res.WriteHeader(200)
		fmt.Fprintf(res, "PATCH request - update user")
	case "DELETE":
		res.WriteHeader(200)
		fmt.Fprintf(res, "DELETE request - delete user")
	default:
		res.WriteHeader(405)
		fmt.Fprintf(res, "Method Not Allowed")
	}
}

func main() {
	fmt.Println("handler started on port 8081")

	http.HandleFunc("/api", apiHandler)
	http.ListenAndServe(":8081", nil)
}
