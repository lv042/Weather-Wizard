package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/pocketbase/dbx"
)

func main() {
	db, _ := dbx.Open("mysql", "user:pass@/example")

	// build a SELECT query
	//   SELECT `id`, `name` FROM `users` WHERE `name` LIKE '%Charles%' ORDER BY `id`
	q := db.Select("id", "name").
		From("users").
		Where(dbx.Like("name", "Charles")).
		OrderBy("id")

	// fetch all rows into a struct array
	var users []struct {
		ID, Name string
	}
	q.All(&users)

	// build an INSERT query
	//   INSERT INTO `users` (`name`) VALUES ('James')
	db.Insert("users", dbx.Params{
		"name": "James",
	}).Execute()

	//print the result
}
