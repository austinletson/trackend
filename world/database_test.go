package world

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var testUser1 = User{
	Username: "austinletson",
	Password: "password",
	Bio:      "Hello this is a bio",
}

var testUser2 = User{
	Username: "austinletson",
	Password: "password1",
	Bio:      "Hello this is a different bio",
}

func TestCreateSqlLiteDb(t *testing.T) {
	sqlLiteData, err := NewSqlLiteDB()
	if err != nil {
		t.Errorf("Error opening database: %v", err)
	}
	defer CloseDb(*sqlLiteData)
}

func TestVerifyPasswordSqlLiteDb(t *testing.T) {
	sqlLiteData, err := NewSqlLiteDB()
	if err != nil {
		t.Errorf("Error opening database: %v", err)
	}
	defer CloseDb(*sqlLiteData)

	sqlLiteData.AddUser(testUser1)
	rightPassword := sqlLiteData.VerifyPassword(testUser1.Username, testUser1.Password)
	if rightPassword != true {
		t.Helper()
		t.Errorf("Valid password doesn't return right password")
	}
	rightPassword = sqlLiteData.VerifyPassword(testUser1.Username, "not password")
	if rightPassword != false {
		t.Helper()
		t.Errorf("Invalid password returns right password")
	}
	sqlLiteData.DeleteUser(testUser1.Username)
	rightPassword = sqlLiteData.VerifyPassword(testUser1.Username, "not password")
	if rightPassword != false {
		t.Helper()
		t.Errorf("No user returns valid password")
	}

}

func TestCreateAndUpdateUser(t *testing.T) {
	sqlLiteData, err := NewSqlLiteDB()
	if err != nil {
		t.Errorf("Error opening database: %v", err)
	}
	defer CloseDb(*sqlLiteData)

	sqlLiteData.AddUser(testUser1)
	err = sqlLiteData.UpdateUser(testUser2)
	if err != nil {
		t.Errorf("Error trying to update user: %v", err)
	}

}
