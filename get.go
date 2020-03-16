package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getUsers(res http.ResponseWriter, req *http.Request) {
	data := dataS.GetUserList()

	if len(data) > 0 {
		users := []user{}
		for _, val := range data {
			users = append(users, val)
		}
		js, _ := json.Marshal(users)
		res.WriteHeader(200)
		res.Write([]byte(fmt.Sprintf(`{"status": "OK", "data": %s}`, js)))
	} else {
		res.WriteHeader(200)
		res.Write([]byte(`{"status": "OK", "message": "No users found.", "data": []}`))
	}
}
