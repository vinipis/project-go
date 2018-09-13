package repo

import (
	"database/sql"

	"github.com/vinipis/project-go/showglobalstatus/model"
)

//OpenConn Realiza a abertura da conexão com o banco de dados
func OpenConn() (connDb *sql.DB) {

	valueCnf := model.MyCnf()
	host := valueCnf[0]
	port := valueCnf[1]
	user := valueCnf[2]
	password := valueCnf[3]

	parameter := user + ":" + password + "@tcp(" + host + ":" + port + ")" + "/"

	connDb, _ = sql.Open("mysql", parameter)

	return connDb
}
