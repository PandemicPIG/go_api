package main

var highestID = 5

var data = []user{
	user{
		UserID: 1,
		Name:   "billy",
		Email:  "billy@email.com",
	},
	user{
		UserID: 2,
		Name:   "james",
		Email:  "james@email.com",
	},
	user{
		UserID: 3,
		Name:   "david",
		Email:  "david@email.com",
	},
	user{
		UserID: 5,
		Name:   "john",
		Email:  "john@email.com",
	},
}

type user struct {
	UserID int    `json:"userId,string"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}
