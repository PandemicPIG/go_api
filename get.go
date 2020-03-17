package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getUsers(res http.ResponseWriter, req *http.Request) {
	js, _ := json.Marshal(userS.GetUserList())
	res.WriteHeader(200)
	res.Write([]byte(fmt.Sprintf(`{"status": "OK", "data": %s}`, js)))
}
