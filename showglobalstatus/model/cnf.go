package model

import (
	"os"

	"github.com/go-ini/ini"
	"github.com/vinipis/project-go/showglobalstatus/structs"
)

//MyCnf realiza a leitura de um arquivo my.cnf e caso não tenha ele insere variaveis default
func MyCnf() (valueCnf []string) {
	valueParameters, validaflag := structs.ParametersTerminal()
	host := valueParameters[0]
	hostflag := validaflag[0]
	port := valueParameters[1]
	portflag := validaflag[1]
	user := valueParameters[2]
	userflag := validaflag[2]
	password := valueParameters[3]
	passflag := validaflag[3]

	cfg, _ := ini.Load(os.Getenv("HOME") + "/.my.cnf")

	host = cfg.Section("client").Key("host").Validate(func(in string) string {
		if in == "" || hostflag == 1 {
			return host
		}
		return in
	})
	valueCnf = append(valueCnf, host)

	port = cfg.Section("client").Key("port").Validate(func(in string) string {
		if in == "" || portflag == 1 {
			return port
		}
		return in
	})
	valueCnf = append(valueCnf, port)

	user = cfg.Section("client").Key("user").Validate(func(in string) string {
		if in == "" || userflag == 1 {
			return user
		}
		return in
	})
	valueCnf = append(valueCnf, user)

	password = cfg.Section("client").Key("password").Validate(func(in string) string {
		if in == "" || passflag == 1 {
			return password
		}
		return in
	})
	valueCnf = append(valueCnf, password)

	return valueCnf
}
