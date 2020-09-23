package app

import (
	"github.com/fajarkarim/bookstore-api-example-with-go/controllers/ping"
	"github.com/fajarkarim/bookstore-api-example-with-go/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users", users.FetchUsers)
	router.GET("/users/:user_id", users.FetchUserById)
}
