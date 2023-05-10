package main

import (
	Init "GoProject2023/init"
	"GoProject2023/models"
)

func init() {
	Init.ConnectDB()
	Init.LoadEnvVar()
}
func main() {
	Init.DB.AutoMigrate(&models.Item{})
	Init.DB.AutoMigrate(&models.User{})
	Init.DB.AutoMigrate(&models.Comment{})
	Init.DB.AutoMigrate(&models.Rating{})
	Init.DB.AutoMigrate(&models.Client{})
	Init.DB.AutoMigrate(&models.Admin{})
	Init.DB.AutoMigrate(&models.Roles{})
}
