package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(res http.ResponseWriter, req *http.Request) {
	if len(data) > 0 {
		js, _ := json.Marshal(data)
		res.Header().Set("Content-Type", "application/json; charset=UTF-8")
		res.WriteHeader(200)
		res.Write(js)
	} else {
		res.WriteHeader(204)
	}
}
