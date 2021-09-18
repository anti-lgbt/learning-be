package main

import (
	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/models"
	"github.com/anti-lgbt/learning-be/routes"
)

func main() {
	if err := config.InitializeConfig(); err != nil {
		config.Logger.Error(err.Error())
		return
	}

	config.DataBase.AutoMigrate(&models.User{}, &models.ProductType{}, &models.Product{}, &models.Comment{})
	r := routes.SetupRouter()
	// running
	r.Listen(":3000")
}
