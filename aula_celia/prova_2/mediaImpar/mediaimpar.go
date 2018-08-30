package mediaimpar

//Media é um funçao para calcular os numeros inpares
func Media(v []int) (valorDaSoma int, qtdimpar int) {

	for cont := 0; cont <= len(v); cont++ {
		if v[cont] != 0 {
			valorDaSoma = valorDaSoma + valorDaSoma
			qtdimpar++
		}
	}
	return
}
