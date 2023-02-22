package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type user struct {
	id    int
	name  string
	money int
	score int
}

func main() {
	db, err := sql.Open("mysql", "root:raye12345@tcp(127.0.0.1:3306)/go_test")
	check(err)
	err = db.Ping()
	check(err)
	fmt.Println("Successfully connected to the database!")
	rows, err := db.Query("SELECT * FROM userInfo")
	check(err)
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.name, &u.money, &u.score)
		check(err)
		fmt.Printf("id:%d\tname:%s\tmoney:%d\tscore:%d\n", u.id, u.name, u.money, u.score)
		//fmt.Println(u)
	}

}
