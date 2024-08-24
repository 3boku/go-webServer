package main

import "fmt"

type hotdog int

type person struct {
	Fname string
	Lname string
}

type secretAgen struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.Fname, p.Lname, `says, Good morning`)
}

func (sa secretAgen) speak() {
	fmt.Println(sa.Fname, sa.Lname, `says, Shanken`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	//var t hotdog
	//t = 7
	//fmt.Println(t)
	//
	//xi := []int{3, 4, 5, 6, 7, 8}
	//fmt.Println(xi)
	//
	//m := map[string]int{
	//	"Todd": 45,
	//	"Job":  42,
	//}
	//
	//fmt.Println(m)

	p1 := person{
		"Hyunseo",
		"Jung",
	}

	p1.speak()

	sa1 := secretAgen{
		person{
			"James",
			"Bond",
		},
		true,
	}

	sa1.speak()
	sa1.person.speak()

	saySomething(p1)
	saySomething(sa1)
}
