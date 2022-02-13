package main

import "fmt"

func main() {
	//declaracion de constantes
	const pi float64 = 3.1416
	const pi2 = 3.141569845

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	/* Zero values: valores por defecto que se dan cuando
	se define una variable pero no se asigna, depende del tipo de variable
	*/
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Área cuadrado

	const base_cuadrado = 10
	area_cuadrado := base_cuadrado * base_cuadrado
	fmt.Println(area_cuadrado)

	// Operadores Aritmeticos

	x := 10
	y := 50

	// suma
	result := x + y
	fmt.Println("Suma:", result)

	//resta
	result = x - y
	fmt.Println("Resta:", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	// Division
	result = y / x
	fmt.Println("Division:", result)

	// Módulo
	result = y % x
	fmt.Println("Módulo:", result)

	//incremental
	x++
	fmt.Println("Incremental de x", x)

	//decremental
	y--
	fmt.Println("Decremental de y", y)

}
