package users

import (
	"github.com/fajarkarim/bookstore-api-example-with-go/domain/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FetchUsers(context *gin.Context) {
	foundUsers := users.FindAll()
	context.JSON(http.StatusOK, foundUsers)
}

func FetchUserById(context *gin.Context) {
	userId := context.Param("user_id")
	user := users.FindById(userId)
	context.JSON(http.StatusOK, user)
}
