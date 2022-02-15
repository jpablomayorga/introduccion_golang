package main

import "fmt"

func main() {
	// Array
	var myArrray [4]int
	myArrray[0] = 5
	myArrray[1] = 20
	myArrray[2] = -5
	fmt.Println(myArrray, len(myArrray), cap(myArrray))

	// Slice
	mySlice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(mySlice, len(mySlice), cap(mySlice))

	//metodos en el slice
	fmt.Println(mySlice[0])
	fmt.Println(mySlice[:3])
	fmt.Println(mySlice[2:4])

	//append

	mySlice = append(mySlice, 7)
	fmt.Println(mySlice)

	//append nueva lista
	newSlice := []int{8, 9, 10, 11, 12}
	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)

}
