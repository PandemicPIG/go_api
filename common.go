package main

import (
	"log"
	"regexp"
	"strings"
	"sync"
)

var wg sync.WaitGroup

var dataS = newDataStore(
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
)

type user struct {
	UserID int    `json:"userId,string"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type dataStore struct {
	data      map[int]user
	highestID int
}

func newDataStore(params ...user) *dataStore {
	var data = map[int]user{}
	var highestID int

	if len(params) > 0 {
		for _, val := range params {
			data[val.UserID] = val
			if val.UserID > highestID {
				highestID = val.UserID
			}
		}
	}

	var newStore = &dataStore{
		data:      data,
		highestID: highestID,
	}

	return newStore
}

func (dataS *dataStore) GetUserList() map[int]user {
	log.Printf("LOG: retrieving user list")
	return dataS.data
}

func (dataS *dataStore) CheckUserExists(userID int) bool {
	log.Printf("LOG: retrieving user")
	return dataS.data[userID].UserID != 0
}

func (dataS *dataStore) AddUser(user user, UserIDchan chan int) {
	log.Printf("INFO: ading user")
	defer wg.Done()

	dataS.highestID++
	user.UserID = dataS.highestID
	user.Email = strings.ToLower(user.Email)
	dataS.data[user.UserID] = user

	UserIDchan <- user.UserID
}

func (dataS *dataStore) EditUser(updatedUser user) {
	log.Printf("INFO: editing user")
	user := dataS.data[updatedUser.UserID]

	if updatedUser.Name != "" {
		user.Name = updatedUser.Name
	}

	if updatedUser.Email != "" {
		user.Email = strings.ToLower(updatedUser.Email)
	}

	dataS.data[user.UserID] = user
}

func (dataS *dataStore) RemoveUser(userID int) {
	log.Printf("INFO: removing user")
	delete(dataS.data, userID)
}

func checkEmailValid(email string) bool {
	// https://gist.github.com/gregseth/5582254
	RFC2822EmailRegex := regexp.MustCompile("(?:[a-zA-Z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-zA-Z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?\\.)+[a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-zA-Z0-9-]*[a-zA-Z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])")
	return RFC2822EmailRegex.MatchString(email)
}

func checkEmailExists(email string) bool {
	lowEmail := strings.ToLower(email)
	for _, v := range dataS.data {
		if v.Email == lowEmail {
			return true
		}
	}
	return false
}
