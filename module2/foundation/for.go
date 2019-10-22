package main

import "fmt"

func main() {
	var i int
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	names := []string{"john", "lenon"}
	m := map[string]int{
		"John":  28,
		"Lenon": 29,
	}
	for i := range names {
		fmt.Println(names[i])
	}

	for key, value := range m {
		fmt.Println(key, value)
	}

}
