package matematica

func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

func Multiplicador(x int, y int) int {
	return x * y
}

func Soma(x int, y int) int {
	return x + y
}

func Subtracao(x int, y int) int {
	return x - y
}

func Divisao(x int, y int) (resultado int) {
	resultado = x / y
	return
}

func DivisorComResto(x int, y int) (resultado int, resto int) {
	resultado = x / y
	resto = x % y
	return
}
