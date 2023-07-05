package main

import (
	"database/sql"
	"fmt"
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

	connStr := "postgres://postgres:password@localhost/dvdrental?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	CheckError(err)
	var (
		city_id     int
		city        string
		country_id  int
		last_update time.Time
	)
	rows, err := db.Query("SELECT * FROM city")
	//database üzerinden çektiğin tablodaki sütunları getirir.
	arr, err := rows.Columns()
	CheckError(err)
	for _, i := range arr {
		fmt.Println(i)
	}
	/*
		for rows.Next() {
			//Scan() :database üzerinden gelen verileri değişkenlere atmamızı sağlar
			err := rows.Scan(&city_id, &city, &country_id, &last_update)
			CheckError(err)
			//log.Printf("bulunan satır içeriği:%q", strconv.Itoa(city_id)+" "+city+" "+strconv.Itoa(country_id))
		}
	*/

	// ID ye göre veri alma
	/*
		rows2, err := db.Query("SELECT * FROM city WHERE ID=355")
		for rows2.Next() {
			err := rows2.Scan(&city_id, &city, &country_id, &last_update)
			CheckError(err)
			log.Printf("bulunan satır içeriği:%q", strconv.Itoa(city_id)+" "+city+" "+strconv.Itoa(country_id))

		}
	*/
	/*
		err = db.QueryRow("SELECT * FROM city limit 1").Scan(&city_id, &city, &country_id, &last_update)
		CheckError(err)
		log.Println("bir satır bulundu", city, city_id)
	*/
	//multiple active result set :MARS
	//BİRDEN FAZLA SORGUNUN TEK BİR QUERY ÜZERİNDE ÇALIŞMASI
	//_,err =db.Exec("DELETE FROM xtable1;DELETE FROM xtable2")

	//Preparing Queries

	//  $1 : dinamik yapı
	stmt, errQ := db.Prepare("SELECT * FROM city WHERE city_id= $1")
	CheckError(errQ)
	rows, errX := stmt.Query(400)
	CheckError(errX)
	for rows.Next() {
		rows.Scan(&city_id, &city, &country_id, &last_update)
		log.Printf("bulunan satır içeriği:%q", strconv.Itoa(city_id)+" "+city+" "+strconv.Itoa(country_id))
	}

}
