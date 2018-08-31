package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//núreros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é:", reflect.TypeOf(32000))

	//sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal... uint8 uint16 uint32 uint64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	//representa um mapeamento da tabela Unicode (int32)
	var i2 rune = 'a'
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do Literal 49.99 é", reflect.TypeOf(49.99)) //float64

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	s1 := "Ola meu nome é vinicius"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	//string com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Vinicius`
	fmt.Println("O tamanho da string é", len(s2))

	//char???
	//var x char = 'b'
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}
