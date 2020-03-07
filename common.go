package main

var data = []user{
	user{
		Name:  "billy",
		Email: "billy@email.com",
	},
	user{
		Name:  "james",
		Email: "james@email.com",
	},
	user{
		Name:  "david",
		Email: "david@email.com",
	},
	user{
		Name:  "john",
		Email: "john@email.com",
	},
}

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
