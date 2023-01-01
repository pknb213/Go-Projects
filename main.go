package main

import (
	"./account"
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

func canIDrking(age int) bool {
	if age > 20 {
		switch koreanAge := age + 2; koreanAge {
		case 10:
			return false
		case 18:
			return true
		}
	}
	return true
}

func pointers(val int) {
	addr := &val
	fmt.Println(val, addr, &val, &addr, *addr)
}

func array() {
	name := [5]string{"a", "b", "C"}
	name[3] = "d"
	name[4] = "f"
	fmt.Println(name)

	slice := []string{"S", "A"}
	slice = append(slice, "B", "C")
	fmt.Println(slice)
}

func mapTest() {
	me := map[string]string{"LastName": "Cheon", "FirstName": "YJ"}
	for k, v := range me {
		fmt.Println(k, v)
	}
	fmt.Println(me)
}

type person struct {
	name    string
	age     int
	favFood []string
}

func structTest() {
	food := []string{"Rice", "Egg", "Sosigy"}
	yj := person{"cyj", 32, food}
	fmt.Println(yj)

}

func main() {
	//fmt.Println("Hello world")
	//fmt.Println(multiply(7, 7))
	//totalLength, upperName := lenAndUpper("Cheon")
	//fmt.Println(totalLength, upperName)
	//repeatMe("Cheon", "Young", "Jo", "Goods")
	//fmt.Println(superAdd(1, 2, 3, 4, 5, 6, 7))
	//fmt.Println(canIDrking(16))
	//pointers(7)
	//array()
	//mapTest()
	//structTest()
	account := account.NewAccount("CYJ")
	account.Deposit(1000)
	fmt.Println(account)
	err := account.Withdraw(2000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
