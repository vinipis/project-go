package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/userbd?charset=utf8")

	res, _ := db.Query("SHOW GLOBAL STATUS")

	var id string
	var nome string
	var users []uint8
	var user map[string]string

	for res.Next() {
		res.Scan(&id, &nome)
		user = map[string]string{id: nome}
		users, _ = json.Marshal(user)
		fmt.Println(string(users))
	}
	db.Close()
}
