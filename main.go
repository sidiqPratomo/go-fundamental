package main

import (
	"fmt"
	"pustaka-api/management"
	// "net/http"
	// "github.com/gin-gonic/gin"
)

func main() {
	// router := gin.Default()

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"name":   "sidiq pratomo",
	// 		"occupy": "sealabs bootcamp",
	// 	})
	// })
	// router.GET("/hello", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"title":    "Hello",
	// 		"subtitle": "Belajar Golang",
	// 	})
	// })

	// router.Run(":8888")

	//=================================================
	// ekxport import from other files from management entity
	user := management.User{ID: 1, FirstName: "Gojo", LastName: "Satoru", Email: "satoru@mailcom", IsActive: true}

	user2 := management.User{
		ID:        2,
		FirstName: "Itadori",
		LastName:  "Yuuji",
		Email:     "Yuuji@mailcom",
		IsActive:  true,
	}
	user3 := management.User{}
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
	users := []management.User{user, user2, user3}
	group := management.Group{"Gamer", "administrator", users, true}
	group.DisplayGroup()
}
