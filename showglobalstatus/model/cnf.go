package model

import (
	"os"

	"github.com/go-ini/ini"
)

//MyCnf realiza a leitura de um arquivo my.cnf e caso não tenha ele insere variaveis default
func MyCnf() (valueCnf []string) {
	var host, port, user, password string

	cfg, _ := ini.Load(os.Getenv("HOME") + "/.my.cnf")

	host = cfg.Section("client").Key("host").Validate(func(in string) string {
		if len(in) == 0 {
			return "127.0.0.1"
		}
		return in
	})
	valueCnf = append(valueCnf, host)
	port = cfg.Section("client").Key("port").Validate(func(in string) string {
		if len(in) == 0 {
			return "3306"
		}
		return in
	})
	valueCnf = append(valueCnf, port)
	user = cfg.Section("client").Key("user").Validate(func(in string) string {
		if len(in) == 0 {
			return "root"
		}
		return in
	})
	valueCnf = append(valueCnf, user)
	password = cfg.Section("client").Key("password").Validate(func(in string) string {
		if len(in) == 0 {
			return "GomariaDB"
		}
		return in
	})
	valueCnf = append(valueCnf, password)

	return valueCnf
}
