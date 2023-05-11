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
	r.POST("/users", controllers.UserRegistration).
		GET("/users", controllers.UsersShow).
		PUT("/users", controllers.DepositCash)
	r.POST("/ratings", controllers.RateItem).
		GET("/ratings", controllers.RatingsShow)
	r.GET("/comments", controllers.CommentsShow).
		POST("/comments", controllers.AddComment)
	r.GET("/items", controllers.ItemsShow).
		POST("/items", controllers.ItemCreate)
	r.GET("/orders", controllers.OrdersShow).
		POST("/orders", controllers.PurchaseItem)
	r.Run()
	//POST("/items", controllers.ItemCreate)
}
