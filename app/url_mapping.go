package app

import "github.com/KaT0819/sample-go-gin-api/app/controllers/users"

func mapUrls() {
	router.POST("users", users.Create)
}
