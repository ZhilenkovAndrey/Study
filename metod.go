package main

import "fmt"

// func main() {
// 	defenition()
// }
// type Square struct {
// 	side int
// }
// func (s Square) Perimetr() {
// 	fmt.Printf("\n%T, %#v  \n", s, s)
// 	fmt.Printf("Периметр фигуры: %d \n", s.side*4)
// }
// func (s *Square) Scale(multiplier int) {
// 	fmt.Printf("\n%T, %#v\n", s, s)
// 	s.side *= multiplier
// }
// func defenition() {
// 	square := Square{side: 4}
// 	pSquare := &square
// 	square2 := Square{side: 2}
// 	square.Perimetr()
// 	square2.Perimetr()
// 	pSquare.Scale(2)
// 	pSquare.Perimetr()
// 	square.Scale(2)
// 	pSquare.Perimetr()
// }
type User struct {
	name     string
	age      int
	password string
	score    []int
}

func main() {
	chel := User{"John", 20, "pass", []int{23, 42, 54, 32, 76, 29}}
	fmt.Println(chel.getName())
	chel.setName("Bob")
	fmt.Println(chel.getName())
	chel.isElder()
	fmt.Println(chel.getHierScore())
}
func (u User) getName() string {
	return u.name
}
func (u *User) setName(name1 string) {
	u.name = name1
}
func (u User) isElder() bool {
	var a bool
	if u.age < 18 {
		println("Только 18+!!!")
		return a
	} else {
		println("Добро пожаловать!")
		return !a
	}
}
func (u User) getHierScore() int {
	hier := 0
	for _, sc := range u.score {
		if sc > hier {
			hier = sc
		}
	}
	return hier
}
