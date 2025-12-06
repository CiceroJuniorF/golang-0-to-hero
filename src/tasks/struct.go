package tasks

import "fmt"

type User struct {
	Name     string
	Document string
	Age      int
}

func TestStruct() {
	user1 := User{"Joao", "123", 20}
	fmt.Println("Struct u1: ", user1)
	fmt.Println("Struct u1: Name: ", user1.Name)
	user2 := User{}
	p := &user2
	fmt.Println("Struct u2: ", user2, p)
	user3 := User{Document: "12345"}
	fmt.Println("Struct u3: ", user3)
}
