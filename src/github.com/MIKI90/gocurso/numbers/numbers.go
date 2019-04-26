package numbers

import "errors"

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

//función de suma excepcion
func ExcepSum(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
	case string:
		return 0, errors.New("El valor A es un string")

	}
	switch b.(type) {
	case string:
		return 0, errors.New("El valor B es un string")
	}
	return a.(int) + b.(int), nil
}
