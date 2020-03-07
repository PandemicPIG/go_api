package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func addUser(res http.ResponseWriter, req *http.Request) {
	var data user
	json.NewDecoder(req.Body).Decode(&data)

	fmt.Println(data)

	res.WriteHeader(201)
}
