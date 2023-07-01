package main

import (
	"fmt"
	"godesde0/ejercicios"
	"godesde0/variables"
	"runtime"
)

func main() {
	variables.ConviertoaTexto(1588)
	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("Esto no es windows, es", os)
	} else {
		fmt.Println("Esto es windows")
	}

	switch os := runtime.GOOS; os {
	case "linuxa":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	case "windows":
		fmt.Println("Esto es windows")
	default:
		fmt.Printf("%s \n", os)
	}

	numero, mensaje := ejercicios.ConvertiraNumero("120")
	fmt.Println(numero)
	fmt.Println(mensaje)
}
