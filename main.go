package main

import (
	"go-UserManagement/config"
	"go-UserManagement/router"
)

//初始化项目
func main() {

	config.InitDB()

	r := router.SetupRouter()
	r.Run(":8080")
}