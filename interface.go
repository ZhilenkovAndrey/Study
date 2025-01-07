package main

import . "fmt"

type numbers struct {
	num1 int64
	num2 int64
}

type NumbersInterface interface {
	sum() int64
	multiply() int64
	division() float64
	substruct() int64
}

func (n numbers) sum() int64 {
	return n.num1 + n.num2
}

func (n numbers) multiply() int64 {
	return n.num1 * n.num2
}

func (n numbers) division() float64 {
	return float64(n.num1) / float64(n.num2)
}

func (n numbers) substruct() int64 {
	return n.num1 - n.num2
}

func main() {
	var i NumbersInterface
	numbers := numbers{5, 8}
	i = numbers
	Println(i.sum(), i.multiply(), i.division(), i.substruct())
}
