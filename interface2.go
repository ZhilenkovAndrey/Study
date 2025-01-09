package main

import . "fmt"

type animal interface {
	makeSound()
}

type greeter interface {
	greet(string) string
}

type russian struct{}
type american struct{}

func (r russian) greet(name string) string {
	return Sprintf("Привет! %s!", name)
}

func (a american) greet(name string) string {
	return Sprintf("Hello! %s!", name)
}

func sayHello(g greeter, name string) {
	Println(g.greet(name))
}

type cat struct{}
type dog struct{}

func (c *cat) makeSound() {
	Println("Mau, mau")
}

func (d *dog) makeSound() {
	Println("Gav, gav")
}

func main() {

	(&cat{}).makeSound()
	(&dog{}).makeSound()

	sayHello(&russian{}, "Олег")
	sayHello(&american{}, "John")
}
