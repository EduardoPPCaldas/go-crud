package main

import (
	"github.com/EduardoPPCaldas/go-crud/initializers"
	"github.com/EduardoPPCaldas/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}