package main

import (
	"fmt"
	"time"
)

func intergerLiterals() {
	println(0b1111)
	println(15)
	println(0xF)
	println(0o17)
}

func switch1() {
	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
	case 15:
		fmt.Println("Today is 15th. Visit a doctor.")
	case 25:
		fmt.Println("Today is 25th. Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}

// ,
func switch2() {

	today := time.Now()
	var t int = today.Day()

	switch t {
	case 5,
		10,
		15:
		fmt.Println("Clean your house.")
	case 25, 26, 27:
		fmt.Println("Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}

func switch3() {
	today := time.Now()
	switch today.Day() {
	case 5:
		fmt.Println("Clean your house.")
		fallthrough
	case 10:
		fmt.Println("Buy some wine.")
		fallthrough
	case 15:
		fmt.Println("Visit a doctor.")
		fallthrough
	case 25:
		fmt.Println("Buy some food.")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}

func switch4() {
	today := time.Now()

	switch {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <= 10:
		fmt.Println("Buy some wine.")
	case today.Day() > 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}
}

func typeswitch(x interface{}) {

	switch i := x.(type) {
	case nil:
		println("x is nil") // type of i is type of x (interface{})
	case int:
		println("int")
	case int32:
		println(i) // type of i is int32
	case float64:
		println(i) // type of i is float64
	case func(int) float64:
		println(i) // type of i is func(int) float64
	case bool, string:
		println("type is bool or string") // type of i is type of x (interface{})
	default:
		println("don't know the type") // type of i is type of x (interface{})
	}
}
