package main

import (
	"fmt"
	// "strconv"
)

// var favoriteFruit string = "Mango"
// var favoriteNumber int = 123

func main() {

	// var favoriteBook string = "Shine"
	//favoriteBook := "Shine"

	// var num uint8 = 111
	// var floatnum float32 = 1.777
	// var strValueWithIntInside string = "123456"
	// var intValue int
	// //intValue, err := strconv.Atoi(strValueWithIntInside)
	// intValue, _ = strconv.Atoi(strValueWithIntInside)

	fmt.Println("Hello, world!")
	// fmt.Println("Int value: " + strconv.Itoa(favoriteNumber))
	// fmt.Printf("Int value: %d\n", num)
	// fmt.Printf("Int float value: %f\n", floatnum)
	// fmt.Print(intValue)
	fmt.Println(getSum(1, 100))
	fmt.Println(getSumAndSub(23, 77))
}

func getSum(value1, value2 int) int {
	return value1 + value2
}

func getSumAndSub(value1, value2 int) (sum int, sub int) {
	sum = value1 + value2
	sub = value2 - value1
	return sum, sub
}
