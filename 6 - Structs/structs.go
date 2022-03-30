package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Age       int
	Address   Address
}

type Address struct {
	City  string
	State string
}

func main() {
	var user User = User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Address: Address{
			City:  "City X",
			State: "State X",
		},
	}

	fmt.Println(user)
	fmt.Println(user.FirstName)
	fmt.Println(user.Address.City)
}
