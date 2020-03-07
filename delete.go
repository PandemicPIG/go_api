package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func deleteUser(res http.ResponseWriter, req *http.Request) {
	var deletedUser user
	json.NewDecoder(req.Body).Decode(&deletedUser)

	fmt.Println(deletedUser)

	// check if user exists
	var idx = -1
	for i, user := range data {
		if user.UserID == deletedUser.UserID {
			idx = i
			break
		}
	}

	if idx != -1 {
		// delete user data
		data = append(data[:idx], data[idx+1:]...)
		res.WriteHeader(204)
	} else {
		res.WriteHeader(404)
	}
}
