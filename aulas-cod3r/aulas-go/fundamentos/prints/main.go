package main

import (
	"fmt"
)

func main() {
	fmt.Print("mesma")
	fmt.Print(" Linha...")

	fmt.Println(" Nova")
	fmt.Println("Linha.")

	x := 3.141516
	//não funciona desse jeito
	//fmt.Println("O valor de X é: " + x)

	xs := fmt.Sprint(x)
	fmt.Println("O valor de X é: " + xs)
	fmt.Println("O valor de X é:", x)

	fmt.Printf("O valor de X é %f", x)
	fmt.Printf("O valor de X é %.2f", x)
	fmt.Printf("O valor de X é %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %.1v %v %v", a, b, b, c, d)

}
