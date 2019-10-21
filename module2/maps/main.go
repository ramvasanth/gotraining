package main

import (
	"fmt"
	"sort"
	"sync"
)

type User struct {
	name string
	age  int
	user *User
}

type subscribers struct {
	list map[User]bool
	mu   sync.Mutex
}

func main() {
	//mapSizeCalculator()
	//simpleMap()
	mapOrderObject()
}
func mapDeclaration() {

}
func objectMap() {
	m := make(map[User]int)
	u1 := &User{}
	u2 := &User{}
	m[User{name: "hey", age: 18, user: u1}] = 100

	fmt.Println(m[User{name: "hey", age: 18, user: u2}])
}

func rangeOverMap() {
	m := make(map[string]int, 100) // initialize with capacity
	m["john"] = 18
	m["lenon"] = 19

	for key, value := range m { // you can range over nil map
		fmt.Println(key, value)
	}

}

func simpleMap() {
	var m map[string]int // nil  map - you can read but write causes runtime
	fmt.Println(m == nil)

	m = make(map[string]int, 100) // initialize with capacity
	m["john"] = 18

	value, found := m["john"]
	fmt.Println(value, found)

	//zero value when the key is not found
	value, found = m["lenon"]
	fmt.Println(value, found)

	//delete a key
	delete(m, "john")
	delete(m, "lenon") // no error when key is not found
}

func mapSizeCalculator() {
	//theMap := make(map[int]int, 100) // initialize with capacity
	//m[1] = 18
	//fmt.Println(len(m) * int(unsafe.Sizeof(m)))
	//unsafe.Sizeof(hmap) + (len(theMap) * 8) + (len(theMap) * 8 * unsafe.Sizeof(x)) + (len(theMap) * 8 * unsafe.Sizeof(y))
}

func mapConcurrency() {
	s := subscribers{}
	s.list = make(map[User]bool)
	add := func(user User) {
		s.mu.Lock()
		s.list[user] = true
		s.mu.Unlock()
	}
	remove := func(user User) {
		s.mu.Lock()
		delete(s.list, user)
		s.mu.Unlock()
	}
	add(User{})
	remove(User{})

}

func mapOrderSimple() {
	ages := map[string]int{
		"a": 1,
		"c": 3,
		"d": 4,
		"b": 2,
	}

	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names) //sort by key
	for _, name := range names {
		fmt.Println(ages[name])
	}
}

func mapOrderObject() {
	ages := map[User]bool{
		User{name: "cat"}:    true,
		User{name: "banana"}: true,
		User{name: "apple"}:  true,
	}

	fmt.Println(ages)
}
