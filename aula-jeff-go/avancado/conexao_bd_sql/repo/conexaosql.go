package repo

import (
	"fmt"

	//github.com/go-sql-driver/mysql não é usado diretamente pela aplicação ( _ )
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//Db armazena a conexão com o banco de dados
var Db *sqlx.DB

//AbreConexaoComBdSQL abrindo conexão com banco de dados
func AbreConexaoComBdSQL() (err error) {
	err = nil

	Db, err = sqlx.Open("mysql", "root:GomariaDB@tcp(127.0.0.1:3306)/carlos?parseTime=true")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println("[AbreConexaoComBdSQL] Ping deu erro")
		return
	}
	return
}
