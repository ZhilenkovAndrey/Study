package main

import "fmt"

// type User struct {
// 	Id   int
// 	Name stirig
// }

func main() {
	var defaultMap map[int64]string
	fmt.Printf("Type: %T, Value: %#v\n\n", defaultMap, defaultMap)

	mapByMake := make(map[string]string)
	fmt.Printf("Type: %T, Value: %#v\n\n", mapByMake, mapByMake)

	mapByMakeWithCap := make(map[string]string, 10)
	fmt.Printf("Type: %T, Value: %#v\n\n", mapByMakeWithCap, mapByMakeWithCap)

	mapByLiteral := map[string]int{"Vasia": 18, "Dima": 20}
	fmt.Printf("Type: %T, Value: %#v\n", mapByLiteral, mapByLiteral)
	fmt.Printf("Lenght: %d\n\n", len(mapByLiteral))

	mapByNew := *new(map[string]string)
	fmt.Printf("Type: %T, Value: %#v\n\n", mapByNew, mapByNew)

	mapByMake["First"] = "Vasia"
	fmt.Printf("Type: %T, Value: %#v\n", mapByMake, mapByMake)
	fmt.Printf("Lenght: %d\n\n", len(mapByMake))

	mapByMake["First"] = "Petia"
	fmt.Printf("Type: %T, Value: %#v\n", mapByMake, mapByMake)
	fmt.Printf("Lenght: %d\n\n", len(mapByMake))
}
