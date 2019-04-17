package structs

import "fmt"

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
