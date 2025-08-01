package ejercicio

import (
	"fmt"
	"strconv"
)

func Get2Values(param string) (int, string) {
	intValue, err := strconv.Atoi(param)

	if err != nil {
		return 0, fmt.Sprintf("No se pudo convertir el valor %s", param)
	}
	if intValue > 100 {
		return intValue, "Es mayor a 100"
	} else {
		return intValue, "Es menor a 100"
	}
}
