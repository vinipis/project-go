package model

import (
	"os"

	"github.com/go-ini/ini"
)

//MyCnf realiza a leitura de um arquivo my.cnf e caso não tenha ele insere variaveis default
func MyCnf() []string {
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
