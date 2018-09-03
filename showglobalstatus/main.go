package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "root:GomariaDB@tcp(127.0.0.1:3306)/carlos?charset=utf8")

	res, _ := db.Query("SHOW GLOBAL STATUS")

	var cont int
	var id string
	var nome string

	for res.Next() {
		res.Scan(&id, &nome)
		user := map[string]string{id: nome}
		users, _ := json.Marshal(user)

		for _, valueJSON := range users {
			montarJSON := string(valueJSON)
			if cont == 0 && montarJSON != "}" {
				fmt.Print(montarJSON)

			} else if montarJSON != "}" && montarJSON != "{" {
				fmt.Print(montarJSON)
			}
		}
		if false != res.Next() {
			fmt.Print(",")
		} else {
			fmt.Print("}")
		}
		cont++
	}
	db.Close()
}
