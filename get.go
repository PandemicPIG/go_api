package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getUsers(res http.ResponseWriter, req *http.Request) {
	var data = []user{
		user{
			Name:  "john",
			Email: "john@email.com",
		},
		user{
			Name:  "james",
			Email: "james@email.com",
		},
		user{
			Name:  "david",
			Email: "david@email.com",
		},
	}

	fmt.Println(data)

	js, _ := json.Marshal(data)

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(200)
	res.Write(js)

}
