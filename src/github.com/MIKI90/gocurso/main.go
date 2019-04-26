package main

import (
	"fmt"
	"strings"
	"time"
	/*"github.com/MIKI90/gocurso/flow"
	"github.com/MIKI90/gocurso/name"
	"github.com/MIKI90/gocurso/numbers"
	"github.com/MIKI90/gocurso/structs"*/)

const helloWorld string = "Hola %s %s bienvenido \n"
const test = "Test"

func main() {
	/*
		firstName := name.GetName()
		lastname := "<modifica apellido>"
		var number = numbers.Sum(50, 50)
		a, b, c := numbers.GetVariables()
		f32, f64 := numbers.GetFloat()
		stringUTF8 := name.GetUnicode()
		fmt.Print("Ingresa tu nombre: ")
		fmt.Scanf("%s", &firstName)
		fmt.Printf(helloWorld, firstName, lastname)
		fmt.Print("hello world")
		fmt.Println(number, a, b, c)
		fmt.Println(f32, f64)
		fmt.Println("cadena con UTF8: ", stringUTF8)
		fmt.Println("hola"[1])
		fmt.Println(string("hola"[1]))
		fmt.Println(len("hola"))
		structs.GetArray()
		structs.GetSlice()
		flow.IfTest()
		flow.SwitchTest()
		forTest()
		strings2()*/
	// fmt.Println(maps.GetMap())
	// fmt.Println(maps.GetMap2())
	// fmt.Println(maps.GetMapName("MIKI"))
	// structs.InterfaceTest()
	// resultSum, err := numbers.ExcepSum(50, 50)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("El resultado de la suma es: %d\n", resultSum)
	//pointerTest()
	go forGo(500)
	go forGo(400)

	time.Sleep(10000 * time.Millisecond)
}

func forTest() {
	for i := 0; i < 5; i++ {
		fmt.Println("[FOR] ", i)
	}

	c := 100
	for c > 0 {
		c -= 10
		fmt.Println("[FOR] solo con una condicion de c > 0 ", c)
	}

	s := 1000

	for {
		s--
		if s == 0 {
			fmt.Println("Termina For 'infinito'")
			break
		}
	}
}

func strings2() {
	var text = "Hello world, Hello miki, Hello Go"
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text, "Hello", "Hola", -1))
	fmt.Println(strings.Replace(text, "Hello", "Hola", 1))
	fmt.Println(strings.Replace(text, "Hello", "Hola", 2))
	fmt.Println(strings.Replace(text, "Hello", "Hola", 3))
	fmt.Println(strings.Split(text, ","))
}

func pointerTest() {
	a := 100
	var b *int
	b = &a
	fmt.Println("muestra puntero")
	fmt.Println(a, b)
	fmt.Println(a, *b)
	fmt.Println(&a, *b)
	fmt.Println("modificando puntero")
	*b = 10
	fmt.Println(a, *b)
	fmt.Println(&a, *b)
	modifyPointer(b)
	fmt.Println("modificando puntero en una funcion")
	fmt.Println(a, *b)
	fmt.Println(&a, *b)
}

func modifyPointer(c *int) {
	*c = 3000
}

func helloGo(index int) {
	fmt.Println("Go rutine number #", index)
}

func forGo(n int) {
	for i := 0; i < n; i++ {
		go helloGo(i)
	}
}
