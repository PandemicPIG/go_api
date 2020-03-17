package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func addUser(res http.ResponseWriter, req *http.Request) {
	var newUser user
	json.NewDecoder(req.Body).Decode(&newUser)

	if newUser.Name == "" || newUser.Email == "" {
		res.WriteHeader(400)
		res.Write([]byte(`{"status": "BAD REQUEST", "message": "Missing name or email."}`))
	} else if !checkEmailValid(newUser.Email) {
		res.WriteHeader(400)
		res.Write([]byte(`{"status": "BAD REQUEST", "message": "Invalid user email."}`))
	} else if checkEmailExists(newUser.Email) {
		res.WriteHeader(400)
		res.Write([]byte(`{"status": "BAD REQUEST", "message": "Email already exists."}`))
	} else {
		UserIDchan := make(chan int)
		wg.Add(1)
		go dataS.AddUser(newUser, UserIDchan)
		res.WriteHeader(201)
		res.Write([]byte(fmt.Sprintf(`{"status": "CREATED", "userId": %d}`, <-UserIDchan)))
		wg.Wait()
		close(UserIDchan)
	}
}
