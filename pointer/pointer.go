package main

import "fmt"


type Student struct {
	ID int
	Name string
	GPA float32
}

type Gamer struct {
	Name string
	Games []string
}

//pointer struct sebagai parameter
// func graduate(student *Student){
// 	student.Name = student.Name + " S.T."
// }

//method dengan pointer receiver
func (student *Student)graduate(){
	student.Name = student.Name + " S.T."
}

//quis
func (gamer *Gamer) AddGame(game string){
	gamer.Games = append(gamer.Games, game)
}

func main() {
	//================================
	// pointer contoh 1

	// numberA := 5
	// numberB := &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// *numberB = 10 
	// fmt.Println(*numberB)
	// fmt.Println(numberA)

	//=========================
	// pointer contoh 2

	// var numberA int = 5
	// var numberB *int = &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// numberA = 20
	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	//===================================
	// contoh 3 pointer di passing ke function

	// number := 5
	// fmt.Println("alamat memori: ", &number)
	// fmt.Println("nilai awal :", number)
	// change(&number, 100)
	// fmt.Println("Nilai akhir :", number)
	// fmt.Println("alamat memori: ", &number)

	//====================================
	//contoh 4 pointer struct sebagai parameter

	// student := Student{1, "Sidiq pratomo", 3.21}
	// fmt.Println(student.Name)

	// graduate(&student)
	// fmt.Println(student.Name)

	//====================================
	// method dengan pointer receiver

	// student := Student{1, "Sidiq pratomo", 3.21}
	// fmt.Println(student.Name)

	// student.graduate()
	// fmt.Println(student.Name)

	//=====================================
	//quis

	gamer := Gamer{Name: "Zelda"}

	gamer.AddGame("god of war")
	gamer.AddGame("red dead redemption")
	gamer.AddGame("dishonored")

	for _, game := range gamer.Games{
		fmt.Println(game)
	}
}

func change(old *int, new int){				//pointer di passing ke function
	*old = new
	fmt.Println("Di dalam function :", *old)
	fmt.Println("alamat memori: ", old)
}

//pointer struct sebagai parameter


