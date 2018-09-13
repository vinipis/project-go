package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vinipis/project-go/showglobalstatus/repo"
)

func main() {
	openConn := repo.OpenConn()
	res, _ := openConn.Query("SHOW GLOBAL STATUS")

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
	openConn.Close()
}
