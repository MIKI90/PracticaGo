package structs

import "fmt"

// PlatziCourse es una estructura
type PlatziCourse struct {
	Name   string
	Slug   string
	Skills []string
}

// PlatziCareer es una estructura que toma los métodos de PlatziCourse
type PlatziCareer struct {
	PlatziCourse
}

//subscripcion al curso
func (p PlatziCourse) Subscribe(name string) {
	fmt.Printf("La persona %s se ha registrado al curso %s\n", name, p.Name)
}

//subscripcion al carrera
func (p PlatziCareer) Subscribe(name string) {
	fmt.Printf("La persona %s se ha registrado al carrera %s\n", name, p.PlatziCourse.Name)
}

//función obtener arreglos
func GetArray() {
	var arr1 [2]string
	arr2 := [3]int{1, 2, 3}
	arr1[0] = "valor 1"
	arr1[1] = "valor 2"
	fmt.Println(arr1, arr2)
}

//función obtener Slice
func GetSlice() {
	var slice1 []string
	slice1 = append(slice1, "mi", "slice", "1")
	fmt.Println(slice1)
}
