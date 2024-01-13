package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

// embedded struct
type Group struct {
	Name string
	Admin string
	Users []User
	IsAvailable bool
}

func main() {
	user := User{}
	user.ID = 1
	user.FirstName = "test"
	user.LastName = "ting"
	user.Email = "testing@mail.com"
	user.IsActive = true
	
	
	user2 := User{}
	user2.ID = 2
	user2.FirstName = "Link"
	user2.LastName = "edIn"
	user2.Email = "LinkedIn@mail.com"
	user2.IsActive = true

	user3 := User{			//bentuk deklarasi lain struct secara langsung
		ID : 3,
		FirstName: "Gojo",
		LastName: "Satoru",
		Email: "satoru@gmail.com",
		IsActive: true,
	}
	_,_,_ = user,user2,user3
	// fmt.Println(user.FirstName) //dasar struct memanggil
	// fmt.Println(user2)
	// fmt.Println(user3)

	// displayUser := displayUser(user3) //struct sebagai parameter
	// fmt.Println(displayUser)

	users := []User{user, user3, user2}

	group := Group{"Gamer", "user", users, true }

	displayGroup(group)
}

//struct sebagai parameter function

func displayUser (user User) string {
	result := fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
	return result
}

//embedded struct
func displayGroup(group Group){
	fmt.Printf("Name: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d \n", len(group.Users))
	
	fmt.Println("Name user : ")
	for _, user := range group.Users{
		fmt.Println(user.FirstName)
	}
}	