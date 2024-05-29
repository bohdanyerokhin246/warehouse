package json

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"site/internal/config"
)

func GetTestInfo(c *gin.Context) {
	type Test struct {
		ProductName string `json:"product_name"`
		ProductID   int32  `json:"product_id"`
	}
	test := Test{
		ProductID:   256,
		ProductName: "Хліб"}

	c.IndentedJSON(200, test)

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
