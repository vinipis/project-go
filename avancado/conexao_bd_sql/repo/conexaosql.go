package repo

import (
	"fmt"
	"os"

	//github.com/go-sql-driver/mysql não é usado diretamente pela aplicação ( _ )
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	ini "gopkg.in/ini.v1"
)

//MyCnf procura um arquivo no computador e le
func MyCnf() []string {
	var host string
	var port string
	var dbname string
	var user string
	var password string
	var array []string

	cfg, err := ini.Load(os.Getenv("HOME") + "/go/src/project-go/avancado/conexao_bd_sql/repo/.my.cnf")
	if err != nil {
		os.Exit(1)
	}

	host = cfg.Section("client").Key("host").String()
	array = append(array, host)
	port = cfg.Section("client").Key("port").String()
	array = append(array, port)
	dbname = cfg.Section("client").Key("dbname").String()
	array = append(array, dbname)
	user = cfg.Section("client").Key("user").String()
	array = append(array, user)
	password = cfg.Section("client").Key("password").String()
	array = append(array, password)

	return array
}

//Db armazena a conexão com o banco de dados
var Db *sqlx.DB

//AbreConexaoComBdSQL abrindo conexão com banco de dados
func AbreConexaoComBdSQL() (err error) {
	err = nil
	array := MyCnf()

	host := array[0]
	//port := array[1]
	dbname := array[2]
	user := array[3]
	password := array[4]

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, dbname)

	Db, err = sqlx.Open("mysql", connectionString)
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println("[conexaosql.go] Ping deu erro")
		return
	}
	return
}
