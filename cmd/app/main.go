package main

import (
	"github.com/gin-gonic/gin"
	"site/internal/database/postgresql"
	"site/internal/transport/html"
	"site/internal/transport/json"
)

var router *gin.Engine

func main() {
	postgresql.Connect()
	router = gin.Default()
	router.Static("/assets/", "web/")
	router.LoadHTMLGlob("web/html/*.html")
	router.GET("/", html.IndexPage)
	router.GET("/test", html.TestPage)
	router.GET("/getTestInfo", json.GetTestInfo)
	router.GET("/registration", html.RegistrationPage)
	router.GET("/authorization", html.AuthorizationPage)
	router.POST("/user/reg", json.UserRegistration)
	router.POST("/user/auth", json.UserAuthorization)
	_ = router.Run(":8080")
}
