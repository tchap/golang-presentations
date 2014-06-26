package main

import "fmt"

// START STRUCT OMIT
type Person struct {
	id      int // not visible from other packages
	Name    string
	Address string
}

// type *Person, allocate on the heap
var p = &Person{Name: "Pepa Novak", Address: "Praha"}

// factory functions are conventional
func NewPerson(name, address string) *Person {
	return &Person{
		Name:    name,
		Address: address,
	}
}

func (p *Person) Greet(name string) {
	fmt.Printf("Hello %s, my name is %s.", name, p.Name)
}
// END STRUCT OMIT

// START INTERFACE OMIT
type Greeter interface {
	Greet(name string)
}
// END INTERFACE OMIT

func main() {
	// Person is a greeter, but no need to specify it.
	var g Greeter = NewPerson("Pepa Novak", "Praha")
	g.Greet("Honza")
}
