package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func addUser(res http.ResponseWriter, req *http.Request) {
	var newUser user
	json.NewDecoder(req.Body).Decode(&newUser)

	fmt.Println(newUser)

	// check if email exists
	var idx = len(data)
	for i, user := range data {
		if user.Email == newUser.Email {
			idx = i
			break
		}
	}

	if idx == len(data) {
		highestID++
		newUser.UserID = highestID
		data = append(data, newUser)
		res.WriteHeader(201)
	} else {
		res.WriteHeader(409)
	}
}
