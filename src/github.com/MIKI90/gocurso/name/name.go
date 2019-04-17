package name

import "fmt"

// GetName obtiene y retorna nombre de la persona
func GetName() string {
	var name string
	name = "test name"
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	return name
}
