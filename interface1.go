package main

import . "fmt"

type Runner interface { //must have: 1 interface - 1 metod
	Run()
}

type Swimmer interface {
	Swimm()
}

type Flyer interface {
	Fly()
}

type Ducker interface { // embanding interfaces
	Runner
	Swimmer
	Flyer
}

type Human struct {
	Name string
}

type Duck struct {
	Name, Surname string
}

func (h Human) Run() { // implimentation
	Sprintf("Человек %s бегает", h.Name)
}

func (d Duck) Run() {
	Println("Утка бегает")
}

func interfaceValues() { // Если у интерфейсного значения нету

	var runner Runner //конкретного типа и значения, то он равен nil
	Printf("Type: %T, Value: %#v\n", runner, runner)

	if runner == nil {
		Println("Runner is nil")
	}

	var unnamedRunner *Human //оздаем переменную конкретного типа со значением nil
	Printf("Type: %T, Value: %#v\n", unnamedRunner, unnamedRunner)

	runner = unnamedRunner                           //присваиваем эу переменную интерфейсному значению, у runner
	Printf("Type: %T, Value: %#v\n", runner, runner) //появляется знание о типе(*Human)

	if runner == nil { //сравнивается значение интерфейсного типа, тип значения не nil,
		Println("Runner is nil") // интерейсное значение не равно nil
	} else {
		Println("Runner not nil")
	}

	namedRunner := &Human{Name: "Джек"}
	Printf("Type: %T, Value: %#v\n", namedRunner, namedRunner)

	runner = namedRunner
	Printf("Type: %T, Value: %#v\n", runner, runner)

	var emptyInterface interface{} = unnamedRunner                   //пустой интефейс, его имплиментирует
	Printf("Type: %T, Value: %#v\n", emptyInterface, emptyInterface) //любой тип в go

	emptyInterface = runner
	Printf("Type: %T, Value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = int64(1)
	Printf("Type: %T, Value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = true //в пустой интерфейс можно положить значение любого типа
	Printf("Type: %T, Value: %#v\n", emptyInterface, emptyInterface)
}

func typeAssertuonAndPolimorfism() {
	var runner = Runner
	Printf("Type: %tT, Value: %#v/n", runner, runner)

	john := &Human{"John"}
	runner = john
	polimorfism(john)

	donald := &Duck{"Donald", "Duck"}
	runner = donald
	polimorfism(donald)
}

func polimorfism(runner Runner) { //вызывая эту функцию у человека или у утки
	Println(runner.Run()) //мы абстрагируемся от конкретных типов
}

func main() {
	interfaceValues()
}
