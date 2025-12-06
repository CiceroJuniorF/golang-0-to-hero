package tasks

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string) Person {
	return Person{Name: name}
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func MethodsExec() {
	p := NewPerson("Jhon")
	p.SetAge(10)
	fmt.Println(p)
}
