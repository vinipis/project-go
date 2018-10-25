package model

import (
	"os"

	"github.com/go-ini/ini"
	"github.com/vinipis/project-go/showglobalstatus/structs"
)

//MyCnf realiza a leitura de um arquivo my.cnf e caso não tenha ele insere variaveis default
func MyCnf() (valueCnf []string) {
	valueParameters, validaflag := structs.ParametersTerminal()
	host, hostflag := valueParameters[0], validaflag[0]
	port, portflag := valueParameters[1], validaflag[1]
	user, userflag := valueParameters[2], validaflag[2]
	password, passflag := valueParameters[3], validaflag[3]

	cfg, err := ini.Load(os.Getenv("HOME") + "/.my.cnf")

	if err == nil {
		host = cfg.Section("client").Key("host").Validate(func(in string) string {
			if in == "" || hostflag == 1 {
				return host
			}
			return in
		})
	}
	valueCnf = append(valueCnf, host)

	if err == nil {
		port = cfg.Section("client").Key("port").Validate(func(in string) string {
			if in == "" || portflag == 1 {
				return port
			}
			return in
		})
	}
	valueCnf = append(valueCnf, port)

	if err == nil {
		user = cfg.Section("client").Key("user").Validate(func(in string) string {
			if in == "" || userflag == 1 {
				return user
			}
			return in
		})
	}
	valueCnf = append(valueCnf, user)

	if err == nil {
		password = cfg.Section("client").Key("password").Validate(func(in string) string {
			if in == "" || passflag == 1 {
				return password
			}
			return in
		})
	}
	valueCnf = append(valueCnf, password)

	return valueCnf
}
