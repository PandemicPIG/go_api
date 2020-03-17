package main

import (
	"regexp"
	"strings"
	"sync"
)

var wg sync.WaitGroup

var userS = newUserStore(
	user{
		UserID: 1,
		Name:   "Billy",
		Email:  "billy@email.com",
	},
	user{
		Name:  "James",
		Email: "james@email.com",
	},
	user{
		Name:  "David",
		Email: "david@email.com",
	},
	user{
		UserID: 3,
		Name:   "John",
		Email:  "john@email.com",
	},
)

func checkEmailValid(email string) bool {
	// https://gist.github.com/gregseth/5582254
	RFC2822EmailRegex := regexp.MustCompile("(?:[a-zA-Z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-zA-Z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?\\.)+[a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-zA-Z0-9-]*[a-zA-Z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])")
	return RFC2822EmailRegex.MatchString(email)
}

func checkEmailExists(email string) bool {
	lowEmail := strings.ToLower(email)
	for _, v := range userS.data {
		if v.Email == lowEmail {
			return true
		}
	}
	return false
}
