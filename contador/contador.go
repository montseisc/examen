package contador

import "fmt"

func Canasta1() string {
	return "Canasta de 1 punto"
}
func Canasta2() string {
	return "Canasta de 2 puntos"
}
func Canasta3() string {
	return "Canasta de 3 puntos"
}

func Anotar(puntos int) (string, error) {
	switch puntos {
	case 1:
		return Canasta1(), nil
	case 2:
		return Canasta2(), nil
	case 3:
		return Canasta3(), nil
	default:
		return "", fmt.Errorf("puntos invalidos: %d (solo 1, 2 o 3)", puntos)
	}
}
