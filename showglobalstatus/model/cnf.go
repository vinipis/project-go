package model

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
	"github.com/vinipis/project-go/showglobalstatus/structs"
)

//MyCnf realiza a leitura de um arquivo my.cnf e caso não tenha ele insere variaveis default
func MyCnf() (valueCnf []string) {
	valueParameters, err := structs.ParametersTerminal()
	host := valueParameters[0]
	port := valueParameters[1]
	user := valueParameters[2]
	password := valueParameters[3]

	cfg, _ := ini.Load(os.Getenv("HOME") + "/.my.cnf")

	host = cfg.Section("client").Key("host").Validate(func(in string) string {
		if in == "" {
			return host
		}
		return in
	})
	valueCnf = append(valueCnf, host)

	port = cfg.Section("client").Key("port").Validate(func(in string) string {
		if in == "" {
			return port
		}
		return in
	})
	valueCnf = append(valueCnf, port)

	user = cfg.Section("client").Key("user").Validate(func(in string) string {
		if in == "" {
			return user
		}
		return in
	})
	valueCnf = append(valueCnf, user)

	password = cfg.Section("client").Key("password").Validate(func(in string) string {
		if in == "" {
			if err != nil {
				fmt.Printf("%+v\n", err)
			}
			return password
		}
		return in
	})
	valueCnf = append(valueCnf, password)

	fmt.Println(host)
	fmt.Println(port)
	fmt.Println(user)
	fmt.Println(password)

	return valueCnf
}
