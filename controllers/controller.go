package controllers

import (
	"Go-CRUD/models"
	Init "GoProject2023/init"
	"GoProject2023/models"
	"github.com/gin-gonic/gin"
)

func ItemShow(c *gin.Context) {
	var items []models.Items
	Init.DB.Find(&items)

	c.JSON(200, gin.H{
		"items": items,
	})
}

func UserRegistration(c *gin.Context) {
	var body struct {
		ID          int
		Items       []models.Items
		Username    string
		Email       string
		Password    string
		SocialMedia string
		Cash        int
	}
	c.Bind(&body)

	user := models.User{
		ID: body.ID, Items: body.Items, Username: body.Username,
		Email: body.Email, Password: body.Password, SocialMedia: body.SocialMedia,
		Cash: body.Cash,
	}
	Init.DB.Create(&user)
	c.JSON(200, gin.H{
		"user": user,
	})
}
