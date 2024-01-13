// exported and unexported unction

package management

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

