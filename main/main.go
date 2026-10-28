package main

import (
	"fmt"
	"log"

	"github.com/montseisc/examen/contador"
)

func main() {
	log.SetPrefix("Contador: ")
	log.SetFlags(0)

	fmt.Println(contador.Canasta1())
	fmt.Println(contador.Canasta2())
	fmt.Println(contador.Canasta3())

	canastas := []int{2, 3, 1, 2, 3}
	for _, puntos := range canastas {
		mensaje, err := contador.Anotar(puntos)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(mensaje)
	}
}
