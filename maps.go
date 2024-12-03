package main

import "fmt"

type User struct {
	Id   int
	Name string
}

func main() {
	var defaultMap map[int64]string //map{nil}
	fmt.Printf("Type: %T, Value: %#v\n\n", defaultMap, defaultMap)

	mapByMake := make(map[string]string) //map{}
	fmt.Printf("Type: %T, Value: %#v\n\n", mapByMake, mapByMake)

	mapByMakeWithCap := make(map[string]string, 10) //map{}
	fmt.Printf("Type: %T, Value: %#v\n\n", mapByMakeWithCap, mapByMakeWithCap)

	mapByLiteral := map[string]int{"Vasia": 18, "Dima": 20} //map{v 18, D 20}
	fmt.Printf("Type: %T, Value: %#v\n", mapByLiteral, mapByLiteral)
	fmt.Printf("Lenght: %d\n\n", len(mapByLiteral))

	mapByNew := *new(map[string]string) //map{nil}
	fmt.Printf("Type: %T, Value: %#v\n\n", mapByNew, mapByNew)

	mapByMake["First"] = "Vasia" //внесение значения
	fmt.Printf("Type: %T, Value: %#v\n", mapByMake, mapByMake)
	fmt.Printf("Lenght: %d\n\n", len(mapByMake))

	mapByMake["First"] = "Petia" //изменение значения
	fmt.Printf("Type: %T, Value: %#v\n", mapByMake, mapByMake)
	fmt.Printf("Lenght: %d\n\n", len(mapByMake))

	fmt.Println(mapByLiteral["Vasia"]) // get value

	fmt.Println(mapByLiteral["Second"]) // get non-existent value

	value, ok := mapByLiteral["Second"] // check value existence
	fmt.Printf("Value: %d, IsExist: %t\n\n", value, ok)

	delete(mapByMake, "First") //deleting by key
	fmt.Printf("Type: %T, Value: %#v\n\n", mapByMake, mapByMake)

	mapIteration := map[string]int{"First": 1, "Second": 2,
		"Third": 3, "Fourth": 4}

	for key, value := range mapIteration { //no subsequence iteration
		fmt.Printf("Key: %s, Value: %d \n", key, value)
	}

	Users := []User{
		{1, "Vasia"},
		{45, "Petya"},
		{57, "john"},
		{45, "Petya"},
	}

	UniqueUsers := make(map[int]struct{}, len(Users))

	for _, user := range Users { //like set
		if _, ok := UniqueUsers[user.Id]; !ok {
			UniqueUsers[user.Id] = struct{}{}
		}
	}

	fmt.Printf("Type: %T, Value: %#v \n\n", UniqueUsers, UniqueUsers)

	UsersMap := make(map[int]User, len(Users))

	for _, user := range Users {
		if _, ok := UsersMap[user.Id]; !ok {
			UsersMap[user.Id] = user
		}
	}

	fmt.Println(FindInSlice(57, Users))
	fmt.Println(FindInMap(57, UsersMap))

}

func FindInMap(id int, UsersMap map[int]User) *User {
	if User, ok := UsersMap[id]; ok {
		return &User
	}
	return nil
}

func FindInSlice(id int, users []User) *User {
	for _, user := range users {
		if user.Id == id {
			return &user
		}
	}
	return nil
}
