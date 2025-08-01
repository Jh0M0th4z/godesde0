package ejercicio

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func TablaMultiplicar() {

	fmt.Println("Ingrese un número: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Número no válido.")
			TablaMultiplicar()
		} else {
			for i := 1; i <= 10; i++ {
				fmt.Println(i, " * ", numero, " = ", i*numero)
			}
		}
	}

}
