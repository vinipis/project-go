package matematica

//Calculo executa qualquer tipo de calculo bjs me liga
func Calculo(funcao func(int, int) int, numA int, numB int) int {
	return funcao(numA, numB)
}

//Multiplicador é um cool
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor é uma bosta precisa de comentario bjs
func Divisor(numA int, numB int) (resultado int) {
	resultado = numA / numB
	return
}

//DivisorComResto pare de pedir comentarios
func DivisorComResto(num1 int, num2 int) (resultado int, resto int) {
	resultado = num1 / num2
	resto = num1 % num2
	return
}
