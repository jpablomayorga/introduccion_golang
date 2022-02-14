package main

import "fmt"

func main() {
	//for condicional
	for i := 0; i <= 100; i++ {
		fmt.Println("El valor de i es:", i)
	}

	//for while
	counter := 0
	for counter < 10 {
		fmt.Println("Valor de contador:", counter)
		counter++
	}

	//for forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}

}
