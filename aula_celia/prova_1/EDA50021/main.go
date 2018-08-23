package main

import "fmt"

func main() {
	/*
		Este programa devera receber dois times (String), dois resultados (INT) e dois aproveitamento de bola (Float)
		e com o case mostrar qual time ganhou, qual resultado da partida e qual o aproveitamento de bola em campo....
	*/

	var (
		time1, time2, sair         string
		qtdGolsTime1, qtdGolsTime2 int
		posseDebola1, posseDebola2 float32
	)
	fmt.Println("Este programa devera receber dois times (String), dois resultados (INT) e dois aproveitamento de bola (Float)")
	fmt.Print("e com o case mostrar qual time ganhou, qual resultado da partida e qual o aproveitamento de bola em campo....")

	fmt.Println(" ")

	fmt.Println("Digite o time 1: ")
	fmt.Scan(&time1)

	fmt.Println("Digite o time 2: ")
	fmt.Scan(&time2)

	fmt.Println("Digite o quantos gols o time 1 fez: ")
	fmt.Scan(&qtdGolsTime1)

	fmt.Println("Digite o quantos gols o time 2 fez: ")
	fmt.Scan(&qtdGolsTime2)

	fmt.Println("Digite a posse de bola do time 1: ")
	fmt.Scan(&posseDebola1)

	fmt.Println("Digite a posse de bola do time 2: ")
	fmt.Scan(&posseDebola2)

	switch {
	case qtdGolsTime1 > qtdGolsTime2:
		fmt.Printf("O Time %s Ganhou de: %v X %v, com aproveitamento de: %.2f\n", time1, qtdGolsTime1, qtdGolsTime2, posseDebola1)
	case qtdGolsTime2 > qtdGolsTime1:
		fmt.Printf("O Time %s Ganhou de: %v X %v, com aproveitamento de: %.2f\n", time2, qtdGolsTime2, qtdGolsTime1, posseDebola2)
	case qtdGolsTime1 == qtdGolsTime2:
		fmt.Printf("Os Times %s e %s empataram. Placar de:  %v X %v com aproveitamento de: %.2f, %.2f\n", time1, time2, qtdGolsTime2, qtdGolsTime2, posseDebola1, posseDebola2)

	}
	fmt.Println("Digite 'E' para sair do programa")
	fmt.Scan(&sair)
}
