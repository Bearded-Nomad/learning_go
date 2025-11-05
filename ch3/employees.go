package main

import "fmt"

type employee struct {
	firstName string
	lastName  string
	id        int
}

func Employees() {
	john := employee{id: 0}
	jane := employee{firstName: "jane", lastName: "fonda", id: 1}
	var rob employee
	rob.firstName = "robert"
	rob.lastName = "downey"
	rob.id = 2

	fmt.Println(john, jane, rob)
}
