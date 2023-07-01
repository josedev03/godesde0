package variables

import "fmt"

func MostrarEnteros() {
	var intComun int
	intDe32 := int32(10)
	intDe64 := int64(100)

	Nombre = "Camilo"
	nombreAndApellido := Nombre + " Rodriguez"

	fmt.Println(intComun)
	fmt.Println(intDe32)
	fmt.Println(intDe64)
	fmt.Println(nombreAndApellido)
}
