package main

import (
	"fmt"
	"math"
)

/*
汉字
Ver
తె
தி

*/

// == , !=  comparable
// <,>, >=,<=  //ordered

type s1 struct {
	Name string
	Age  int
}

type s2 struct {
	Name string
	Age  int
}

type cArray [10]int

func main() {
	// comparable ==, !=
	// ordered <,> <=, >=

}

func runeLiterals() {
	// utf-8 - 4 bytes
	//

	fmt.Println(len("字"), rune('字'))
	fmt.Println(len("V"), rune('V'))
	fmt.Println(len("தி"))
	fmt.Println(string(int32(23383)))
}
func stringRawandInterpret() {
	// internal structure-> []byte
	//8 bits - 255

	fmt.Println("This is my name \n --")
	fmt.Println(`This is my name \n --`)
}

func digitRepresentation() {
	//base 10,  2  ,    16 , 8
	fmt.Println(15, 0B1111, 0xF, 0o17)
}
func stringRanging() {
	for _, v := range "தி" {
		fmt.Printf("%T\n", v)
	}
	fmt.Println(len("A") == len("தி"))

	for i := 0; i < len("தி"); i++ {
		fmt.Println(i)
	}
}
func f(i ...string) {
	fmt.Printf("%T\n", i)
}

type customFunc func(string, string) error
type customInt int
type customMap map[string]string

func boolCompare() {
	a := true
	b := false

	fmt.Println(a == b)

}

func stringCompare() {

	fmt.Println("ac" > "aa")

}

func integerAndFloatCompare() {
	a := 3.1415
	if a != math.Pi {
		fmt.Println("a is not pi")
	}
}

func complexCompare() {
	a := complex(-3.25, 2)
	b := -3.25 + 2i
	if a == b {
		fmt.Println("a complex as b")
	}
}

type Person1 struct {
	a string
	b int
}

type Person2 struct {
	a string
	b int
}

func structCompare1() {
	cp1 := Person1{}
	cp2 := Person2{}
	cp1 = Person1(cp2)
	fmt.Println(cp1)
}

func structCompare() {
	p1 := struct {
		a string
		b int
	}{"left", 4}

	p2 := struct {
		a string
		b int
	}{a: "left", b: 4}
	if p1 == p2 {
		fmt.Println("Same position")
	}
}

func arrayCompare() {
	pair1 := [2]int{4, 2}
	pair2 := [2]int{2, 4}
	if pair1 != pair2 {
		fmt.Println("different pair")
	}

}

func arrayIdentical() {
	pair1 := [2]int{4, 2}
	pair2 := [2]int{2, 4}
	pair2 = pair1

}

func SliceIdentical() {
	pair1 := []int{4, 2}
	pair2 := []int{2, 4}
	pair2 = pair1
}

func structIden() {
	p1 := struct {
		a string
		b []int
	}{}

	p2 := struct {
		a string
		b []int
	}{}
	p1 = p2
}
