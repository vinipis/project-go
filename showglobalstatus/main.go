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

	parameter := user + ":" + password + "@tcp(" + host + ":" + port + ")" + "/"

	connDb, _ := sql.Open("mysql", parameter)

	res, _ := connDb.Query("SHOW GLOBAL STATUS")

	var (
		counter    int
		key, value string
	)

	for res.Next() {
		res.Scan(&key, &value)
		mapCreation := map[string]string{key: value}
		transformsJSON, _ := json.Marshal(mapCreation)

		for _, valueJSON := range transformsJSON {
			mountJSON := string(valueJSON)
			if counter == 0 && mountJSON != "}" {
				fmt.Print(mountJSON)

			} else if mountJSON != "}" && mountJSON != "{" {
				fmt.Print(mountJSON)
			}
		}
		if false != res.Next() {
			fmt.Print(",")
		} else {
			fmt.Print("}")
		}
		counter++
	}
	connDb.Close()
}
