package main

import (
	"fmt"
	"log"

	"github.com/montseisc/examen/contador")

func main() {
	log.SetPrefix("Contador: ")
	log.SetFlags(0)

	// llamada directa a cada funcion del modulo
	fmt.Println(contador.Canasta1())
	fmt.Println(contador.Canasta2())
	fmt.Println(contador.Canasta3())

	// contador de puntos con una lista de canastas
	canastas := []int{2, 3, 1, 2, 3}
	total := 0
	for _, puntos := range canastas {
		mensaje, err := contador.Anotar(puntos)
		// evaluacion de la salida de los datos
		if err != nil {
			log.Fatal(err)
		}
		total += puntos
		fmt.Println(mensaje)
	}
	fmt.Println("Total de puntos:", total)
}
