package users

import (
	"github.com/fajarkarim/bookstore-api-example-with-go/datasource/db_users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindAll(c *gin.Context) {
	var (
		user  User
		users []User
	)
	const query = `SELECT * FROM users;`
	rows, err := db_users.Client.Query(query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated,
			&user.Password, &user.Status)
		users = append(users, user)
	}

	defer rows.Close()
	c.JSON(http.StatusOK, users)
}
