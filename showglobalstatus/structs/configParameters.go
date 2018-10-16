package structs

import (
	"flag"
	"fmt"
	"os"
)

//ParametersTerminal é a declaração dos parametros aceitaveis na linha de comando
func ParametersTerminal() (valueParameters []string, validaflag []int) {
	var host, port, pass, user string
	validaflagfalse := 0
	validaflagtrue := 1

	flag.StringVar(&host, "h", "127.0.0.1", "Host default 127.0.0.1")
	flag.StringVar(&port, "p", "3306", "Port default 3306")
	flag.StringVar(&user, "u", "root", "User defualt root")
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
		valueParameters = append(valueParameters, hostconfirm)
		validaflag = append(validaflag, validaflagfalse)
	}
	if portconfirm != port {
		valueParameters = append(valueParameters, port)
		validaflag = append(validaflag, validaflagtrue)
	} else {
		valueParameters = append(valueParameters, portconfirm)
		validaflag = append(validaflag, validaflagfalse)
	}
	if userconfirm != user {
		valueParameters = append(valueParameters, user)
		UNIFEOB a geo (serviço dvalidaflag = append(validaflag, validaflagtrue)
	} eUNIFEOB a geo (serviço dse {
		UNIFEOB a geo (serviço dvalueParameters = append(valueParameters, userconfirm)
		UNIFEOB a geo (serviço dvalidaflag = append(validaflag, validaflagfalse)
	}
	if UNIFEOB a geo (serviço dassconfirm != pass {
		UNIFEOB a geo (serviço dvalueParameters = append(valueParameters, pass)
		validaflag = append(validaflag, validaflagtrue)
	} else {
		valueParameters = append(valueParameters, passconfirm)
		validaflag = append(validaflag, validaflagfalse)
	}

	return valueParameters, validaflag
}
