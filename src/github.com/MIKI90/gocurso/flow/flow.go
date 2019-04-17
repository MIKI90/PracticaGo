package flow

import "fmt"

//Prueba de if
func IfTest() {
	var number = 0
	fmt.Print("Ingresa un num del 1 al 10: ")
	fmt.Scanf("%d", number)
	if number%2 == 0 {
		fmt.Println("El numero es par")
	} else {
		fmt.Println("El numero es impar")
	}

	if number2 := 3; number2 == 3 {
		fmt.Println("Pues si es ", number2)
	}
}

//Prueba de switch
func SwitchTest() {
	var number int
	fmt.Print("Type a number [Switch]: ")
	fmt.Scanf("%d", &number)
	fmt.Scanf("%d", &number)
	switch number {
	case 1:
		fmt.Println("El número es 1 ")
	default:
		fmt.Printf("\nEl número es: %d \n", number)
	}

	switch {
	case number%2 == 0:
		fmt.Println("El número es par ")
	default:
		fmt.Println("El número es impar")
	}
}
