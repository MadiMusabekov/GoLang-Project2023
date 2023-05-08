package main

import Init "GoProject2023/init"

func init() {
	Init.LoadEnvVar()
	Init.ConnectDB()
}
func main() {
	r := gin.Default()
	r.PUT("/items/:id", controller.)
}
