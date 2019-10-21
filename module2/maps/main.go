package main

import "fmt"

type User struct {
	name string
	age  int
	user *User
}

func main() {
	m := make(map[User]int)
	u1 := &User{}
	u2 := &User{}
	m[User{name: "hey", age: 18, user: u1}] = 100

	fmt.Println(m[User{name: "hey", age: 18, user: u2}])
}
