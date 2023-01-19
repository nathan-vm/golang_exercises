package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println(p.name, "say hello")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{"Nathan", 24}

	p1.speak()    // <- Isso é um shortcut para (&p1).speak()
	(&p1).speak() // <- Isso é a maneira "correta"

	saySomething(&p1)
	// saySomething(p1) <-- Não funciona.
}
