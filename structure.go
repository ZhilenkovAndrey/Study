package main

import "fmt"

type People struct {
	Name    string
	Surname string
	Age     int
}

func (p People) Talk() {
	fmt.Printf(" %s talking!\n", p.Name)
}
func (p *People) grewUpAge() {
	p.Age++
	fmt.Printf("%s grew up, now he is %d!!!\n", p.Name, p.Age)
}
func main() {
	Vasya := People{"Vasya", "Pupckin", 21}
	fmt.Println(Vasya)
	Kolya := People{"Kolya", "Pupckin", 22}
	fmt.Println(Kolya)
	Vasya.Talk()
	Kolya.Talk()
	Vasya.grewUpAge()
	Kolya.grewUpAge()
	fmt.Println(Vasya)
	fmt.Println(Kolya)
}
