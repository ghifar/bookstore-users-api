package app

import "github.com/ghifar/bookstore-users-api/controller"

func mapUrls() {
	router.GET("/ping", controller.Ping)

	router.POST("/users", controller.CreateUser)
	router.GET("/users/:user_id", controller.GetUser)
	//router.GET("/users/search", controller.FindUser)
}