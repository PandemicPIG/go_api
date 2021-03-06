package main

import (
	"log"
	"strings"
)

type user struct {
	UserID int    `json:"userId,string"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type userStore struct {
	data      map[int]user
	highestID int
}

func newUserStore(params ...user) *userStore {
	var data = map[int]user{}
	var noIDUsers = []user{}
	var highestID int

	if len(params) > 0 {
		for _, val := range params {
			if val.UserID > 0 && val.Name != "" && val.Email != "" {
				val.Email = strings.ToLower(val.Email)
				data[val.UserID] = val
				log.Printf("INFO: user %s added to store", val.Name)
				if val.UserID > highestID {
					highestID = val.UserID
				}
			} else if val.Name != "" && val.Email != "" {
				noIDUsers = append(noIDUsers, val)
			} else {
				log.Printf("ERROR: Fileds missing for: %d %s %s", val.UserID, val.Name, val.Email)
			}
		}
	}

	var newStore = &userStore{
		data:      data,
		highestID: highestID,
	}

	newStore.addNoIDUsers(noIDUsers)

	return newStore
}

func (store *userStore) addNoIDUsers(noIDUsers []user) {
	for _, val := range noIDUsers {
		UserIDchan := make(chan int)
		wg.Add(1)
		go store.AddUser(val, UserIDchan)
		log.Printf("INFO: user %s added to store with new id: %d", val.Name, <-UserIDchan)
		wg.Wait()
		close(UserIDchan)
	}
}

func (store *userStore) GetUserList() map[int]user {
	log.Printf("LOG: retrieving user list")
	return store.data
}

func (store *userStore) CheckUserExists(userID int) bool {
	log.Printf("LOG: retrieving user")
	return store.data[userID].UserID != 0
}

func (store *userStore) AddUser(user user, UserIDchan chan int) {
	log.Printf("INFO: ading user")
	defer wg.Done()

	store.highestID++
	user.UserID = store.highestID
	user.Email = strings.ToLower(user.Email)
	store.data[user.UserID] = user

	UserIDchan <- user.UserID
}

func (store *userStore) EditUser(updatedUser user) {
	log.Printf("INFO: editing user")
	user := store.data[updatedUser.UserID]

	if updatedUser.Name != "" {
		user.Name = updatedUser.Name
	}

	if updatedUser.Email != "" {
		user.Email = strings.ToLower(updatedUser.Email)
	}

	store.data[user.UserID] = user
}

func (store *userStore) RemoveUser(userID int) {
	log.Printf("INFO: removing user")
	delete(store.data, userID)
}
