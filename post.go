package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func addUser(res http.ResponseWriter, req *http.Request) {
	var newUser user
	json.NewDecoder(req.Body).Decode(&newUser)

	data = append(data, newUser)

	fmt.Println(newUser)

	res.WriteHeader(201)
}
