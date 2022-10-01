package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	contract  Contract
}

type Contract struct {
	mobile string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func updateSlice(s []string) {
	s[0] = "1"
}

func main() {
	s := []string{"2", "3"}
	updateSlice(s)
	fmt.Println(s)
	fmt.Println(len(s))
}
