package main

import (
"fmt"
"os"
//"encoding/json"

_"github.com/go-sql-driver/mysql"	
"database/sql"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:s3cr3t@tcp(127.0.0.1:3306)/carlos?charset=utf8")

	res, err:= db.Query("SHOW GLOBAL STATUS")
	checkError(err)

	var name string
	var variable string
	//var result map[string]string
	var result []string

	for  res.Next(){
		if err := res.Scan(&name, &variable); err != nil{
			checkError(err)
		}
		result = []string{name, variable}
/*	 	user := json.NewEncoder(os.Stdout)
		user.Encode(result)*/
	}
	//fmt.Printf(name)
	fmt.Println(result)

	db.Close()
}