package api

import (
	"github.com/austinletson/trackend/world"
	"testing"
)

type testData struct {
	users []world.User
}

func TestGetUsers(t *testing.T) {

}

// Mock data functions for test data

func (data *testData) AddUser(user world.User) {
	data.users = append(data.users, user)
}

func (data *testData) DeleteUser(username string) {
	for i := 0; i < len(data.users); i++ {
		if data.users[i].Username == username {
			data.users = append(data.users[:i], data.users[i+1:]...)
		}
	}
}
func (data *testData) VerifyPassword(username string, password string) bool {
	for i := 0; i < len(data.users); i++ {
		if data.users[i].Username == username && data.users[i].Password == password {
			return true
		}
	}
	return false
}

func (data *testData) GetUsers() []string {
	var usernames []string
	for _, user := range data.users {
		usernames = append(usernames, user.Username)
	}
	return usernames
}
