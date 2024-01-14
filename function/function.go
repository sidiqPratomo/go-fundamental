package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	// sentence := printMyResults("saya sedang")
	// fmt.Println(sentence)
	// printMyResult("belajar")
	// printMyResult("golang")

	// result := add(10, 20)
	// fmt.Println(result)

	// luas, keliling := calculate(10, 2) //function multiple return dan kemudian dipecah
	// fmt.Println(luas)
	// fmt.Println(keliling)

	// scores := []int{10, 5, 8, 9, 7}
	// total := sum(scores) //function dengan input parameter slice
	// fmt.Println(total)

	// result, err := hitung(10, 2, "*")
	// if err != nil {
	// 	fmt.Println("terjadi kesalahan")
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(result)

	// //named return value
	// a, b, c := getFullName2()
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	//variadic function
	jumlahTotal := sumAll(10, 90, 30, 50, 40)
	fmt.Println(jumlahTotal)
	fmt.Println(strings.Repeat("=", 20))
	slice := []int{10,20,30,40,50} //input parameter berupa slice
	jumlahTotal = sumAll(slice...) //variable slice berupa slice harus dipecah sebelum di passing dengan slice... ditambah 3 titik
	fmt.Println(jumlahTotal)

	// function as value
	sayGoodBye := getGoodBye

	hasil := sayGoodBye("Eko") // function di simpan di variabel
	fmt.Println(hasil)			// gunanya untuk mempassing function lain sebagai parameter
	fmt.Println(getGoodBye("Eko")) //jadi untuk function as parameter

	//function as parameter
	sayHelloWithFilter("Eko", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	//anonymous function
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("eko", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("eko", func(name string) bool {
		return name == "root"
	})
}

func printMyResult(sentence string) {
	fmt.Println(sentence)
}

func printMyResults(sentence string) string {
	newSentence := sentence + " Belajar Golang"
	return newSentence
}

func add(number int, numberTwo int) int {
	result := number + numberTwo
	return result
}

//function multiple return
func calculate(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}

// function dengan predefined return value
func calculates(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar

	keliling = 2 * (panjang + lebar)

	return luas, keliling
}

// quiz
func sum(numbers []int) int {		//function dengan input parameter slice
	// scores := []int{10, 5, 8, 9, 7}
	var total int
	for _, number := range numbers {
		total = total + number
	}
	return total
}

func hitung(number, numberTwo int, operation string) (int, error) {
	var result int
	var errorResult error

	switch operation {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "*":
		result = number * numberTwo
	case "/":
		result = number / numberTwo
	default:
		errorResult = errors.New("Unknown operation")
	}
	return result, errorResult
}

//Named return value
func getFullName2() (firstName string, middleName string, lastName string) {
	firstName = "Eko"
	middleName = "Kurniawan"
	lastName = "Khannedy"

	return
}

//variadic function 
func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

//function as value
func getGoodBye(name string) string {
	return "Good Bye " + name
}

//function sebagai parameter

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}else {
		return name
	}
}

//anonymous function
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

//1.input
//2.proses
//3. output
