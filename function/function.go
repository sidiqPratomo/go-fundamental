package main

import "fmt"

func main() {
	// sentence := printMyResults("saya sedang")
	// fmt.Println(sentence)
	// printMyResult("belajar")
	// printMyResult("golang")

	// result := add(10, 20)
	// fmt.Println(result)

	luas, keliling := calculate(10, 2)
	fmt.Println(luas)
	fmt.Println(keliling)
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

//1.input
//2.proses
//3. output
