package main

import . "fmt"

type animal interface {
	walker
	runner
}

type bird interface {
	walker
	flyer
}

type walker interface {
	walk()
}

type runner interface {
	run()
}

type flyer interface {
	fly()
}

type cat struct{}
type eagle struct{}

func (e eagle) walk() {
	Println("Eagle walk")
}

func (e eagle) fly() {
	Println("Eagle fly")
}

func walk(w walker) {
	w.walk()
}

func fly(f flyer) {
	f.fly()
}

func (c *cat) walk() {
	Println("Cat walk")
}

func (c *cat) run() {
	Println("Cat run")
}

func main() {
	var cat animal = &cat{}
	var eagle bird = &eagle{}
	cat.walk()
	cat.run()

	walk(cat)
	walk(eagle)
	fly(eagle)
}
