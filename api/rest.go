package api

import (
	"net/http"

	"github.com/austinletson/trackend/world"
	"github.com/gin-gonic/gin"
)

type Env struct {
	Data world.UserOp
}

type restUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (env *Env) ApiIndex() {
	r := gin.Default()
	r.POST("/register", env.apiRegisterUser)
	r.GET("/users", env.apiGetUsers)
	r.Run()

}

func (env *Env) apiGetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": env.Data.GetUsers(),
	})
}

func (env *Env) apiRegisterUser(c *gin.Context) {
	var rUser restUser
	c.BindJSON(&rUser)

	if rUser.Username == "" || rUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Must provide username and password fields",
			"status":  http.StatusBadRequest,
		})
		return
	}

	env.Data.AddUser(restUserToWorldUser(rUser))
	c.JSON(http.StatusOK, gin.H{
		"message": "Username and password added",
		"status":  http.StatusOK,
	})
}

func worldUserToRestUser(wUser world.User) (rUser restUser) {
	rUser = restUser{
		Username: wUser.Username,
		Password: wUser.Password,
	}
	return rUser
}

func restUserToWorldUser(rUser restUser) (wUser world.User) {
	wUser = world.User{
		Username: rUser.Username,
		Password: rUser.Password,
	}
	return wUser
}
