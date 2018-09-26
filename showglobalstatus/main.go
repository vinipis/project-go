package main

import (
	"encoding/json"
	"fmt"

	"github.com/caarlos0/env"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vinipis/project-go/showglobalstatus/repo"
	"github.com/vinipis/project-go/showglobalstatus/structs"
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
	fmt.Println()
	fmt.Println()
	fmt.Println()

	cfg := structs.Parameters{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", cfg)

	openConn.Close()
}
