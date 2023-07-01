package ejercicios

import (
	"strconv"
)

func ConvertiraNumero(texto string) (int, string) {
	var mensaje string
	numero, error := strconv.Atoi(texto)

	if error != nil {
		return 0, "Hubo un error" + error.Error()
	}

	if numero < 100 {
		mensaje = "El numero es menor a 100"
	} else {
		mensaje = "El numero es mayor a 100"
	}

	return numero, mensaje
}
