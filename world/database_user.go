package world

import "log"

// User database queriesqq
const userTableName = "users"
const addUserQuery = "INSERT INTO " + userTableName + " (username, password, bio) " +
	"VALUES (?, ?, ?)"

const deleteUserQuery = "DELETE FROM " + userTableName + " " +
	"WHERE username=?"

const verifyPasswordQuery = "SELECT username, Password " +
	"FROM " + userTableName + " " +
	"WHERE username=? AND password=?"

const getUsersQuery = "SELECT username " +
	"FROM " + userTableName

const updateUserQuery = "UPDATE " + userTableName + " " +
	"SET username=?, password=?, bio=? WHERE username=?"

func (data *SqlLiteDB) AddUser(user User) {
	statement, err := data.db.Prepare(addUserQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(user.Username, user.Password, user.Bio)
	if err != nil {
		log.Fatal(err)
	}
}

func (data *SqlLiteDB) DeleteUser(username string) {
	statement, err := data.db.Prepare(deleteUserQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(username)
	if err != nil {
		log.Fatal(err)
	}
}

func (data *SqlLiteDB) VerifyPassword(username string, password string) bool {
	rows, err := data.db.Query(verifyPasswordQuery, username, password)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var usernameFromDb, passwordFromDb string
		err := rows.Scan(&usernameFromDb, &passwordFromDb)
		if err != nil {
			log.Fatal(err)
		}
		if username == usernameFromDb && password == passwordFromDb {
			return true
		}
	}
	return false
}

func (data *SqlLiteDB) GetUsers() []string {
	rows, err := data.db.Query(getUsersQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var usernames []string
	for rows.Next() {
		var usernameFromDb string
		err := rows.Scan(&usernameFromDb)
		if err != nil {
			log.Fatal(err)
		}
		usernames = append(usernames, usernameFromDb)
	}

	return usernames
}

func (data *SqlLiteDB) UpdateUser(user User) error {
	statement, err := data.db.Prepare(updateUserQuery)
	if err != nil {
		return err
	}
	_, err = statement.Exec(user.Username, user.Password, user.Bio, user.Username)
	if err != nil {
		return err
	}
	return nil
}
