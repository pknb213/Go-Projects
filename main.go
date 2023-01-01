package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("Is Done~")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
	//return len(name), strings.ToUpper(name)
}

func multiply(a, b int) int {
	return a * b
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) (total int) {
	total = 0
	for _, numbers := range numbers {
		total += numbers
	}
	//for i := 0; i < len(numbers); i++ {
	//	fmt.Println(numbers[i])
	//}
	return
}

func canIDring(age int) bool {
	if age < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println("Hello world")
	fmt.Println(multiply(7, 7))
	totalLength, upperName := lenAndUpper("Cheon")
	fmt.Println(totalLength, upperName)
	repeatMe("Cheon", "Young", "Jo", "Goods")
	fmt.Println(superAdd(1, 2, 3, 4, 5, 6, 7))
	fmt.Println(canIDring(16))
}
