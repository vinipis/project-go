package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
)

func myCnf() []string {
	var host, port, user, password string
	var array []string

	cfg, _ := ini.Load(os.Getenv("HOME") + "/.my.cnf")

	host = cfg.Section("client").Key("host").Validate(func(in string) string {
		if len(in) == 0 {
			return "127.0.0.1"
		}
		return in
	})
	array = append(array, host)
	port = cfg.Section("client").Key("port").Validate(func(in string) string {
		if len(in) == 0 {
			return "3306"
		}
		return in
	})
	array = append(array, port)
	user = cfg.Section("client").Key("user").Validate(func(in string) string {
		if len(in) == 0 {
			return "root"
		}
		return in
	})
	array = append(array, user)
	password = cfg.Section("client").Key("password").Validate(func(in string) string {
		if len(in) == 0 {
			return "GomariaDB"
		}
		return in
	})
	array = append(array, password)

	return array
}

func main() {
	array := myCnf()
	host := array[0]
	port := array[1]
	user := array[2]
	password := array[3]

	parametro := user + ":" + password + "@tcp(" + host + ":" + port + ")" + "/"

	db, _ := sql.Open("mysql", parametro)

	res, _ := db.Query("SHOW GLOBAL STATUS")

	var (
		cont     int
		id, nome string
	)

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
