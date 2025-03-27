package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func (p person) print() {
	fmt.Println(p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	var yannic person

	yannic.firstName = "yannic"
	yannic.contact.email = "yannic@loritz.ch"

	pointerYannic := &yannic
	pointerYannic.updateName("goht di nix ah ")

	yannic.print()
}
