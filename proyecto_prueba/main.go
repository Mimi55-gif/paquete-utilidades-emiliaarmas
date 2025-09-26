package main

import (
	"fmt"

	utilidades "github.com/Mimi55-gif/paquete-utilidades-emiliaarmas"
)

func main() {
	// Conversor
	var valor float64
	var opcion int
	fmt.Print("Ingresa el valor en USD: ")
	fmt.Scanln(&valor)
	fmt.Println("1) EUR, 2) LB, 3) WON, 4) BTC")
	fmt.Scanln(&opcion)
	fmt.Println("Convertido:", utilidades.Conversor(valor, opcion))

	// Contador de vocales
	var frase string
	fmt.Print("Ingresa una frase: ")
	fmt.Scanln(&frase)
	fmt.Println("Conteo de vocales:", utilidades.ContadorVocales(frase))
}
