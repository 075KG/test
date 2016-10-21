// hello
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	host     = "127.0.0.1"
	port     = "3306"
	user     = "root"
	pw       = "root"
	datebase = "mysql"
)

func init_db(host, port, user, pw, datebase string) *sql.DB {
	conn := user + ":" + pw + "@tcp(" + host + ":" + port + ")/" + datebase + "?charset=utf8"
	db, err := sql.Open("mysql", conn)

	if err != nil {
		log.Println(err)
	}
	fmt.Println("connection ok...")

	if err := db.Ping(); err != nil {
		log.Println("ping error...")
	}

	return db
}

func Query(db *sql.DB, str string) *sql.Rows {
	rows, err := db.Query(str)
	if err != nil {
		log.Println(err)
	}
	return rows
}

func Insert(db *sql.DB, str string) {
	insert, err := db.Exec(str)
	if err != nil {
		log.Println(err)
	}
	id, err := insert.RowsAffected()
	fmt.Println("affect the row: ", id)
}

func main() {
	//connect
	db := init_db(host, port, user, pw, datebase)
	defer db.Close()

	//insert
	//insert:= "INSERT INTO test(name,password) VALUES('bob', 'bob')"
	//Insert(db, insert)

	//query
	rows := Query(db, "select name from test where uid=3")
	defer rows.Close()
	for rows.Next() {
		//var uid int
		var name string
		rows.Scan(&name)
		fmt.Println("name:", name)
		//var name, password string
		//rows.Scan(&uid, &name, &password)
		//fmt.Println("uid:", uid, "name:", name, "password:", password)
	}

}
