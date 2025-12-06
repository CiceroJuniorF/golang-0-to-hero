package tasks

import "fmt"

type UserMap struct {
	Name, Document string
}

func MakeASimpleMap() {
	const USER1, USER2 = "user1", "user2"
	x := map[string]*UserMap{
		USER1: {Name: "Jhon", Document: "189761"},
		USER2: {Name: "Peter", Document: "198129"},
	}
	fmt.Println("Make A Simple Map: ", x)
	fmt.Println("Make A Simple Map: ", USER1, x[USER1])
	fmt.Println("Make A Simple Map: ", USER2, x[USER2])

	x[USER1].Name = "Marcel"
	fmt.Println("The value:", USER1, x[USER1])

	x[USER2].Name = "Marcel"
	fmt.Println("The value:", USER2, x[USER2])

	delete(x, USER1)
	fmt.Println("The value:", USER1, x[USER1])

	v, ok := x[USER1]
	fmt.Println("The value:", v, "Present?", ok)

}
