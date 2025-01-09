package main

import . "fmt"

type Runner interface { //must have: 1 interface - 1 metod
	Run() string
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

func (h Human) Run() string { // implimentation
	return Sprintf("Человек %s бегает", h.Name)
}

func (d Duck) Run() string {
	return "Утка бегает"
}

func interfaceValues() Runner { // Если у интерфейсного значения нету

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

	return runner
}

func typeAssertionAndPolimorfism() Runner {
	var runner Runner
	Printf("Type: %T, Value: %#v\n", runner, runner)

	john := &Human{"John"}
	runner = john
	polimorfism(john)

	donald := &Duck{"Donald", "Duck"}
	runner = donald
	polimorfism(donald)

	return runner
}

func polimorfism(runner Runner) { //вызывая эту функцию у человека или у утки
	Println(runner.Run()) //мы абстрагируемся от конкретных типов
}

func (h Human) writeCode() {
	Println("Human write code")
}

func (d Duck) fly() string {
	return "duck fly"
}
func (d Duck) Swimm() string {
	return "duck swimm"
}

func typeAssertion(runner Runner) { //type assertion - проверка типа
	Printf("Type: %T, Value: %#v\n", runner, runner)
	if human, ok := runner.(*Human); ok {
		Printf("Type: %T, Value: %#v\n", human, human)
		human.writeCode()
	}
	if duck, ok := runner.(*Duck); ok {
		Printf("Type: %T, Value: %#v\n", duck, duck)
		Println(duck.fly())
	}
}

func typeSwitch(runner Runner) { //проверка типа
	switch v := runner.(type) {
	case *Human:
		Println(v.Run())
	case *Duck:
		Println(v.Swimm())
	default:
		Printf("Type: %T, Value: %#v\n", v, v)
	}
}

func main() {
	interfaceValues()
	typeAssertionAndPolimorfism()
	typeAssertion(interfaceValues())
	typeSwitch(typeAssertionAndPolimorfism())
}
