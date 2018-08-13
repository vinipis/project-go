package operadora

import (
	"fundamentos/pacotes/prefixo"
	"strconv"
)

var NomeOperadora = "TIM Telecom"

var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
