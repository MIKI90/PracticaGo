package maps

// regresa un mapa
func GetMap() map[string]int {
	mapTest := make(map[string]int)
	mapTest["llave1"] = 1
	mapTest["llave2"] = 2
	return mapTest
}

// regresa un mapa con valores iniciales
func GetMap2() map[string]int {
	mapTest := map[string]int{
		"Juan":  18,
		"Yohan": 24,
		"MIKI":  100,
	}
	mapTest["llave1"] = 1
	mapTest["llave2"] = 2
	delete(mapTest, "Juan")
	delete(mapTest, "Yohan")
	return mapTest
}

// regresa un mapa con valores iniciales
func GetMapName(name string) int {
	mapTest := map[string]int{
		"Juan":  18,
		"Yohan": 24,
		"MIKI":  100,
	}
	return mapTest[name]
}
