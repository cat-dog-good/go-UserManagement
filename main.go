package main

import "go-UserManagement/router"

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}