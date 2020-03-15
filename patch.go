package main

import (
	"encoding/json"
	"net/http"
)

func updateUser(res http.ResponseWriter, req *http.Request) {
	var updatedUser user
	json.NewDecoder(req.Body).Decode(&updatedUser)

	// check if user exists
	userExists := dataS.CheckUserExists(updatedUser.UserID)

	if userExists {
		if updatedUser.UserID == 0 {
			res.WriteHeader(400)
			res.Write([]byte(`{"status": "BAD REQUEST", "message": "Missing of invalid user ID."}`))
		} else if updatedUser.Email == "" && updatedUser.Name == "" {
			res.WriteHeader(400)
			res.Write([]byte(`{"status": "BAD REQUEST", "message": "Missing user name and email."}`))
		} else if updatedUser.Email != "" && !checkEmail(updatedUser.Email) {
			res.WriteHeader(400)
			res.Write([]byte(`{"status": "BAD REQUEST", "message": "Invalid user email."}`))
		} else {
			// update user data
			go dataS.EditUser(updatedUser)
			res.WriteHeader(200)
			res.Write([]byte(`{"status": "OK", "message": "User updated."}`))
		}
	} else {
		res.WriteHeader(404)
		res.Write([]byte(`{"status": "NOT FOUND", "message": "No user found."}`))
	}
}
