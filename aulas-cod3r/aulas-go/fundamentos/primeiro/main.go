//Programas executaveis iniciam pelo pacote main
package main

/*
os codigos em Go são organizados em pacotes
e para usa-lós é necessario declarar um ou varios imports
*/
import (
	"fmt"
)

func main() {
	//a porta de entrada de um programa Go é a função main
	fmt.Print("hello ")
	fmt.Print("World!!")
}
