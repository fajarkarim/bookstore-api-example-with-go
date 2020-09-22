package datasource

import "database/sql"

func init() {
	dataSourceName := "root:@tcp/db_users?charset=utf8"

	db, err := sql.Open("msql", dataSourceName)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
}
