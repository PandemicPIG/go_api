package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func updateUser(res http.ResponseWriter, req *http.Request) {
	var updatedUser user
	json.NewDecoder(req.Body).Decode(&updatedUser)

	fmt.Println(updatedUser)

	// check if user exists
	var idx = -1
	for i, user := range data {
		if user.UserID == updatedUser.UserID {
			idx = i
			break
		}
	}

	if idx != -1 {
		// update user data
		if updatedUser.Name != "" {
			data[idx].Name = updatedUser.Name
		}

		if updatedUser.Email != "" {
			data[idx].Email = updatedUser.Email
		}

		res.WriteHeader(204)
	} else {
		res.WriteHeader(404)
	}
}
