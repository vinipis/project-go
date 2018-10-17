package structs

import (
	"flag"
)

//ParametersTerminal é a declaração dos parametros aceitaveis na linha de comando
func ParametersTerminal() (valueParameters []string, validaflag []int) {
	var host, port, pass, user string
	validaflagfalse := 0
	validaflagtrue := 1

	flag.StringVar(&host, "h", "", "Host default 127.0.0.1")
	flag.StringVar(&port, "p", "", "Port default 3306")
	flag.StringVar(&user, "u", "", "User defualt root")
	flag.StringVar(&pass, "o", "", "Default pass has no pass")

	hostconfirm := host
	portconfirm := port
	userconfirm := user
	passconfirm := pass

	flag.Parse()

	if hostconfirm != host {
		valueParameters = append(valueParameters, host)
		validaflag = append(validaflag, validaflagtrue)
	} else {
		valueParameters = append(valueParameters, "127.0.0.1")
		validaflag = append(validaflag, validaflagfalse)
	}
	if portconfirm != port {
		valueParameters = append(valueParameters, port)
		validaflag = append(validaflag, validaflagtrue)
	} else {
		valueParameters = append(valueParameters, "3306")
		validaflag = append(validaflag, validaflagfalse)
	}
	if userconfirm != user {
		valueParameters = append(valueParameters, user)
		validaflag = append(validaflag, validaflagtrue)
	} else {
		valueParameters = append(valueParameters, "root")
		validaflag = append(validaflag, validaflagfalse)
	}
	if passconfirm != pass {
		valueParameters = append(valueParameters, pass)
		validaflag = append(validaflag, validaflagtrue)
	} else {
		valueParameters = append(valueParameters, "")
		validaflag = append(validaflag, validaflagfalse)
	}

	return valueParameters, validaflag
}
