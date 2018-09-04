package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
)

func myCnf() []string {
	var host, port, user, password string
	var array []string

	cfg, err := ini.Load(os.Getenv("HOME") + "/.my.cnf")
	if err != nil {
		os.Exit(1)
	}

	host = cfg.Section("client").Key("host").String()
	array = append(array, host)
	port = cfg.Section("client").Key("port").String()
	array = append(array, port)
	user = cfg.Section("client").Key("user").String()
	array = append(array, user)
	password = cfg.Section("client").Key("password").String()
	array = append(array, password)

	return array
}

func main() {
	array := myCnf()
	host := array[0]
	port := array[1]
	user := array[2]
	password := array[3]

	teste := user + ":" + password + "@tcp(" + host + ":" + port + ")" + "/"

	db, _ := sql.Open("mysql", teste)

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

	arquivo, _ := ioutil.ReadFile(os.Getenv("HOME") + "/.my.cnf")
	str := string(arquivo)
	fmt.Println("")
	fmt.Println("")
	fmt.Println(str)

	db.Close()
}
