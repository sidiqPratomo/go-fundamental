package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	Name        string
	Admin       string
	Users       []User
	IsAvailable bool
}

func main() {
	user := User{ID: 1, FirstName: "Gojo", LastName: "Satoru", Email: "satoru@mailcom", IsActive: true}

	user2 := User{
		ID:        2,
		FirstName: "Itadori",
		LastName:  "Yuuji",
		Email:     "Yuuji@mailcom",
		IsActive:  true,
	}
	user3 := User{}
	user3.ID = 3
	user3.FirstName = "Fushiguro"
	user3.LastName = "Megumi"
	user3.Email = "Megumi@mailcom"
	user3.IsActive = false

	_,_ = user, user3

	// displayUser(user3) using method user.display()
	result := user2.Display()
	fmt.Println(result)

	// displayGroup using method
	users := []User{user, user2, user3}
	group := Group{"Gamer", "administrator", users, true}
	group.DisplayGroup()

}

func displayUser(user User) string {
	result := fmt.Sprintf("Name : %s %s, Email : %s",user.FirstName, user.LastName, user.Email)
	return result
}

//method displayUser
func (user User) Display() string{
	return fmt.Sprintf("Name : %s %s, Email : %s",user.FirstName, user.LastName, user.Email)
}

func (group Group) DisplayGroup(){
	fmt.Printf("Name : %s \n", group.Name)
	fmt.Printf("Member count : %d \n", len(group.Users))
	fmt.Println("Users name: ")
	for _, user := range group.Users {
		fmt.Println(user.FirstName, user.LastName)
	}
}