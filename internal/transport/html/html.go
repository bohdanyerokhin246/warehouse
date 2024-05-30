package html

import (
	"github.com/gin-gonic/gin"
)

func IndexPage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func RegistrationPage(c *gin.Context) {
	c.HTML(200, "registration.html", gin.H{})
}

func AuthorizationPage(c *gin.Context) {
	c.HTML(200, "authorization.html", gin.H{})
}

func AdminPage(c *gin.Context) {
	c.HTML(200, "admin.html", gin.H{})
}
