package main

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {

	var (
		actor_id    int
		first_name  string
		last_name   string
		last_update time.Time
	)

	connStr := "postgres://postgres:password@localhost/dvdrental?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	CheckError(err)
	/*
		rows, errX := db.Query("SELECT * FROM actor")
		fmt.Println(rows.Columns())
		CheckError(errX)

			for rows.Next() {
				err = rows.Scan(&actor_id, &first_name, &last_name, &last_update)
				CheckError(err)
				log.Printf("bulunan içerik :%q", strconv.Itoa(actor_id)+" "+first_name+" "+last_name)
			}
	*/

	rows, err := db.Query("SELECT * FROM actor WHERE actor_id=$1", 5)

	for rows.Next() {
		err = rows.Scan(&actor_id, &first_name, &last_name, &last_update)
		CheckError(err)
		log.Printf("bulunan içerik :%q", strconv.Itoa(actor_id)+" "+first_name+" "+last_name)
	}

	//db.Prepare("SELECT * FROM actor WHERE city_id= $1")

}
