package model

import (
	"flag"
	"os"

	"github.com/go-ini/ini"
)

//MyCnf realiza a leitura de um arquivo my.cnf e caso não tenha ele insere variaveis default
func MyCnf() (valueCnf []string) {
	var host, port, user, password, Parameter string

	cfg, _ := ini.Load(os.Getenv("HOME") + "/.my.cnf")

	host = cfg.Section("client").Key("host").Validate(func(in string) string {
		flag.StringVar(&Parameter, "h", "127.0.0.1", "specify Parameter to use.  defaults to 127.0.0.1.")
		flag.Parse()
		if in == "" || flag.NFlag() != 0 {
			return Parameter
		}
		return in
	})
	valueCnf = append(valueCnf, host)

	port = cfg.Section("client").Key("port").Validate(func(in string) string {
		flag.StringVar(&Parameter, "p", "3306", "specify Parameter to use.  defaults to 3306.")
		flag.Parse()
		if in == "" || flag.NFlag() != 0 {
			return Parameter
		}
		return in
	})
	valueCnf = append(valueCnf, port)

	user = cfg.Section("client").Key("user").Validate(func(in string) string {
		flag.StringVar(&Parameter, "u", (os.Getenv("USER")), "specify Parameter to use.  defaults to root.")
		flag.Parse()
		if in == "" || flag.NFlag() != 0 {
			return Parameter
		}
		return in
	})
	valueCnf = append(valueCnf, user)

	password = cfg.Section("client").Key("password").Validate(func(in string) string {
		flag.StringVar(&Parameter, "o", "", "specify port to use.  defaults to password.")
		flag.Parse()
		if in == "" || flag.NFlag() != 0 {
			return Parameter
		}
		return in
	})
	valueCnf = append(valueCnf, password)

	return valueCnf
}
