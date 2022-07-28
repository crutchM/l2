package Patterns

import (
	"fmt"
)

type Employee interface {
	FullName()
	Accept(visitor Visitor)
}
type Developer struct {
	FirstName string
	LastName  string
	Income    float32
	Age       int
}

func (d Developer) FullName() {
	fmt.Println("Developer", d.FirstName, " ", d.LastName)
}

func (d Developer) Accept(v Visitor) {
	v.VisitDeveloper(d)
}

func (d Director) Accept(v Visitor) {
	v.VisitDirector(d)
}

type Director struct {
	FirstName string
	LastName  string
	Income    float32
	Age       int
}

func (d Director) FullName() {
	fmt.Println("Director ", d.FirstName, " ", d.LastName)
}

type Visitor interface {
	VisitDeveloper(d Developer)
	VisitDirector(d Director)
}

type Income struct {
	rate float32
}

func (i Income) VisitDeveloper(d Developer) {
	fmt.Println(d.Income + d.Income*i.rate/100)
}

func (i Income) VisitDirector(d Director) {
	fmt.Println(d.Income + d.Income*i.rate/100)
}

func test() {
	backend := Developer{"Bob", "Bilbo", 1000, 32}
	boss := Director{"Bob", "Baggins", 2000, 40}
	backend.FullName()
	backend.Accept(Income{20})
	boss.FullName()
	boss.Accept(Income{10})
}
