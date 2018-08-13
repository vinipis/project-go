package model

import "errors"

type Imovel struct {
	X     int    `json:"Coordenada_x"`
	Y     int    `json:"Coordenada_Y"`
	Nome  string `json:"Nome"`
	valor int
}

var ErrValorDeImovelInvalido = errors.New("Valor informado não é valido para um imovel...")
var ErrValorDeImovelMuitoAlto = errors.New("O valor informado é muito alto...")

func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 1000000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}
	i.valor = valor
	return
}

func (i *Imovel) GetValor() int {
	return i.valor
}
