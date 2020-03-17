package main

import (
	"encoding/json"
	"net/http"
)

func deleteUser(res http.ResponseWriter, req *http.Request) {
	var user user
	json.NewDecoder(req.Body).Decode(&user)

	// check if user exists
	userExists := userS.CheckUserExists(user.UserID)

	if userExists {
		go userS.RemoveUser(user.UserID)
		res.WriteHeader(200)
		res.Write([]byte(`{"status": "OK", "message": "User deleted."}`))
	} else {
		res.WriteHeader(404)
		res.Write([]byte(`{"status": "NOT FOUND", "message": "No user found."}`))
	}
}
