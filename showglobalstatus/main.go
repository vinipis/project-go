package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

//PrintSlice precisa ter comentario
func PrintSlice(result []string) {
	fmt.Print(result)
}

func main() {
	db, err := sql.Open("mysql", "root:GomariaDB@tcp(127.0.0.1:3306)/carlos?charset=utf8")
	checkError(err)

	res, err := db.Query("SHOW GLOBAL STATUS")
	checkError(err)

	var id string
	var nome string

	for res.Next() {

		res.Scan(&id, &nome)
		user := json.NewEncoder(os.Stdout)
		users := map[string]string{id: nome}
		user.Encode(users)
	}

	db.Close()
}
