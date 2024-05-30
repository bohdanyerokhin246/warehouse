package json

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"site/internal/config"
	"site/internal/database/postgresql"
)

func CreateUser(c *gin.Context) {
	userID, err := postgresql.CreateNewUser()
	if err != nil {
		c.JSON(200, gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"Error": nil,
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

	fmt.Println(user)

	c.JSON(200, gin.H{
		"Error": nil,
	})
}
