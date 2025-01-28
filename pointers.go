package main

import . "fmt"

func square(num int) { //не меняет значение переменной
	num *= num
}

func squarePointer(num *int) { //изменяет значение переменной
	*num *= *num
}

func hasWallet(money *int) bool {
	return money != nil
}

func main() {
	num := 3
	square(num)
	Printf("variable value not changed, num = %d\n", num)

	squarePointer(&num)
	Printf("Variable value has changed, num = %d\n", num)

	var wallet1 *int
	Printf("Wallet1 exists = %t\n", hasWallet(wallet1))

	wallet2 := 0
	Printf("Wallet2 exists = %t\n", hasWallet(&wallet2))

	wallet3 := 100
	Printf("Wallet3 exists = %t\n", hasWallet(&wallet3))
}
