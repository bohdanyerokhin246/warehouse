package json

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"site/internal/config"
	"site/internal/database/postgresql"
)

func CreateUser(c *gin.Context) {
	var user config.User
	e := c.BindJSON(&user)
	if e != nil {
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}

	userID, err := postgresql.CreateNewUser(user)

	c.JSON(200, gin.H{
		"Error": err,
		"User":  userID,
	})

}

func UserRegistration(c *gin.Context) {
	var user config.User
	e := c.BindJSON(&user)
	if e != nil {
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}

	fmt.Println(user)

	c.JSON(200, gin.H{
		"Error": nil,
	})
}

func UserAuthorization(c *gin.Context) {

	var user config.User
	e := c.BindJSON(&user)
	if e != nil {
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}

	name, err := postgresql.UserAuthorization(user.Login, user.Password)

	c.JSON(200, gin.H{
		"Name":  name,
		"Error": err,
	})
}
