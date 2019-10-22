package main

import "fmt"

func breakcase1() {
	/* local variable definition */
	var a int = 10

	/* for loop execution */
	for a < 20 {
		fmt.Printf("value of a: %d\n", a)
		a++
		if a > 15 {
			/* terminate the loop using break statement */
			break
		}
	}
}

func breakCase2() {
	for outer := 0; outer < 5; outer++ {
		if outer == 3 {
			fmt.Println("Breaking out of outer loop")
			break // break here
		}
		fmt.Println("The value of outer is", outer)
		for inner := 0; inner < 5; inner++ {
			if inner == 2 {
				fmt.Println("Breaking out of inner loop")
				break // break here
			}
			fmt.Println("The value of inner is", inner)
		}
	}
	fmt.Println("Exiting program")
}

func continuecase() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Continuing loop")
			continue // break here
		}
		fmt.Println("The value of i is", i)
	}
	fmt.Println("Exiting program")
}
