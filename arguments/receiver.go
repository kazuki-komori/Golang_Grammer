package arguments

import "fmt"

type Person struct {
	name string
	age  int
}

// ポインタによるレシーバ
func (p *Person) setPointPerson() {
	p.name = "山田"
	p.age = 10
}

func ReceiverPointer() {
	p := new(Person)
	p.setPointPerson()
	fmt.Println(p.name, p.age)
}

//値によるレシーバ
func (p Person) printPerson() {
	fmt.Println(p.name, p.name)
}

func ReceiverPrinter() {
	p := Person{name: "山田", age: 10}
	p.printPerson()
}
