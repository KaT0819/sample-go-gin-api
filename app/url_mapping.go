package app

import "sample-go-gin-api/controllers/users"

func mapUrls() {
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
}
