package main

import (
	"./account"
	"./mydict"
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
	fmt.Println(account.Balance(), account.Owner(), account.String(), "\n\n\n")

	dictionary := mydict.Dictionary{"first": "Frist word"}
	fmt.Println(dictionary["first"])
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	word := "hello"
	definition2 := "Grreeting"
	err2 := dictionary.Add(word, definition2)
	if err != nil {
		fmt.Println(err2)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println(hello)
	err3 := dictionary.Add(word, definition)
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Println("\n")

	dictionary2 := mydict.Dictionary{}
	baseWord := "hello"
	dictionary2.Add(baseWord, "First")
	err4 := dictionary2.Update(baseWord, "Second")
	if err4 != nil {
		fmt.Println(err4)
	}
	word1, _ := dictionary2.Search(baseWord)
	fmt.Println(word1)
	dictionary2.Delete(baseWord)
	word2, err5 := dictionary2.Search(baseWord)
	if err5 != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word2)
	}
}
