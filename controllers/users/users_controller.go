package users

import (
	"github.com/fajarkarim/bookstore-api-example-with-go/domain/users"
	"github.com/gin-gonic/gin"
)

func FetchUsers(context *gin.Context) {
	users.FindAll(context)
}
