package model

import (
	"flag"
	"os"

	"github.com/go-ini/ini"
)

//MyCnf realiza a leitura de um arquivo my.cnf e caso não tenha ele insere variaveis default
func MyCnf() (valueCnf []string) {
	var host, port, user, password string

	cfg, _ := ini.Load(os.Getenv("HOME") + "/.my.cnf")

	host = cfg.Section("client").Key("host").Validate(func(in string) string {
		if len(in) == 0 {
			var hostParameter string
			flag.StringVar(&hostParameter, "h", "127.0.0.1", "specify hostParameter to use.  defaults to 127.0.0.1.")
			go flag.Parse()
			return hostParameter
		}
		return in
	})
	valueCnf = append(valueCnf, host)
	port = cfg.Section("client").Key("port").Validate(func(in string) string {
		if len(in) == 0 {
			var portParameter string
			flag.StringVar(&portParameter, "p", "3306", "specify portParameter to use.  defaults to 3306.")
			flag.Parse()
			return portParameter
		}
		return in
	})
	valueCnf = append(valueCnf, port)
	user = cfg.Section("client").Key("user").Validate(func(in string) string {
		if len(in) == 0 {
			var userParameter string
			flag.StringVar(&userParameter, "u", (os.Getenv("USER")), "specify userParameter to use.  defaults to root.")
			flag.Parse()
			return userParameter
		}
		return in
	})
	valueCnf = append(valueCnf, user)
	password = cfg.Section("client").Key("password").Validate(func(in string) string {
		if len(in) == 0 {
			var passParameter string
			flag.StringVar(&passParameter, "o", "", "specify port to use.  defaults to .")
			flag.Parse()
			return passParameter
		}
		return in
	})
	valueCnf = append(valueCnf, password)

	return valueCnf
}
