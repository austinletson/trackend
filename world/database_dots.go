package world

// User database queries
const dotTableName = "dot"
const insertDotQuery = "INSERT INTO" + dotTableName + " (id, username, name, filename) " +
	"VALUES (?, ?, ?, ?)"
const getUserDotsQuery = "SELECT id, username, filename " +
	"FROM " + dotTableName + " " +
	"WHERE username = ?"

func (data *SqlLiteDB) insertDot(dot Dot) {

}
