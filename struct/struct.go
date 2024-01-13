package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	email     string
	isActive  bool
}

func main() {
	user := User{}
	user.ID = 1
	user.FirstName = "test"
	user.LastName = "ting"
	user.email = "testing@mail.com"
	user.isActive = true
	
	
	user2 := User{}
	user2.ID = 2
	user2.FirstName = "Link"
	user2.LastName = "edIn"
	user2.email = "LinkedIn@mail.com"
	user2.isActive = true

	user3 := User{			//bentuk deklarasi lain struct secara langsung
		ID : 3,
		FirstName: "Gojo",
		LastName: "Satoru",
		email: "satoru@gmail.com",
		isActive: true,
	}

	fmt.Println(user.FirstName)
	fmt.Println(user2)
	fmt.Println(user3)

	displayUser := displayUser(user3)

	fmt.Println(displayUser)
}

//struct sebagai parameter function

func displayUser (user User) string {
	result := fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.email)
	return result
}