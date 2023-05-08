package migrate

import (
	Init "Project/GoProject2023/init"
	"Project/GoProject2023/models"
)

func init() {
	Init.ConnectDB()
	Init.LoadEnvVar()
}
func main() {
	Init.DB.AutoMigrate(&models.Comments{})
	Init.DB.AutoMigrate(&models.Rating{})
	Init.DB.AutoMigrate(&models.User{})
	Init.DB.AutoMigrate(&models.Client{})
	Init.DB.AutoMigrate(&models.Admin{})
	Init.DB.AutoMigrate(&models.Items{})
	Init.DB.AutoMigrate(&models.Roles{})
}
