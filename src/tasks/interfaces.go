package tasks

import "fmt"

type I interface {
	Abs() int
	Negative() int
}

type ImplementsI struct {
	Number int
}

func (i ImplementsI) Abs() int {
	if i.Number < 0 {
		return -i.Number
	}
	return i.Number
}

func (i ImplementsI) Negative() int {
	if i.Number > 0 {
		return -i.Number
	}
	return i.Number
}

func (i ImplementsI) String() string {
	return fmt.Sprintf("Number Implements I %d", i.Number)
}

type ImplementsI2 struct {
	Number int
}

func (i ImplementsI2) Abs() int {
	if i.Number < 0 {
		return -i.Number
	}
	return i.Number
}

func (i ImplementsI2) Negative() int {
	if i.Number > 0 {
		return -i.Number
	}
	return i.Number
}

func emptyInterface(v interface{}) {
	fmt.Println(v)
}

func ExecInterfaces() {
	var i I = ImplementsI{Number: -100}

	fmt.Println(i.Abs())
	fmt.Println(i.Negative())

	emptyInterface("ABC")
	emptyInterface(i)
	emptyInterface(123)

	t, ok := i.(ImplementsI)

	if ok {
		emptyInterface(t)
	}

	t2, ok2 := i.(ImplementsI2)

	if ok2 {
		emptyInterface(t2)
	}

	switch i.(type) {
	case ImplementsI:
		fmt.Println("ImplementsI")
	case ImplementsI2:
		fmt.Println("ImplementsI2")

	}
}
