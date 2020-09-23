package users

import (
	"fmt"
	"github.com/fajarkarim/bookstore-api-example-with-go/datasource/db_users"
	"github.com/fajarkarim/bookstore-api-example-with-go/errors"
	"log"
)

const tableName = "users"

func FindAll() []User {
	var (
		user  User
		users []User
	)
	query := fmt.Sprintf(`SELECT * FROM %s;`, tableName)
	rows, err := db_users.Client.Query(query)

	if err != nil {
		errors.NewInternalServerError("Failed Query to Database", err)
	}

	for rows.Next() {
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated,
			&user.Password, &user.Status)
		users = append(users, user)
	}

	defer rows.Close()
	return users
}

func FindById(userID string) User {
	var user User
	const query = `SELECT * FROM users where id = ?;`

	row := db_users.Client.QueryRow(query, userID)

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated,
		&user.Password, &user.Status)

	if err != nil {
		log.Println("error ", err)
	}

	return user
}
