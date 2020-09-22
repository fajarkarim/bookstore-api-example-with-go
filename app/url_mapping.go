package app

import "github.com/fajarkarim/bookstore-api-example-with-go/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
