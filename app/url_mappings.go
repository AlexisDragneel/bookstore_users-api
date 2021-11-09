package app

import "github.com/alexisdragneel/bookstore_users-api/controllers"

func MapUrls() {
	router.GET("/ping", controllers.Ping)

	router.GET("/users", controllers.GetUser)
	router.GET("/users/:user_id", controllers.FindUser)
	router.POST("/users", controllers.CreateUser)
}
