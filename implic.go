package main

import (
	"fmt"
)

func implic(a bool, b bool) bool {
	//if a == true && b == false {
	//		return false
	//	} else {
	//		return true
	//	}
	return !a || b
}
func main() {
	fmt.Println(implic(true, true))
	fmt.Println(implic(false, false))
	fmt.Println(implic(true, false))
	fmt.Println(implic(false, true))
}
