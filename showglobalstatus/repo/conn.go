package repo

import (
	"database/sql"

	"github.com/vinipis/project-go/showglobalstatus/model"
)

//OpenConn Realiza a abertura da conexão com o banco de dados
func OpenConn() (connDb *sql.DB) {

	array := model.MyCnf()
	host := array[0]
	port := array[1]
	user := array[2]
	password := array[3]

	parameter := user + ":" + password + "@tcp(" + host + ":" + port + ")" + "/"

	connDb, _ = sql.Open("mysql", parameter)

	return connDb
}
