package main

import (
	"database/sql"
	"fmt"
	"log"

	//"time"

	_ "github.com/lib/pq"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	/*
		var (
			actor_id    int
			first_name  string
			last_name   string
			last_update time.Time
		)
	*/
	connStr := "postgres://postgres:password@localhost/dvdrental?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	CheckError(err)
	query := "INSERT INTO actor(actor_id,first_name,last_name,last_update) VALUES(201,'havva nur','durudeniz','2013-05-26 14:47:57.62')"
	stmt, err := db.Exec(query)
	CheckError(err)
	/*
		actor_id = 5
		first_name = "havva nur"
		last_name = "durudeniz"
		last_update = time.Now()
		res, err := stmt.Exec(actor_id, first_name, last_name, last_update)
		CheckError(err)

		fmt.Println(res.RowsAffected())
	*/
	fmt.Println(stmt)
}
