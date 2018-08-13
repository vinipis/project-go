package operacoes

import "errors"

//ErrOperadorInvalido precisa ser comentado
var ErrOperadorInvalido = errors.New("Digite um operador valido ( +  |  -  |  *  |  / )")

//QuatroOperacoes precisa de comentario
func QuatroOperacoes(num1 int, operador string, num2 int) (Adicao int, subtracao int, multiplicacao int, divisao int, resto int, resultado int, err error) {

	if operador == "+" {
		resultado = num1 + num2
		return
	}

	if operador == "-" {
		resultado = num1 - num2
		return
	}

	if operador == "*" {
		resultado = num1 * num2
		return
	}

	if operador == "/" {
		resultado = num1 / num2
		return
	}

	if operador == "%" {
		resultado = num1 % num2
		return
	}
	err = ErrOperadorInvalido
	return
}
