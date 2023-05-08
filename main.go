package main

import (
	"GoProject2023/controllers"
	Init "GoProject2023/init"
	"github.com/gin-gonic/gin"
)

func init() {
	Init.LoadEnvVar()
	Init.ConnectDB()
}
func main() {
	r := gin.Default()
	r.POST("/users", controllers.UserRegistration)
	r.GET("/users", controllers.UsersShow)
	r.Run()
}
