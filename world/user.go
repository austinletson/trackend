package world

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Bio      string `json:"bio" binding:"required"`
}

type UserOp interface {
	AddUser(user User)
	DeleteUser(username string)
	VerifyPassword(username string, password string) bool
	GetUsers() []string
}
