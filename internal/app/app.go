package app

import (
	"github.com/gin-gonic/gin"
	"site/internal/database/postgresql"
	"site/internal/transport/html"
	"site/internal/transport/json"
)

func InitStructures() {
	//var Cartridge = config.Cartridge{}
	//var Computer = config.Computer{}
}

func Run() {

	//Init router and postgresql
	r := gin.Default()
	postgresql.Connect()

	r.Static("/assets/", "web/")
	r.LoadHTMLGlob("web/html/*.html")

	//HTML handlers
	r.GET("/", html.IndexPage)
	r.GET("/registration", html.RegistrationPage)
	r.GET("/authorization", html.AuthorizationPage)
	r.GET("/adminPage", html.AdminPage)

	//JSON user endpoint
	r.POST("/user/create", json.CreateUser)
	r.POST("/user/reg", json.UserRegistration)
	r.POST("/user/auth", json.UserAuthorization)

	//JSON categories endpoint
	r.POST("/category/create", json.CreateCategory)
	r.POST("/category/get", json.GetCategories)
	r.POST("/category/update", json.UpdateCategory)
	r.POST("/category/delete", json.DeleteCategory)

	_ = r.Run(":8080")
}
