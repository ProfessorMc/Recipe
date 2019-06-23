package main

import (
	"fmt"
	"github.com/ProfessorMc/Recipe/spoilers/friend"
)

var friends []*friend.Friend

func lab2() {
	fmt.Println("-- Lab2 Friends --")
	friends = make([]*friend.Friend,0)
	// Add mom
	mom := friend.NewFriend("Mom", 55)
	friends = append(friends, mom)

	// Add dad
	dad := friend.NewFriend("Dad", 60)
	dad.SetLiesAboutAge(5)
	dad.CelebrateBirthday()
	friends = append(friends, dad)

	// Add chuck
	chuck := friend.NewFriend("Chuck", 35)
	friends = append(friends, chuck)

	greetGuests()
}

func greetGuests() {
	for _, friend := range friends {
		greeting := friend.GreetFriend()
		fmt.Print(greeting)
	}
}
