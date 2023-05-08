package controller

import (
	"Assignment3/Init"
	"Project/GoLang-Project2023/model"
	"github.com/gin-gonic/gin"
)

func Registration(c *gin.Context) {
	var body struct {
		Items       []model.Items
		Id          int
		Username    string
		Email       string
		Password    string
		SocialMedia string
	}
	c.Bind(body)
	user := model.User{Items: body.Items, Id: body.Id, Username: body.Username, Email: body.Email, Password: body.Password, SocialMedia: body.SocialMedia}

	Init.DB.Create(&user)
	c.JSON(200, gin.H{
		"user": user,
	})
}
