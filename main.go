package main

import "fmt"

type contactInfo struct {
	email         string
	contactNumber string
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// cards := newDeck()

	//cards.print()

	// hand, remainingCards := deal(cards, 3)

	// hand.print()
	// fmt.Println("\n")
	// remainingCards.print()

	//fmt.Println(cards.toString())
	// cards.saveToFile("mycards")
	// cards := newDeckFromFile("mycards")
	// cards.shuffle()
	// cards.print()

	//Structs Examples

	p1 := person{firstName: "Ayesha", lastName: "kaleem"}

	p1.print()
	var p2 person
	p2.firstName = "Aisha"
	p2.lastName = "Kalim"
	p2.print()

	p3 := person{
		firstName: "ayesha0",
		lastName:  "kaleem0",
		contactInfo: contactInfo{
			email:         "akal**@gmail.com",
			contactNumber: "+9854256***4",
		},
	}

	(&p3).updateName("ramsha")
	p3.print()

	p4 := map[string]string{
		"fname": "ayesha",
		"lname": "kaleem",
	}
	printMap(p4)
	// remove key from map
	delete(p4, "lname")
	printMap(p4)

	//create empty map
	p5 := make(map[string]string)
	printMap(p5)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("\n %+v", p)
}
func printMap(p map[string]string) {
	for k, v := range p {
		fmt.Println(k, v)
	}
}
