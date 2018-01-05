package main

import "fmt"

type foo int

type person struct {
	first string
	last  string
	age   int
}

type doubleZero struct {
	person
	secretAgent bool
}

func (p person) fullName() string {
	return p.first + " " + p.last
}

func (d doubleZero) fullName() string {
	thisName := ""
	if d.secretAgent {
		thisName = "**** ****"
	} else {
		thisName = d.person.fullName()
	}
	return thisName
}

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)

	var yourAge int
	yourAge = 53
	fmt.Printf("%T %v \n", yourAge, yourAge)

	yourAge = int(myAge)
	fmt.Printf("%T %v \n", myAge, myAge)

	p1 := person{"Adam", "First", 2}
	p2 := person{"Eve", "First", 2}
	p3 := doubleZero{
		person: person{
			first: "Tom",
			last:  "Sawyer",
			age:   33,
		},
		secretAgent: false,
	}
	p4 := doubleZero{
		person: person{
			first: "Tom",
			last:  "Sawyer",
			age:   33,
		},
		secretAgent: true,
	}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.fullName(), p2.age)
	fmt.Println(p3)
	fmt.Println(p3.fullName())
	fmt.Println(p4.fullName())
	// myAge = int(yourAge)

}
