package numbers

//función devuelve numeros enteros
func GetVariables() (int, int32, int64) {
	return 1, 2, 3
}

//función que devulve floats
func GetFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

//función de suma
func Sum(a int, b int) int {
	return a + b
}
