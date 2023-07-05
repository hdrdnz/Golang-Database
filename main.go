package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	} else {
		fmt.Println("connected")
	}
}

func main() {

	var (
		first_name string
		last_name  string
	)

	connStr := "postgres://postgres:password@localhost/dvdrental?sslmode=disable"
	connStr2 := "postgres://postgres:password@localhost/Example?sslmode=disable"
	db2, err := sql.Open("postgres", connStr2)
	db, err := sql.Open("postgres", connStr)
	CheckError(err)
	//db.Exec("INSERT INTO actor(first_name,last_name,last_update) VALUES('Ahmet','Yilmaz'),'2013-05-26 14:47:57.62'")
	rows, err := db.Query("SELECT first_name, last_name FROM actor")
	CheckError(err)

	//veri tabanını kapatman gerekiyor.
	defer db.Close()
	for rows.Next() {

		err := rows.Scan(&first_name, &last_name)
		CheckError(err)
		fmt.Println("bulunan satır ierigi:", first_name+" "+last_name)
		fmt.Println(first_name, last_name)
		//createStatement := "'customer'('ID' int(11), 'firstName' varchar(45),'lastName' varchar(45))"

	}

	query := `CREATE TABLE IF NOT EXISTS users(user_id int , first_name text,  
		last_name text)`

	res, err := db2.Exec(query)
	CheckError(err)
	//rowsAffected ile kaç kayıt etkilendiği ekrana döner.
	//LastInsertId : en son eklediğin kaydın ID sini döner.
	rowCount, err := res.RowsAffected()
	fmt.Println(rowCount)
	CheckError(err)

	_, err = db2.Exec(` CREATE TABLE  IF NOT EXISTS products(product_ID int, product_name varchar(45))`)
	CheckError(err)
	_, err = db2.Exec(`CREATE TABLE IF NOT EXISTS orders(urun_ID int, musteri_ID int, adet int,toplam_fiyat int)`)
	CheckError(err)

}
