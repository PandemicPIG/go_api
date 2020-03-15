package main

import (
	"log"
	"net/http"
)

func apiHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/api" {
		http.Error(res, "Forbidden", 403)
		return
	}

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")

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
	port := ":8081"
	log.Printf("INFO: api started on port %s", port)

	http.HandleFunc("/api", apiHandler)
	http.ListenAndServe(port, nil)
}
